"""This module defines the google_chrome repository rule.

The google_chrome repository rule installs Google Chrome and the required fonts and fontconfig
settings to properly render webpages in Puppeteer-based[1] screenshot tests.

This rule does not install any of the runtime libraries required by Chrome or Chromium. Runtime
libraries must be provided by the host system, or by the RBE toolchain container image when
building on RBE. When building binaries with Bazel, it is important to hermetically provide all
required libraries, as this ensures reproducible builds. However, we only use Chrome and Chromium
(via Puppeteer) for tests, so exact library versions are less of a concern. This has the additional
benefit of keeping this rule small, as both Chrome and Chromium depend on a large number of runtime
libraries.

Puppeteer tests must set the FONTCONFIG_SYSROOT[2] environment variable to point to the repository
directory generated by this rule. This ensures that the correct fonts are used in screenshot tests.

Note that Puppeteer bundles its own Chromium binary, and therefore ignores the Chrome binary
provided by this rule. However, said binary is used by Karma tests.

Why is this rule necessary, when the browser_repositories[3] repository rule exists? The Chromium
binary provided by the browser_repositories repository rule is indeed sufficient for Karma tests
(see example here[4] and here[5]). However, it does not provide the fonts or fontconfig settings
needed to properly render webpages in Puppeteer-based screenshot tests. Given that the
google_chrome repository rule downloads a Debian package with the necessary fonts, downloading an
additional package to provide the Chrome binary seems marginally less complex than adding a
dependency on an additional external Bazel repository.

[1] https://pptr.dev
[2] https://www.freedesktop.org/software/fontconfig/fontconfig-user.html
[3] https://github.com/bazelbuild/rules_webtesting/blob/e9cf17123068b1123c68219edf9b274bf057b9cc/web/versioned/browsers-0.3.3.bzl#L18
[4] https://github.com/google/skia/blob/89fea08f6b16a73b3f824e37dbf1039bb08bf91d/WORKSPACE.bazel#L156-L174
[5] https://github.com/google/skia/blob/89fea08f6b16a73b3f824e37dbf1039bb08bf91d/bazel/karma_test.bzl#L75
"""

load("//bazel:gcs_mirror.bzl", "gcs_mirror_url")

_DEB_PACKAGES_LINUX_AMD64 = [
    # Google Chrome's debian repository is found at https://dl.google.com/linux/chrome/deb. The URL
    # and SHA256 hash below were taken on 2022-05-25 from the Release[1] file in the repository's
    # stable channel. Said file is found at
    # https://dl.google.com/linux/chrome/deb/dists/stable/Release.
    #
    # [1] https://wiki.debian.org/DebianRepository/Format
    {
        "url": "https://dl.google.com/linux/chrome/deb/pool/main/g/google-chrome-stable/google-chrome-stable_102.0.5005.61-1_amd64.deb",
        "sha256": "fb8386ad122b328c7166647ae7deeb14d452cd8b66b01c51c1abccd4c6441680",
    },

    # https://packages.debian.org/buster/all/fonts-liberation/download
    {
        "url": "https://ftp.debian.org/debian/pool/main/f/fonts-liberation/fonts-liberation_1.07.4-9_all.deb",
        "sha256": "c936aebbfd0af7851399ae5ab08bb01744f5e3381f7678fb87cc77114f95ef53",
    },
]

def _google_chrome_impl(repository_ctx):
    is_linux = repository_ctx.os.name.lower().startswith("linux")

    if not is_linux:
        # Support for other operating systems can be added as needed.
        fail("OS not yet supported: %s." % repository_ctx.os.name)

    # Download and extract Debian packages.
    repository_ctx.report_progress("Downloading and installing .deb packages.")
    for deb in _DEB_PACKAGES_LINUX_AMD64:
        repository_ctx.download_and_extract(
            url = gcs_mirror_url(url = deb["url"], sha256 = deb["sha256"]),
            output = "__tmp__",
            sha256 = deb["sha256"],
        )
        repository_ctx.extract(
            archive = "__tmp__/data.tar.xz",
            output = ".",
        )
        repository_ctx.delete("__tmp__")

    # Generate /etc/fonts/fonts.conf file.
    #
    # An alternative to providing our own /etc/fonts/fonts.conf file is to install the
    # fontconfig[1] and fontconfig-config[2] Debian packages, which provide good default settings.
    # These packages provide an /etc/fonts/fonts.conf file and a number of
    # /etc/fonts/conf.d/*.conf files which are loaded from the former via an
    # <include>conf.d</include> element. However, for some unknown reason (a bug, perhaps?) the
    # fontconfig library fails to scan the /etc/fonts/conf.d directory when the
    # FONTCONFIG_SYSROOT[3] environment variable is set (as is the case for Puppeteer and Karma
    # tests). Thus, as a workaround, we provide our own /etc/fonts/fonts.conf file with a minimal
    # set of settings based on the config files provided by the fontconfig-config Debian package.
    # This is enough to guarantee good quality Puppeteer screenshots.
    #
    # The below file maps common font families, such as Times New Roman, Arial and Courier New, to
    # their generic names, such as serif, sans-serif, and monospace, respectively. Then, it maps
    # serif, sans-serif and monospace to Liberation Sans, Liberation Serif, and Liberation Mono,
    # respectively. Finally, if an unknown font is requested, it defaults to sans-serif, i.e.
    # Liberation Sans.
    #
    # [1] https://packages.debian.org/buster/fontconfig
    # [2] https://packages.debian.org/buster/fontconfig-config
    # [3] https://www.freedesktop.org/software/fontconfig/fontconfig-user.html
    repository_ctx.file("etc/fonts/fonts.conf", content = """<?xml version="1.0"?>
<!DOCTYPE fontconfig SYSTEM "fonts.dtd">
<fontconfig>
    <dir>/usr/share/fonts</dir>
    <cachedir>/var/cache/fontconfig</cachedir>

    <!--
    Mark common families with their generics so we'll get something reasonable.

    Based on
    https://gitlab.freedesktop.org/fontconfig/fontconfig/-/blob/d863f6778915f7dd224c98c814247ec292904e30/conf.d/45-latin.conf.
    -->

    <!--
    Serif faces
    -->
    <alias>
        <family>Bitstream Vera Serif</family>
        <default><family>serif</family></default>
    </alias>
    <alias>
        <family>Cambria</family>
        <default><family>serif</family></default>
    </alias>
    <alias>
        <family>Constantia</family>
        <default><family>serif</family></default>
    </alias>
    <alias>
        <family>DejaVu Serif</family>
        <default><family>serif</family></default>
    </alias>
    <alias>
        <family>Elephant</family>
        <default><family>serif</family></default>
    </alias>
    <alias>
        <family>Garamond</family>
        <default><family>serif</family></default>
    </alias>
    <alias>
        <family>Georgia</family>
        <default><family>serif</family></default>
    </alias>
    <alias>
        <family>Liberation Serif</family>
        <default><family>serif</family></default>
    </alias>
    <alias>
        <family>Luxi Serif</family>
        <default><family>serif</family></default>
    </alias>
    <alias>
        <family>MS Serif</family>
        <default><family>serif</family></default>
    </alias>
    <alias>
        <family>Nimbus Roman No9 L</family>
        <default><family>serif</family></default>
    </alias>
    <alias>
        <family>Nimbus Roman</family>
        <default><family>serif</family></default>
    </alias>
    <alias>
        <family>Palatino Linotype</family>
        <default><family>serif</family></default>
    </alias>
    <alias>
        <family>Thorndale AMT</family>
        <default><family>serif</family></default>
    </alias>
    <alias>
        <family>Thorndale</family>
        <default><family>serif</family></default>
    </alias>
    <alias>
        <family>Times New Roman</family>
        <default><family>serif</family></default>
    </alias>
    <alias>
        <family>Times</family>
        <default><family>serif</family></default>
    </alias>
    <!--
    Sans-serif faces
    -->
    <alias>
        <family>Albany AMT</family>
        <default><family>sans-serif</family></default>
    </alias>
    <alias>
        <family>Albany</family>
        <default><family>sans-serif</family></default>
    </alias>
    <alias>
        <family>Arial Unicode MS</family>
        <default><family>sans-serif</family></default>
    </alias>
    <alias>
        <family>Arial</family>
        <default><family>sans-serif</family></default>
    </alias>
    <alias>
        <family>Bitstream Vera Sans</family>
        <default><family>sans-serif</family></default>
    </alias>
    <alias>
        <family>Britannic</family>
        <default><family>sans-serif</family></default>
    </alias>
    <alias>
        <family>Calibri</family>
        <default><family>sans-serif</family></default>
    </alias>
    <alias>
        <family>Candara</family>
        <default><family>sans-serif</family></default>
    </alias>
    <alias>
        <family>Century Gothic</family>
        <default><family>sans-serif</family></default>
    </alias>
    <alias>
        <family>Corbel</family>
        <default><family>sans-serif</family></default>
    </alias>
    <alias>
        <family>DejaVu Sans</family>
        <default><family>sans-serif</family></default>
    </alias>
    <alias>
        <family>Helvetica</family>
        <default><family>sans-serif</family></default>
    </alias>
    <alias>
        <family>Haettenschweiler</family>
        <default><family>sans-serif</family></default>
    </alias>
    <alias>
        <family>Liberation Sans</family>
        <default><family>sans-serif</family></default>
    </alias>
    <alias>
        <family>MS Sans Serif</family>
        <default><family>sans-serif</family></default>
    </alias>
    <alias>
        <family>Nimbus Sans L</family>
        <default><family>sans-serif</family></default>
    </alias>
    <alias>
        <family>Nimbus Sans</family>
        <default><family>sans-serif</family></default>
    </alias>
    <alias>
        <family>Luxi Sans</family>
        <default><family>sans-serif</family></default>
    </alias>
    <alias>
        <family>Tahoma</family>
        <default><family>sans-serif</family></default>
    </alias>
    <alias>
        <family>Trebuchet MS</family>
        <default><family>sans-serif</family></default>
    </alias>
    <alias>
        <family>Twentieth Century</family>
        <default><family>sans-serif</family></default>
    </alias>
    <alias>
        <family>Verdana</family>
        <default><family>sans-serif</family></default>
    </alias>
    <!--
    Monospace faces
    -->
    <alias>
        <family>Andale Mono</family>
        <default><family>monospace</family></default>
    </alias>
     <alias>
        <family>Bitstream Vera Sans Mono</family>
        <default><family>monospace</family></default>
    </alias>
    <alias>
        <family>Consolas</family>
        <default><family>monospace</family></default>
    </alias>
    <alias>
        <family>Courier New</family>
        <default><family>monospace</family></default>
    </alias>
    <alias>
        <family>Courier</family>
        <default><family>monospace</family></default>
    </alias>
    <alias>
        <family>Cumberland AMT</family>
        <default><family>monospace</family></default>
    </alias>
    <alias>
        <family>Cumberland</family>
        <default><family>monospace</family></default>
    </alias>
    <alias>
        <family>DejaVu Sans Mono</family>
        <default><family>monospace</family></default>
    </alias>
    <alias>
        <family>Fixedsys</family>
        <default><family>monospace</family></default>
    </alias>
    <alias>
        <family>Inconsolata</family>
        <default><family>monospace</family></default>
    </alias>
    <alias>
        <family>Liberation Mono</family>
        <default><family>monospace</family></default>
    </alias>
    <alias>
        <family>Luxi Mono</family>
        <default><family>monospace</family></default>
    </alias>
    <alias>
        <family>Nimbus Mono L</family>
        <default><family>monospace</family></default>
    </alias>
    <alias>
        <family>Nimbus Mono</family>
        <default><family>monospace</family></default>
    </alias>
    <alias>
        <family>Nimbus Mono PS</family>
        <default><family>monospace</family></default>
    </alias>
    <alias>
        <family>Terminal</family>
        <default><family>monospace</family></default>
    </alias>

    <!--
    If the font still has no generic name, add sans-serif.

    Based on
    https://gitlab.freedesktop.org/fontconfig/fontconfig/-/blob/d863f6778915f7dd224c98c814247ec292904e30/conf.d/49-sansserif.conf.
    -->
    <match target="pattern">
        <test qual="all" name="family" compare="not_eq">
            <string>sans-serif</string>
        </test>
        <test qual="all" name="family" compare="not_eq">
            <string>serif</string>
        </test>
        <test qual="all" name="family" compare="not_eq">
            <string>monospace</string>
        </test>
        <edit name="family" mode="append_last">
            <string>sans-serif</string>
        </edit>
    </match>

    <!--
    Default to the Liberation font family.
    -->
    <alias>
        <family>sans-serif</family>
        <prefer>
            <family>Liberation Sans</family>
        </prefer>
    </alias>
    <alias>
        <family>monospace</family>
        <prefer>
            <family>Liberation Mono</family>
        </prefer>
    </alias>
    <alias>
        <family>serif</family>
        <prefer>
            <family>Liberation Serif</family>
        </prefer>
    </alias>
</fontconfig>
""")

    # Generate BUILD.bazel file.
    repository_ctx.file("BUILD.bazel", content = """
filegroup(
    name = "all_files",
    srcs = glob(
        include = ["**/*"],
    ),
    visibility = ["//visibility:public"],
)
""")

google_chrome = repository_rule(
    implementation = _google_chrome_impl,
    doc = "Hermetically installs Google Chrome, and all required libraries and fonts.",
)