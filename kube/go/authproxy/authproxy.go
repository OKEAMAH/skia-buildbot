// auth-proxy is a reverse proxy that runs in front of applications and takes
// care of authentication.
//
// This is useful for applications like Promentheus that doesn't handle
// authentication itself, so we can run it behind auth-proxy to restrict access.
//
// The auth-proxy application also adds the X-WEBAUTH-USER header to each
// authenticated request and gives it the value of the logged in users email
// address, which can be used for audit logging. The application running behind
// auth-proxy should then use:
//
//     https://pkg.go.dev/go.skia.org/infra/go/alogin/proxylogin
//
// When using --cria_group this application should be run using work-load
// identity with a service account that as read access to CRIA, such as:
//
//     skia-auth-proxy-cria-reader@skia-public.iam.gserviceaccount.com
//
// See also:
//
//     https://chrome-infra-auth.appspot.com/auth/groups/project-skia-auth-service-access
//
//     https://grafana.com/blog/2015/12/07/grafana-authproxy-have-it-your-way/
package authproxy

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"go.skia.org/infra/go/allowed"
	"go.skia.org/infra/go/cleanup"
	"go.skia.org/infra/go/common"
	"go.skia.org/infra/go/httputils"
	"go.skia.org/infra/go/skerr"
	"go.skia.org/infra/go/sklog"
	"go.skia.org/infra/kube/go/authproxy/auth"
	"golang.org/x/oauth2/google"
)

const (
	appName            = "auth-proxy"
	serverReadTimeout  = time.Hour
	serverWriteTimeout = time.Hour
	drainTime          = time.Minute
)

// Send the logged in user email in the following header. This allows decoupling
// of authentication from the core of the app. See
// https://grafana.com/blog/2015/12/07/grafana-authproxy-have-it-your-way/ for
// how Grafana uses this to support almost any authentication handler.
const webAuthHeaderName = "X-WEBAUTH-USER"

type proxy struct {
	allowPost    bool
	passive      bool
	reverseProxy http.Handler
	authProvider auth.Auth
}

func newProxy(target *url.URL, authProvider auth.Auth, allowPost bool, passive bool) *proxy {
	return &proxy{
		reverseProxy: httputil.NewSingleHostReverseProxy(target),
		authProvider: authProvider,
		allowPost:    allowPost,
		passive:      passive,
	}
}

func (p proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	email := p.authProvider.LoggedInAs(r)
	r.Header.Del(webAuthHeaderName)
	r.Header.Add(webAuthHeaderName, email)
	if r.Method == "POST" && p.allowPost {
		p.reverseProxy.ServeHTTP(w, r)
		return
	}
	if !p.passive {
		if email == "" {
			http.Redirect(w, r, p.authProvider.LoginURL(w, r), http.StatusSeeOther)
			return
		}
		if !p.authProvider.IsViewer(r) {
			http.Error(w, "403 Forbidden", http.StatusForbidden)
			return
		}
	}
	p.reverseProxy.ServeHTTP(w, r)
}

// App is the auth-proxy application.
type App struct {
	port        string
	promPort    string
	criaGroup   string
	local       bool
	targetPort  string
	allowPost   bool
	allowedFrom string
	passive     bool

	target       *url.URL
	authProvider auth.Auth
	server       *http.Server
}

// Flagset constructs a flag.FlagSet for the App.
func (a *App) Flagset() *flag.FlagSet {
	fs := flag.NewFlagSet(appName, flag.ExitOnError)
	fs.StringVar(&a.port, "port", ":8000", "HTTP service address (e.g., ':8000')")
	fs.StringVar(&a.promPort, "prom-port", ":20000", "Metrics service address (e.g., ':10110')")
	fs.StringVar(&a.criaGroup, "cria_group", "", "The chrome infra auth group to use for restricting access. Example: 'google/skia-staff@google.com'")
	fs.BoolVar(&a.local, "local", false, "Running locally if true. As opposed to in production.")
	fs.StringVar(&a.targetPort, "target_port", ":9000", "The port we are proxying to.")
	fs.BoolVar(&a.allowPost, "allow_post", false, "Allow POST requests to bypass auth.")
	fs.StringVar(&a.allowedFrom, "allowed_from", "", "A comma separated list of of domains and email addresses that are allowed to access the site. Example: 'google.com'")
	fs.BoolVar(&a.passive, "passive", false, "If true then allow unauthenticated requests to go through, while still adding logged in users emails in via the webAuthHeaderName.")

	return fs
}

// New returns a new *App.
func New(ctx context.Context) (*App, error) {
	var ret App

	err := common.InitWith(
		appName,
		common.PrometheusOpt(&ret.promPort),
		common.MetricsLoggingOpt(),
		common.FlagSetOpt(ret.Flagset()),
	)
	if err != nil {
		return nil, skerr.Wrap(err)
	}

	err = ret.validateFlags()
	if err != nil {
		return nil, skerr.Wrap(err)
	}

	var allow allowed.Allow
	if ret.criaGroup != "" {
		ctx := context.Background()
		ts, err := google.DefaultTokenSource(ctx, "email")
		if err != nil {
			return nil, skerr.Wrap(err)
		}
		criaClient := httputils.DefaultClientConfig().WithTokenSource(ts).With2xxOnly().Client()
		allow, err = allowed.NewAllowedFromChromeInfraAuth(criaClient, ret.criaGroup)
		if err != nil {
			return nil, skerr.Wrap(err)
		}
	} else {
		allow = allowed.NewAllowedFromList(strings.Split(ret.allowedFrom, ","))
	}

	authInstance := auth.New()
	authInstance.SimpleInitWithAllow(ret.port, ret.local, nil, nil, allow)
	targetURL := fmt.Sprintf("http://localhost%s", ret.targetPort)
	target, err := url.Parse(targetURL)
	if err != nil {
		return nil, skerr.Wrap(err)
	}
	ret.authProvider = authInstance
	ret.target = target
	ret.registerCleanup()

	return &ret, nil
}

func (a *App) registerCleanup() {
	cleanup.AtExit(func() {
		if a.server != nil {
			sklog.Info("Shutdown server gracefully.")
			ctx, cancel := context.WithTimeout(context.Background(), drainTime)
			err := a.server.Shutdown(ctx)
			if err != nil {
				sklog.Error(err)
			}
			cancel()
		}
	})

}

// Run starts the application serving, it does not return unless there is an
// error or the passed in context is cancelled.
func (a *App) Run(ctx context.Context) error {
	var h http.Handler = newProxy(a.target, a.authProvider, a.allowPost, a.passive)
	h = httputils.HealthzAndHTTPS(h)
	server := &http.Server{
		Addr:           a.port,
		Handler:        h,
		ReadTimeout:    serverReadTimeout,
		WriteTimeout:   serverWriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	a.server = server

	sklog.Infof("Ready to serve on port %s", a.port)
	err := server.ListenAndServe()
	if err == http.ErrServerClosed {
		// This is an orderly shutdown.
		return nil
	}
	return skerr.Wrap(err)
}

func (a *App) validateFlags() error {
	if a.criaGroup != "" && a.allowedFrom != "" {
		return fmt.Errorf("Only one of the flags in [--auth_group, --allowed_from] can be specified.")
	}
	if a.criaGroup == "" && a.allowedFrom == "" {
		return fmt.Errorf("At least one of the flags in [--auth_group, --allowed_from] must be specified.")
	}

	return nil
}

// Main constructs and runs the application. This function will only return on failure.
func Main() error {
	ctx := context.Background()
	app, err := New(ctx)
	if err != nil {
		return skerr.Wrap(err)
	}

	return app.Run(ctx)
}