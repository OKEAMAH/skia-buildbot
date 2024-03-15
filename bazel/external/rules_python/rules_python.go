package rules_python

import (
	"os/exec"
	"path/filepath"
	"runtime"

	"go.skia.org/infra/bazel/go/bazel"
	"go.skia.org/infra/go/skerr"
)

// FindPython3 returns the path to the `python3` binary provided by rules_python[1].
//
// Calling this function from any Go package will automatically establish a Bazel dependency on the
// corresponding external Bazel repository.
//
// [1] https://github.com/bazelbuild/rules_python
func FindPython3() (string, error) {
	if !bazel.InBazel() {
		return exec.LookPath("python3")
	}
	if runtime.GOOS == "linux" {
		// This path was determined by looking at the `interpreter` constant defined in
		// @python3_10//:defs.bzl, e.g.:
		//
		//     $ cat $(bazel info output_base)/external/python3_10/defs.bzl
		//     # Generated by python/repositories.bzl
		//     host_platform = "x86_64-unknown-linux-gnu"
		//     interpreter = "@python3_10_x86_64-unknown-linux-gnu//:bin/python3"
		return filepath.Join(bazel.RunfilesDir(), "external", "python3_10_x86_64-unknown-linux-gnu", "bin", "python3"), nil
	}
	return "", skerr.Fmt("unsupported runtime.GOOS: %q", runtime.GOOS)
}
