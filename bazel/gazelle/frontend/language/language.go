package language

import (
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/label"
	"github.com/bazelbuild/bazel-gazelle/language"
	"github.com/bazelbuild/bazel-gazelle/rule"
	"go.skia.org/infra/bazel/gazelle/frontend/common"
	"go.skia.org/infra/bazel/gazelle/frontend/configurer"
	"go.skia.org/infra/bazel/gazelle/frontend/parsers"
	"go.skia.org/infra/bazel/gazelle/frontend/resolver"
	"go.skia.org/infra/go/util"
)

// Language implements the language.Language interface.
type Language struct {
	configurer.Configurer
	resolver.Resolver

	// TargetDirectories is a set of known good directories for which we can currently generate valid
	// build targets. This Gazelle extension will not generate build targets for any other directories
	// in the repository.
	//
	// The value of this map indicates whether to recurse into the directory.
	//
	// If nil, no directories will be ignored.
	//
	// TODO(lovisolo): Delete after this Gazelle extension is fully fleshed out.
	TargetDirectories map[string]bool
}

// isTargetDirectory returns true if this Gazelle extension should generate or update the BUILD file
// in the given directory.
func (l *Language) isTargetDirectory(dir string) bool {
	if l.TargetDirectories == nil {
		return true
	}
	for targetDir, recursive := range l.TargetDirectories {
		if dir == targetDir || (recursive && strings.HasPrefix(dir, targetDir+"/")) {
			return true
		}
	}
	return false
}

// Kinds implements the language.Language interface.
//
// Interface documentation:
//
// Kinds returns a map of maps rule names (kinds) and information on how to
// match and merge attributes that may be found in rules of those kinds. All
// kinds of rules generated for this language may be found here.
func (l *Language) Kinds() map[string]rule.KindInfo {
	return map[string]rule.KindInfo{
		"karma_test": {
			NonEmptyAttrs:  map[string]bool{"src": true},
			MergeableAttrs: map[string]bool{"src": true},
			ResolveAttrs:   map[string]bool{"deps": true},
		},
		"nodejs_test": {
			NonEmptyAttrs:  map[string]bool{"src": true},
			MergeableAttrs: map[string]bool{"src": true},
			ResolveAttrs:   map[string]bool{"deps": true},
		},
		"sass_library": {
			NonEmptyAttrs:  map[string]bool{"srcs": true},
			MergeableAttrs: map[string]bool{"srcs": true},
			ResolveAttrs:   map[string]bool{"deps": true},
		},
		"sk_demo_page_server": {
			NonEmptyAttrs:  map[string]bool{"sk_page": true},
			MergeableAttrs: map[string]bool{"sk_page": true},
		},
		"sk_element": {
			MatchAny:       true,
			NonEmptyAttrs:  map[string]bool{"ts_srcs": true, "sass_srcs": true},
			MergeableAttrs: map[string]bool{"ts_srcs": true, "sass_srcs": true},
			ResolveAttrs:   map[string]bool{"sass_deps": true, "sk_element_deps": true, "ts_deps": true},
		},
		"sk_element_puppeteer_test": {
			NonEmptyAttrs:  map[string]bool{"src": true, "sk_demo_page_server": true},
			MergeableAttrs: map[string]bool{"src": true, "sk_demo_page_server": true},
			ResolveAttrs:   map[string]bool{"deps": true},
		},
		"sk_page": {
			NonEmptyAttrs:  map[string]bool{"html_file": true, "ts_entry_point": true, "scss_entry_point": true},
			MergeableAttrs: map[string]bool{"html_file": true, "ts_entry_point": true, "scss_entry_point": true},
			ResolveAttrs:   map[string]bool{"sass_deps": true, "sk_element_deps": true, "ts_deps": true},
		},
		"ts_library": {
			NonEmptyAttrs:  map[string]bool{"srcs": true},
			MergeableAttrs: map[string]bool{"srcs": true},
			ResolveAttrs:   map[string]bool{"deps": true},
		},
	}
}

// Loads implements the language.Language interface.
//
// Interface documentation:
//
// Loads returns .bzl files and symbols they define. Every rule generated by
// GenerateRules, now or in the past, should be loadable from one of these
// files.
func (l *Language) Loads() []rule.LoadInfo {
	return []rule.LoadInfo{
		{
			Name: "//infra-sk:index.bzl",
			Symbols: []string{
				"karma_test",
				"nodejs_test",
				"sass_library",
				"sk_demo_page_server",
				"sk_element",
				"sk_element_puppeteer_test",
				"sk_page",
				"ts_library",
			},
		},
	}
}

// importsParsedFromRuleSourcesImpl implements the common.ImportsParsedFromRuleSources interface.
type importsParsedFromRuleSourcesImpl struct {
	sassImports []string
	tsImports   []string
}

// GetSassImports implements the common.ImportsParsedFromRuleSources interface.
func (i *importsParsedFromRuleSourcesImpl) GetSassImports() []string {
	return i.sassImports
}

// GetTypeScriptImports implements the common.ImportsParsedFromRuleSources interface.
func (i *importsParsedFromRuleSourcesImpl) GetTypeScriptImports() []string {
	return i.tsImports
}

var _ common.ImportsParsedFromRuleSources = &importsParsedFromRuleSourcesImpl{}

// GenerateRules implements the language.Language interface.
//
// GenerateRules generates build rules for source files in a directory. GenerateRules is called in
// each directory where an update is requested in depth-first order.
//
// This method does not populate the deps argument of any generate rules. Dependencies are resolved
// in Resolver.Resolve(), which happens after GenerateRules is called in each directory where an
// update is requested.
func (l *Language) GenerateRules(args language.GenerateArgs) language.GenerateResult {
	// Unit tests use a made-up directory structure, so we can skip these checks.
	if !l.IsUnitTest {
		// Skip known directories with third-party code.
		for _, dir := range strings.Split(args.Rel, "/") {
			if util.In(dir, []string{"node_modules", "bower_components"}) {
				return language.GenerateResult{}
			}
		}

		// Limit generation of build targets to a hard-coded list of known good directories.
		// TODO(lovisolo): Delete after this Gazelle extension is fully fleshed out.
		if !l.isTargetDirectory(args.Rel) {
			return language.GenerateResult{}
		}
	}

	// Return values.
	var rules []*rule.Rule
	var imports []common.ImportsParsedFromRuleSources

	allFiles := append(args.RegularFiles, args.GenFiles...)

	// Directories are classified into three different groups based on their name, which determines
	// the kinds of rules that will be generated:
	//
	// - Unclassified directories:
	//   - Rules generated:
	//     - One ts_library rule for each *.ts file that does not end with "_test.ts".
	//     - One nodejs_test rule for each file ending with "_nodejs_test.ts".
	//     - One kara_test rule for each file ending with "_test.ts" and not "_nodejs_test.ts".
	//     - One sass_library rule for each *.scss file.
	//
	// - Directories with a custom element:
	//   - Pattern: //<app name>/modules/<custom element name ending in -sk, e.g. my-element-sk>
	//   - Rules generated:
	//     - sk_element:
	//       - Only generated if a my-element-sk.ts file is found.
	//       - If an index.ts file is found, it will be added to the ts_srcs argument.
	//       - If a my-element-sk.scss file is found, it will be added to the sass_srcs argument.
	//     - sk_page:
	//       - Only generated if files my-element-sk-demo.html a my-element-sk-demo.ts are found.
	//       - If a my-element-sk-demo.scss file is found, it will be added as the scss_entry_point
	//         argument.
	//     - sk_demo_page_server:
	//       - Only generated if an sk_page rule is produced.
	//     - sk_element_puppeteer_test:
	//       - Only generated if a my-element-sk_puppeteer_test.ts file is found, and if an
	//         sk_demo_page_server is produced.
	//     - Any other files follow the same criteria as for unclassified directories. This includes
	//       the element's karma_test, demo and test data defined in separate *.ts files, etc.
	//
	// - Directories with application pages:
	//   - Pattern: //<app name>/pages.
	//   - Rules generated:
	//     - sk_page:
	//       - One is generated for each pair of <page name>.html and <page name>.ts files found.
	//       - If a <page name>.scss file is found, it will be added as the scss_entry_point argument.
	//     - One ts_library rule for any other *.ts file that does not end with "_test.ts".
	//     - One sass_library rule for any other *.scss file.

	// Application page directories follow the "<app name>/pages" pattern.
	if isAppPageDir(args.Dir) {
		// This map will store the source files of any application pages found in the current directory.
		pages := map[string]*skPageSrcs{}
		getPage := func(name string) *skPageSrcs {
			if pages[name] == nil {
				pages[name] = &skPageSrcs{}
			}
			return pages[name]
		}

		// Populate the pages map with all HTML, TypeScript and Sass files found in the directory.
		for _, f := range allFiles {
			if strings.HasSuffix(f, "_test.ts") {
				log.Printf("Ignoring TypeScript test file found in directory with application pages: %s", filepath.Join(args.Dir, f))
				continue
			}

			name := strings.TrimSuffix(f, filepath.Ext(f)) // e.g. my-page.html -> my-page
			switch filepath.Ext(f) {
			case ".html":
				getPage(name).html = f
			case ".ts":
				getPage(name).ts = f
			case ".scss":
				getPage(name).scss = f
			}
		}

		// Generate sk_page targets for all pages for which we have all the necessary files (i.e.
		// my-page.html and my-page.ts), or stand-alone ts_library and sass_library targets for any
		// TypeScript and Sass files that do not belong to a page.
		for _, page := range pages {
			// A page is valid if it has an HTML file and a TypeScript file.
			if page.isValid() {
				r, i := generateSkPageRule(page, args.Dir)
				rules = append(rules, r)
				imports = append(imports, i)
			} else {
				if page.ts != "" {
					r, i := generateTSLibraryRule(page.ts, args.Dir)
					rules = append(rules, r)
					imports = append(imports, i)
				}
				if page.scss != "" {
					r, i := generateSassLibraryRule(page.scss, args.Dir)
					rules = append(rules, r)
					imports = append(imports, i)
				}
			}
		}

		return makeGenerateResult(args, rules, imports)
	}

	// Custom element directories follow the "<app name>/modules/<element-name-sk>" pattern.
	isCustomElementDir, customElementName := extractCustomElementNameFromDir(args.Dir)

	// If we are in a custom element directory, it will contain at most one custom element and one
	// custom page. Let's find the source files for both, and generate the corresponding sk_element,
	// sk_page and sk_demo_page_server rules.
	customElementSrcs := &skElementSrcs{}
	demoPageSrcs := &skPageSrcs{}
	skDemoPageServerLabel := label.NoLabel // We'll need this later for the Puppeteer test.
	if isCustomElementDir {
		// Iterate over all files and add them to the appropriate structs.
		indexTsFound := false
		for _, f := range allFiles {
			switch f {
			case "index.ts":
				indexTsFound = true
			case customElementName + ".ts":
				customElementSrcs.ts = f
			case customElementName + ".scss":
				customElementSrcs.scss = f
			case customElementName + "-demo.html":
				demoPageSrcs.html = f
			case customElementName + "-demo.ts":
				demoPageSrcs.ts = f
			case customElementName + "-demo.scss":
				demoPageSrcs.scss = f
			}
		}

		// An index.ts file alone does not make an sk_element, so we will include it in the returned
		// skElementSrcs struct only if the struct has other sources as well.
		if indexTsFound && !customElementSrcs.isEmpty() {
			customElementSrcs.indexTs = "index.ts"
		}

		// Generate the rules.
		if customElementSrcs.isValid() {
			r, i := generateSkElementRule(customElementName, customElementSrcs, args.Dir)
			rules = append(rules, r)
			imports = append(imports, i)
		}
		if demoPageSrcs.isValid() {
			skPage, i := generateSkPageRule(demoPageSrcs, args.Dir)
			rules = append(rules, skPage)
			imports = append(imports, i)

			skDemoPageServerRule, i := generateSkDemoPageServerRule(label.Label{Repo: "", Pkg: "", Name: skPage.Name(), Relative: true})
			rules = append(rules, skDemoPageServerRule)
			imports = append(imports, i)

			skDemoPageServerLabel = label.Label{Repo: "", Pkg: "", Name: skDemoPageServerRule.Name(), Relative: true}
		}
	}

	// Generate rules for all other files found in the directory.
	for _, f := range allFiles {
		if isCustomElementDir {
			// Skip any files that belong to the custom element or demo page.
			if (customElementSrcs.isValid() && customElementSrcs.has(f)) || (demoPageSrcs.isValid() && demoPageSrcs.has(f)) {
				continue
			}
		}

		if strings.HasSuffix(f, ".scss") {
			r, i := generateSassLibraryRule(f, args.Dir)
			rules = append(rules, r)
			imports = append(imports, i)
		} else if strings.HasSuffix(f, "_nodejs_test.ts") {
			r, i := generateNodeJSTestRule(f, args.Dir)
			rules = append(rules, r)
			imports = append(imports, i)
		} else if strings.HasSuffix(f, "_puppeteer_test.ts") {
			if skDemoPageServerLabel != label.NoLabel {
				r, i := generateSkElementPuppeteerTestRule(f, args.Dir, skDemoPageServerLabel)
				rules = append(rules, r)
				imports = append(imports, i)
			} else if isCustomElementDir {
				log.Printf("Not generating an sk_element_puppeteer_test rule for %s because %s has no demo page.", filepath.Join(args.Rel, f), customElementName)
			} else {
				log.Printf("Not generating an sk_element_puppeteer_test rule for %s because %s does not follow the custom element directory naming convention (<app>/modules/<element name>-sk).", filepath.Join(args.Rel, f), args.Rel)
			}
		} else if strings.HasSuffix(f, "_test.ts") {
			r, i := generateKarmaTestRule(f, args.Dir)
			rules = append(rules, r)
			imports = append(imports, i)
		} else if strings.HasSuffix(f, ".ts") {
			r, i := generateTSLibraryRule(f, args.Dir)
			rules = append(rules, r)
			imports = append(imports, i)
		}
	}

	return makeGenerateResult(args, rules, imports)
}

// makeGenerateResult returns a language.GenerateResult with the results of generating build rules
// for a directory.
func makeGenerateResult(args language.GenerateArgs, rules []*rule.Rule, imports []common.ImportsParsedFromRuleSources) language.GenerateResult {
	if len(rules) != len(imports) {
		log.Panicf("Arguments rules and imports must be of the same length (lengths: %d, %d; directory: %s)", len(rules), len(imports), args.Rel)
	}

	// Sort the rules and imports slices by rule name to guarantee a deterministic result.
	type ruleImportsPair struct {
		rule    *rule.Rule
		imports common.ImportsParsedFromRuleSources
	}
	var ruleImportsPairs []ruleImportsPair
	for i, r := range rules {
		ruleImportsPairs = append(ruleImportsPairs, ruleImportsPair{r, imports[i]})
	}
	sort.Slice(ruleImportsPairs, func(i, j int) bool {
		return ruleImportsPairs[i].rule.Name() < ruleImportsPairs[j].rule.Name()
	})
	rules = nil
	imports = nil
	for _, ri := range ruleImportsPairs {
		rules = append(rules, ri.rule)
		imports = append(imports, ri.imports)
	}

	// The Imports field in language.GenerateResult is of type []interface{}, so we need to cast our
	// imports slice to []interface{}.
	var importsAsEmptyInterfaces []interface{}
	for _, i := range imports {
		importsAsEmptyInterfaces = append(importsAsEmptyInterfaces, i)
	}

	return language.GenerateResult{
		Gen:     rules,
		Imports: importsAsEmptyInterfaces,
		Empty:   generateEmptyRules(args),
	}
}

var (
	// appPagesDirRegexp matches directories where application pages are found (sk_page targets), e.g.
	// "myapp/pages".
	//
	// In order to support absolute paths, this regexp does not start with ^.
	appPagesDirRegexp = regexp.MustCompile(`(?P<app_name>(?:[[:alnum:]]|_|-)+)/pages$`)

	// skElementModuleDirRegexp matches directories that might contain an sk_element, e.g.
	//"myapp/modules/my-element-sk".
	//
	// In order to support absolute paths, this regexp does not start with ^.
	skElementModuleDirRegexp = regexp.MustCompile(`(?P<app_name>(?:[[:alnum:]]|_|-)+)/modules/(?P<element_name>(?:[[:alnum:]]|_|-)+-sk)$`)
)

// isAppPageDir returns true if the directory matches the "<app name>/pages" pattern, which
// indicates it might contain application pages (sk_page targets).
func isAppPageDir(dir string) bool {
	return appPagesDirRegexp.MatchString(dir)
}

// extractCustomElementNameFromDir determines whether the given directory corresponds to a custom
// element based on the "<app name>/modules/<element-name-sk>" pattern, and returns the element name
// if the directory matches said pattern.
func extractCustomElementNameFromDir(dir string) (bool, string) {
	match := skElementModuleDirRegexp.FindStringSubmatch(dir)
	if len(match) != 3 {
		return false, ""
	}
	return true, match[2]
}

// skElementSrcs groups together the various sources that could make an sk_element target.
type skElementSrcs struct {
	indexTs string // index.ts
	ts      string // my-element-sk.ts
	scss    string // my-element-sk.scss
}

// isValid returns true if the struct contains the necessary sources to build an sk_element, or
// false otherwise.
func (e *skElementSrcs) isValid() bool {
	return e.ts != ""
}

// isEmpty returns true if the structure does not contain any source files.
func (e *skElementSrcs) isEmpty() bool {
	return *e == skElementSrcs{}
}

// has returns true if the struct includes the given source file, or false otherwise.
func (e *skElementSrcs) has(src string) bool {
	return src == e.indexTs || src == e.ts || src == e.scss
}

// skPageSrcs groups together the various sources that could make an sk_page target.
type skPageSrcs struct {
	html string // my-element-sk-demo.html
	ts   string // my-element-sk-demo.ts
	scss string // my-element-sk-demo.scss
}

// isValid returns true if the struct contains the necessary sources to build an sk_page, or false
// otherwise.
func (p *skPageSrcs) isValid() bool {
	return p.html != "" && p.ts != ""
}

// has returns true if the struct includes the given source file, or false otherwise.
func (p *skPageSrcs) has(src string) bool {
	return src == p.html || src == p.ts || src == p.scss
}

// generateSkElementRule generates a sk_element rule for the given sources.
func generateSkElementRule(name string, srcs *skElementSrcs, dir string) (*rule.Rule, common.ImportsParsedFromRuleSources) {
	tsSrcs := []string{srcs.ts}
	if srcs.indexTs != "" {
		tsSrcs = append(tsSrcs, srcs.indexTs)
		sort.Strings(tsSrcs)
	}

	rule := rule.NewRule("sk_element", name)
	rule.SetAttr("ts_srcs", tsSrcs)
	if srcs.scss != "" {
		rule.SetAttr("sass_srcs", []string{srcs.scss})
	}
	rule.SetAttr("visibility", []string{"//visibility:public"})

	imports := &importsParsedFromRuleSourcesImpl{}
	for _, tsSrc := range tsSrcs {
		imports.tsImports = append(imports.tsImports, extractImportsFromTypeScriptFile(filepath.Join(dir, tsSrc))...)
	}
	if srcs.scss != "" {
		imports.sassImports = extractImportsFromSassFile(filepath.Join(dir, srcs.scss))
	}

	return rule, imports
}

// generateSkPageRule generates a sk_page rule for the given sources.
func generateSkPageRule(srcs *skPageSrcs, dir string) (*rule.Rule, common.ImportsParsedFromRuleSources) {
	rule := rule.NewRule("sk_page", makeRuleNameFromFileName(srcs.html, ""))
	rule.SetAttr("html_file", srcs.html)
	rule.SetAttr("ts_entry_point", srcs.ts)
	if srcs.scss != "" {
		rule.SetAttr("scss_entry_point", srcs.scss)
	}

	imports := &importsParsedFromRuleSourcesImpl{
		tsImports: extractImportsFromTypeScriptFile(filepath.Join(dir, srcs.ts)),
	}
	if srcs.scss != "" {
		imports.sassImports = extractImportsFromSassFile(filepath.Join(dir, srcs.scss))
	}

	return rule, imports
}

// generateSkDemoPageServerRule generates a sk_demo_page_server rule for the given sk_page.
func generateSkDemoPageServerRule(skPage label.Label) (*rule.Rule, common.ImportsParsedFromRuleSources) {
	rule := rule.NewRule("sk_demo_page_server", "demo_page_server")
	rule.SetAttr("sk_page", skPage.String())
	return rule, &importsParsedFromRuleSourcesImpl{}
}

// generateSassLibraryRule generates a sass_library rule for the given Sass file.
func generateSassLibraryRule(file, dir string) (*rule.Rule, common.ImportsParsedFromRuleSources) {
	rule := rule.NewRule("sass_library", makeRuleNameFromFileName(file, "_sass_lib"))
	rule.SetAttr("srcs", []string{file})
	rule.SetAttr("visibility", []string{"//visibility:public"})
	return rule, &importsParsedFromRuleSourcesImpl{sassImports: extractImportsFromSassFile(filepath.Join(dir, file))}
}

// generateKarmaTestRule generates a karma_test rule for the given TypeScript file.
func generateKarmaTestRule(file, dir string) (*rule.Rule, common.ImportsParsedFromRuleSources) {
	rule := rule.NewRule("karma_test", makeRuleNameFromFileName(file, ""))
	rule.SetAttr("src", file)
	return rule, &importsParsedFromRuleSourcesImpl{tsImports: extractImportsFromTypeScriptFile(filepath.Join(dir, file))}
}

// generateNodeJSTestRule generates a nodejs_test rule for the given TypeScript file.
func generateNodeJSTestRule(file, dir string) (*rule.Rule, common.ImportsParsedFromRuleSources) {
	rule := rule.NewRule("nodejs_test", makeRuleNameFromFileName(file, ""))
	rule.SetAttr("src", file)
	return rule, &importsParsedFromRuleSourcesImpl{tsImports: extractImportsFromTypeScriptFile(filepath.Join(dir, file))}
}

// generateSkElementPuppeteerTestRule generates a sk_element_puppeteer_test rule for the given
// TypeScript file and sk_demo_page_server.
func generateSkElementPuppeteerTestRule(file, dir string, skDemoPageServer label.Label) (*rule.Rule, common.ImportsParsedFromRuleSources) {
	rule := rule.NewRule("sk_element_puppeteer_test", makeRuleNameFromFileName(file, ""))
	rule.SetAttr("src", file)
	rule.SetAttr("sk_demo_page_server", skDemoPageServer.String())
	return rule, &importsParsedFromRuleSourcesImpl{tsImports: extractImportsFromTypeScriptFile(filepath.Join(dir, file))}
}

// generateTSLibraryRule generates a ts_library rule for the given TypeScript file.
func generateTSLibraryRule(file, dir string) (*rule.Rule, common.ImportsParsedFromRuleSources) {
	rule := rule.NewRule("ts_library", makeRuleNameFromFileName(file, "_ts_lib"))
	rule.SetAttr("srcs", []string{file})
	rule.SetAttr("visibility", []string{"//visibility:public"})
	return rule, &importsParsedFromRuleSourcesImpl{tsImports: extractImportsFromTypeScriptFile(filepath.Join(dir, file))}
}

// makeRuleNameFromFileName returns e.g. "baz_ts_lib" when given "foo/bar/baz.ts" and "_ts_lib".
func makeRuleNameFromFileName(file, suffix string) string {
	file = strings.ToLower(path.Base(file))
	return strings.TrimSuffix(file, filepath.Ext(file)) + suffix
}

// extractImportsFromSassFile returns the verbatim paths of the import statements found in the given
// Sass file.
func extractImportsFromSassFile(path string) []string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicf("Error reading file %q: %v", path, err)
	}
	return parsers.ParseSassImports(string(b[:]))
}

// extractImportsFromTypeScriptFile returns the verbatim paths of the import statements found in the
// given TypeScript file.
func extractImportsFromTypeScriptFile(path string) []string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicf("Error reading file %q: %v", path, err)
	}

	// Ignore CSS / Sass imports from TypeScript files (Webpack idiom).
	var imports []string
	for _, imp := range parsers.ParseTSImports(string(b[:])) {
		if !strings.HasSuffix(imp, ".css") && !strings.HasSuffix(imp, ".scss") {
			imports = append(imports, imp)
		}
	}
	return imports
}

// generateEmptyRules returns a list of rules that cannot be built with the files found in the
// directory, for example because a file in its srcs argument does not exist anymore.
//
// Gazelle will merge these rules with the existing rules, and if any of their attributes marked as
// non-empty are empty after the merge, they will be deleted.
func generateEmptyRules(args language.GenerateArgs) []*rule.Rule {
	var emptyRules []*rule.Rule

	// If no BUILD.bazel file exists in the current directory, there's nothing to do.
	if args.File == nil {
		return emptyRules
	}

	allFilesInDir := map[string]bool{}
	for _, f := range append(args.RegularFiles, args.GenFiles...) {
		allFilesInDir[f] = true
	}

	someFilesFound := func(files ...string) bool {
		for _, f := range files {
			if allFilesInDir[f] {
				return true
			}
		}
		return false
	}

	allFilesFound := func(files ...string) bool {
		for _, f := range files {
			if !allFilesInDir[f] {
				return false
			}
		}
		return true
	}

	allRulesByNameInDir := map[string]*rule.Rule{}
	for _, r := range args.File.Rules {
		allRulesByNameInDir[r.Name()] = r
	}

	parseRelLabelFromAttribute := func(r *rule.Rule, attr string) string {
		if r.AttrString(attr) == "" {
			return ""
		}
		l, err := label.Parse(r.AttrString(attr))
		if err != nil {
			log.Panicf(`Unable to parse attribute %q of rule %q: %v`, attr, r.Name(), err)
		}
		// We assume the label is always relative, e.g. ":foo", not "//path/to:foo".
		if !l.Relative {
			log.Panicf(`Label in attribute %q of rule %q should be relative, but was %q`, attr, r.Name(), l.String())
		}
		return l.Name
	}

	ruleFound := func(kind, name string) bool {
		r := allRulesByNameInDir[name]
		return r != nil && r.Kind() == kind
	}

	isEmptyRule := func(kind, name string) bool {
		for _, r := range emptyRules {
			if r.Kind() == kind && r.Name() == name {
				return true
			}
		}
		return false
	}

	// An existing sk_demo_page_server rule is empty (i.e. should be deleted) if:
	//
	//   1) its associated sk_page rule no longer exists,
	//   or
	//   2) if its associated sk_page exists, but is empty (i.e. should be deleted), e.g. because its
	//      source files no longer exist.
	//
	// Similarly, an existing sk_element_puppeteer_test rule is empty (i.e. should be deleted) if:
	//
	//   1) its associated sk_demo_page_server rule no longer exists,
	//   or
	//   2) if its associated sk_demo_page_server rule exists, but is empty (i.e. should be deleted),
	//      e.g. because its associated sk_page rule no longer exists, or is empty.
	//
	// To address condition 2) of both of the above rules, we populate the emptyRules slice in the
	// following order:
	//
	//   - All rules except for sk_demo_page_server and sk_element_puppeteer_test rules.
	//   - All sk_demo_page_server rules.
	//   - All sk_element_puppeteer_test rules.
	//
	// This will allow us to query the emptyRules slice in the loop below for any empty sk_page or
	// sk_demo_page_server rules when processing sk_demo_page_server or sk_element_puppeteer_test
	// rules, respectively.

	var allRules, skDemoPageServerRules, skElementPuppeteerTestRules []*rule.Rule
	for _, r := range args.File.Rules {
		switch r.Kind() {
		case "sk_demo_page_server":
			skDemoPageServerRules = append(skDemoPageServerRules, r)
		case "sk_element_puppeteer_test":
			skElementPuppeteerTestRules = append(skElementPuppeteerTestRules, r)
		default:
			allRules = append(allRules, r)
		}
	}
	allRules = append(allRules, skDemoPageServerRules...)
	allRules = append(allRules, skElementPuppeteerTestRules...)

	for _, curRule := range allRules {
		var empty bool

		switch curRule.Kind() {
		case "karma_test":
			empty = !someFilesFound(curRule.AttrString("src"))
		case "nodejs_test":
			empty = !someFilesFound(curRule.AttrString("src"))
		case "sass_library":
			empty = !someFilesFound(curRule.AttrStrings("srcs")...)
		case "sk_demo_page_server":
			skPage := parseRelLabelFromAttribute(curRule, "sk_page")
			empty = !ruleFound("sk_page", skPage) || isEmptyRule("sk_page", skPage)
		case "sk_element":
			empty = !someFilesFound(curRule.AttrStrings("ts_srcs")...)
		case "sk_element_puppeteer_test":
			skDemoPageServer := parseRelLabelFromAttribute(curRule, "sk_demo_page_server")
			empty = !allFilesFound(curRule.AttrString("src")) || !ruleFound("sk_demo_page_server", skDemoPageServer) || isEmptyRule("sk_demo_page_server", skDemoPageServer)
		case "sk_page":
			empty = !allFilesFound(curRule.AttrString("html_file"), curRule.AttrString("ts_entry_point"))
		case "ts_library":
			empty = !someFilesFound(curRule.AttrStrings("srcs")...)
		}

		if empty {
			emptyRules = append(emptyRules, rule.NewRule(curRule.Kind(), curRule.Name()))
		}
	}

	return emptyRules
}

// Fix implements the language.Language interface.
func (l *Language) Fix(c *config.Config, f *rule.File) {}

var _ language.Language = &Language{}