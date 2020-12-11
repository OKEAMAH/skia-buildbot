"""This module defines the karma_test rule."""

load("@infra-sk_npm//@bazel/typescript:index.bzl", "ts_library")
load("@infra-sk_npm//@bazel/rollup:index.bzl", "rollup_bundle")
load("@infra-sk_npm//karma:index.bzl", _generated_karma_test = "karma_test")

def karma_test(name, srcs, deps, entry_point = None):
    """Runs unit tests in a browser with Karma and the Mocha test runner.

    When executed with `bazel test`, a headless Chrome browser will be used. This supports testing
    multiple karma_test targets in parallel, and works on RBE.

    When executed with `bazel run`, it prints out a URL to stdout that can be opened in the browser,
    e.g. to debug the tests using the browser's developer tools. Source maps are generated.

    When executed with `ibazel test`, the test runner never exits, and tests will be rerun every
    time a source file is changed.

    When executed with `ibazel run`, it will act the same way as `bazel run`, but the tests will be
    rebuilt automatically when a source file changes. Reload the browser page to see the changes.

    Args:
      name: The name of the target.
      srcs: The *.ts test files.
      deps: The ts_library dependencies for the source files.
      entry_point: File in srcs to be used as the entry point to generate the JS bundle executed by
        the test runner. Optional if srcs contains only one file.
    """

    if len(srcs) > 1 and not entry_point:
        fail("An entry_point must be specified when srcs contains more than one file.")

    if entry_point and entry_point not in srcs:
        fail("The entry_point must be included in srcs.")

    if len(srcs) == 1:
        entry_point = srcs[0]

    ts_library(
        name = name + "_lib",
        srcs = srcs,
        deps = deps + [
            # Add common test dependencies for convenience.
            "@infra-sk_npm//@types/mocha",
            "@infra-sk_npm//@types/chai",
            "@infra-sk_npm//@types/sinon",
        ],
    )

    rollup_bundle(
        name = name + "_bundle",
        entry_point = entry_point,
        deps = [
            name + "_lib",
            "@infra-sk_npm//@rollup/plugin-node-resolve",
            "@infra-sk_npm//@rollup/plugin-commonjs",
            "@infra-sk_npm//rollup-plugin-sourcemaps",
        ],
        format = "umd",
        config_file = "//infra-sk:rollup.config.js",
    )

    # This rule is automatically generated by rules_nodejs from Karma's package.json file.
    _generated_karma_test(
        name = name,
        size = "large",
        data = [
            name + "_bundle",
            "//infra-sk/karma_test:karma.conf.js",
            "@infra-sk_npm//karma-chrome-launcher",
            "@infra-sk_npm//karma-sinon",
            "@infra-sk_npm//karma-mocha",
            "@infra-sk_npm//karma-chai",
            "@infra-sk_npm//karma-chai-dom",
            "@infra-sk_npm//karma-spec-reporter",
            "@infra-sk_npm//mocha",
        ],
        templated_args = [
            "start",
            "$(execpath //infra-sk/karma_test:karma.conf.js)",
            "$$(rlocation $(location %s_bundle))" % name,
        ],
        tags = [
            # Necessary for it to work with ibazel.
            "ibazel_notify_changes",
        ],
    )