// The presubmit binary runs many checks on the differences between the current commit and its
// parent branch. This is usually origin/main. If there is a chain of CLs (i.e. branches), this
// will only consider the diffs between the current commit and the parent branch.
// If all presubmits pass, this binary will output nothing (by default) and have exit code 0.
// If any presubmits fail, there will be errors logged to stdout and the exit code will be non-zero.
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var (
		upstream = flag.String("upstream", "origin/main", "The upstream repo to diff against.")
		verbose  = flag.Bool("verbose", false, "If extra logging is desired")
	)
	flag.Parse()
	ctx := withOutputWriter(context.Background(), os.Stdout)

	branchBaseCommit := findBranchBase(ctx, *upstream)
	if branchBaseCommit == "" {
		// This either means the user hasn't committed their changes or is just on the main branch
		// somewhere. Either way, we don't want to run the presubmits. It could mutate those
		// un-committed changes or just be a no-op, since presumably code in the past passed the
		// presubmit checks.
		logf(ctx, "No commits since %s. Presubmit passes by default.\nDid you forget to commit changes?\n", *upstream)
		os.Exit(0)
	}
	if *verbose {
		logf(ctx, "Base commit is %s\n", branchBaseCommit)
	}

	changedFiles, deletedFiles := computeDiffFiles(ctx, branchBaseCommit)
	if *verbose {
		logf(ctx, "Changed files:\n%v\n", changedFiles)
		logf(ctx, "Deleted files:\n%s\n", deletedFiles)
	}
	ok := checkLongLines(ctx, changedFiles)
	ok = ok && checkTODOHasOwner(ctx, changedFiles)
	ok = ok && checkForStrayWhitespace(ctx, changedFiles)
	ok = ok && checkHasNoTabs(ctx, changedFiles)
	ok = ok && checkBannedGoAPIs(ctx, changedFiles)
	ok = ok && checkJSDebugging(ctx, changedFiles)
	ok = ok && checkNonASCII(ctx, changedFiles)

	if ok {
		os.Exit(0)
	}
	logf(ctx, "Presubmit errors detected!\n")
	os.Exit(1)
}

const (
	gitErrorMessage = `Error running git - is git on your path?

If running with Bazel, you need to invoke this like:
bazel run //cmd/presubmit --run_under="cd $PWD &&"
`
	refSeperator = "$|" // Some string we hope that no users start their branch names with
)

// findBranchBase returns the git commit of the parent branch. If there is a chain of CLs
// (i.e. branches), this will return the parent branch's most recent commit. If we are on the
// upstream branch, this will return empty string. Otherwise, it will return the commit on the
// upstream branch where branching occurred. It shells out to git, which is presumed to be on
// PATH in order to find this information.
func findBranchBase(ctx context.Context, upstream string) string {
	// rev-list is one of the git "plumbing" commands and the output should be relatively stable.
	// https://mirrors.edge.kernel.org/pub/software/scm/git/docs/git.html#_low_level_commands_plumbing
	cmd := exec.CommandContext(ctx, "git", "rev-list", "HEAD", "^"+upstream,
		// %D means "ref names", which is the commit hash and any branch name associated with it
		// %p means the shortened parent hash
		`--format=%D`+refSeperator+`%p`)
	output, err := cmd.CombinedOutput()
	if err != nil {
		logf(ctx, string(output)+"\n")
		logf(ctx, err.Error()+"\n")
		panic(gitErrorMessage)
	}
	return parseRevList(string(output))
}

// parseRevList looks for the most recent branch, as indicated by the "ref name". Failing to find
// that, it will return the last parent commit, which will connect to the upstream branch.
func parseRevList(output string) string {
	output = strings.TrimSpace(output)
	if output == "" {
		return ""
	}
	lines := strings.Split(output, "\n")
	lastCommit := ""
	lastParentCommit := ""
	for _, line := range lines {
		if strings.HasPrefix(line, "commit ") {
			lastCommit = strings.TrimPrefix(line, "commit ")
			continue
		}
		if strings.Contains(line, refSeperator) {
			parts := strings.Split(line, refSeperator)
			lastParentCommit = parts[1]
			if parts[0] == "" {
				continue
			}
		}
		if strings.HasPrefix(line, "HEAD -> ") {
			continue
		}
		// This means we found a branch name that the current branch is dependent on.
		return lastCommit
	}
	// We got to the end without finding another branch, so that means this commit must be based
	// directly on the main branch. The last parent shows us where in that main branch we are
	// based.
	return lastParentCommit
}

// computeDiffFiles returns a slice of changed (modified or added) files with the lines touched
// and a slice of deleted files. It shells out to git, which is presumed to be on PATH in order
// to find this information.
func computeDiffFiles(ctx context.Context, branchBase string) ([]fileWithChanges, []string) {
	// git diff-index is considered to be a "git plumbing" API, so its output should be pretty
	// stable across git version (unlike ordinary git diff, which can do things like show
	// tabs as multiple spaces).
	cmd := exec.CommandContext(ctx, "git", "diff-index", branchBase,
		// Don't show any surrounding context (i.e. lines which did not change)
		"--unified=0",
		"--patch-with-raw",
		"--no-color")
	output, err := cmd.CombinedOutput()
	if err != nil {
		logf(ctx, string(output)+"\n")
		logf(ctx, err.Error()+"\n")
		panic(gitErrorMessage)
	}
	return parseGitDiff(string(output))
}

var (
	gitDiffLine = regexp.MustCompile(`^diff --git (?P<fileA>.*) (?P<fileB>.*)$`)
	lineAnchor  = regexp.MustCompile(`^@@ -(?P<deleted>\d+)(?P<delLines>,\d+)? \+(?P<added>\d+)(?P<addLines>,\d+)? @@.*$`)
)

const (
	deletedFilePrefix = "deleted file mode"
	addedLinePrefix   = "+"
)

// fileWithChanges represents an added or modified files along with the lines changed (touched).
// Lines that were deleted are not tracked.
type fileWithChanges struct {
	fileName     string
	touchedLines []lineOfCode
}

func (f fileWithChanges) String() string {
	rv := f.fileName + "\n"
	for _, line := range f.touchedLines {
		rv += "  " + line.String() + "\n"
	}
	return rv
}

type lineOfCode struct {
	contents string
	num      int
}

func (c lineOfCode) String() string {
	return fmt.Sprintf("% 4d:%s", c.num, c.contents)
}

// parseGitDiff looks through the provided `git diff` output and finds the files that were
// added or modified, as well as the new version of any lines touched. It also returns
// a slice of deleted files.
func parseGitDiff(diffOutput string) ([]fileWithChanges, []string) {
	var changed []fileWithChanges
	var deleted []string
	lines := strings.Split(diffOutput, "\n")
	currFileDeleted := false
	lastLineIndex := -1
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if match := gitDiffLine.FindStringSubmatch(line); len(match) > 0 {
			// A new file was changed, reset our counters, trackers, and add an entry for it.
			var newFile fileWithChanges
			newFile.fileName = strings.TrimPrefix(match[2], "b/")
			changed = append(changed, newFile)
			currFileDeleted = false
			lastLineIndex = -1
		}
		if currFileDeleted || len(changed) == 0 {
			continue
		}
		currFile := &changed[len(changed)-1]
		if strings.HasPrefix(line, deletedFilePrefix) {
			deleted = append(deleted, currFile.fileName)
			changed = changed[:len(changed)-1] // trim off changed file
			currFileDeleted = true
			continue
		}
		if match := lineAnchor.FindStringSubmatch(line); len(match) > 0 {
			// The added group has the line number corresponding to the next line starting with a +
			var err error
			lastLineIndex, err = strconv.Atoi(match[3])
			if err != nil {
				panic("Got an integer where none was expected: " + line + "\n" + err.Error())
			}
			continue
		}
		if lastLineIndex < 0 {
			// Have not found a line index yet, ignore the lines like:
			// ++ b/PRESUBMIT.py
			continue
		}
		if strings.HasPrefix(line, addedLinePrefix) {
			currFile.touchedLines = append(currFile.touchedLines, lineOfCode{
				contents: strings.TrimPrefix(line, addedLinePrefix),
				num:      lastLineIndex,
			})
			lastLineIndex++
		}
	}

	return changed, deleted
}

// checkLongLines looks through all touched lines and returns false if any of them (not covered by
// exceptions) have lines longer than 100 lines, an arbitrary measurement.
// Based on https://chromium.googlesource.com/chromium/tools/depot_tools.git/+/19b3eb5adbe00e9da40375cb5dc47380a46f3041/presubmit_canned_checks.py#488
func checkLongLines(ctx context.Context, files []fileWithChanges) bool {
	const maxLineLength = 100
	ignoreFileExts := []string{".go", ".html", ".py"}
	ignoreFiles := []string{"package-lock.json", "go.sum", "infra/bots/tasks.json"}
	ok := true
	for _, f := range files {
		if contains(ignoreFiles, f.fileName) {
			continue
		}
		if contains(ignoreFileExts, filepath.Ext(f.fileName)) {
			continue
		}
		for _, line := range f.touchedLines {
			if len(line.contents) > maxLineLength {
				logf(ctx, "%s:%d Line too long (%d/%d)\n", f.fileName, line.num, len(line.contents), maxLineLength)
				ok = false
			}
		}
	}
	return ok
}

var todoWithoutOwner = regexp.MustCompile(`TODO[^(]`)

// checkTODOHasOwner looks through all touched lines and returns false if any of them (not covered
// by exceptions) have a TODO without an owner or a bug.
// Based on https://chromium.googlesource.com/chromium/tools/depot_tools.git/+/19b3eb5adbe00e9da40375cb5dc47380a46f3041/presubmit_canned_checks.py#464
func checkTODOHasOwner(ctx context.Context, files []fileWithChanges) bool {
	ignoreFiles := []string{
		// These files have TODO in their function names and test data.
		"cmd/presubmit/presubmit.go", "cmd/presubmit/presubmit_test.go",
	}
	ok := true
	for _, f := range files {
		if contains(ignoreFiles, f.fileName) {
			continue
		}
		for _, line := range f.touchedLines {
			if todoWithoutOwner.MatchString(line.contents) || strings.HasSuffix(line.contents, "TODO") {
				logf(ctx, "%s:%d TODO without owner or bug\n", f.fileName, line.num)
				ok = false
			}
		}
	}
	return ok
}

var trailingWhitespace = regexp.MustCompile(`\s+$`)

// checkForStrayWhitespace goes through all touched lines and returns false if any of them end with
// a whitespace character (e.g. tabs, spaces). newlines should have been stripped out of
// the content of a line earlier.
// Based on https://chromium.googlesource.com/chromium/tools/depot_tools.git/+/19b3eb5adbe00e9da40375cb5dc47380a46f3041/presubmit_canned_checks.py#476
func checkForStrayWhitespace(ctx context.Context, files []fileWithChanges) bool {
	ok := true
	for _, f := range files {
		for _, line := range f.touchedLines {
			if trailingWhitespace.MatchString(line.contents) {
				logf(ctx, "%s:%d Trailing whitespace\n", f.fileName, line.num)
				ok = false
			}
		}
	}
	return ok
}

// checkHasNoTabs goes through all touched lines and returns false if any of them (barring
// exceptions for Golang and Makefiles) have tabs anywhere.
// Based on https://chromium.googlesource.com/chromium/tools/depot_tools.git/+/19b3eb5adbe00e9da40375cb5dc47380a46f3041/presubmit_canned_checks.py#441
func checkHasNoTabs(ctx context.Context, files []fileWithChanges) bool {
	ignoreFileExts := []string{".go", ".mk"}
	ignoreFiles := []string{"Makefile"}
	ok := true
	for _, f := range files {
		if contains(ignoreFiles, f.fileName) || contains(ignoreFiles, filepath.Base(f.fileName)) {
			continue
		}
		if contains(ignoreFileExts, filepath.Ext(f.fileName)) {
			continue
		}
		for _, line := range f.touchedLines {
			if strings.Contains(line.contents, "\t") {
				logf(ctx, "%s:%d Tab character not allowed\n", f.fileName, line.num)
				ok = false
			}
		}
	}
	return ok
}

type bannedGoAPI struct {
	regex      *regexp.Regexp
	suggestion string
	exceptions []*regexp.Regexp
}

// checkBannedGoAPIs goes through all touched lines in go files and returns false if any of them
// have APIs that we wish not to use. It logs suggested replacements in that case.
func checkBannedGoAPIs(ctx context.Context, files []fileWithChanges) bool {
	bannedAPIs := []bannedGoAPI{
		{regex: regexp.MustCompile(`reflect\.DeepEqual`), suggestion: "DeepEqual in go.skia.org/infra/go/testutils"},
		{regex: regexp.MustCompile(`github\.com/golang/glog`), suggestion: "go.skia.org/infra/go/sklog"},
		{regex: regexp.MustCompile(`github\.com/skia-dev/glog`), suggestion: "go.skia.org/infra/go/sklog"},
		{regex: regexp.MustCompile(`http\.Get`), suggestion: "NewTimeoutClient in go.skia.org/infra/go/httputils"},
		{regex: regexp.MustCompile(`http\.Head`), suggestion: "NewTimeoutClient in go.skia.org/infra/go/httputils"},
		{regex: regexp.MustCompile(`http\.Post`), suggestion: "NewTimeoutClient in go.skia.org/infra/go/httputils"},
		{regex: regexp.MustCompile(`http\.PostForm`), suggestion: "NewTimeoutClient in go.skia.org/infra/go/httputils"},
		{regex: regexp.MustCompile(`os\.Interrupt`), suggestion: "AtExit in go.skia.org/go/cleanup"},
		{regex: regexp.MustCompile(`signal\.Notify`), suggestion: "AtExit in go.skia.org/go/cleanup"},
		{regex: regexp.MustCompile(`syscall\.SIGINT`), suggestion: "AtExit in go.skia.org/go/cleanup"},
		{regex: regexp.MustCompile(`syscall\.SIGTERM`), suggestion: "AtExit in go.skia.org/go/cleanup"},
		{regex: regexp.MustCompile(`syncmap\.Map`), suggestion: "sync.Map, added in go 1.9"},
		{regex: regexp.MustCompile(`assert\s+"github\.com/stretchr/testify/require"`), suggestion: `non-aliased import; this can be confused with package "github.com/stretchr/testify/assert"`},
		{
			regex:      regexp.MustCompile(`"git"`),
			suggestion: `Executable in go.skia.org/infra/go/git`,
			exceptions: []*regexp.Regexp{
				// These don't actually shell out to git; the tests look for "git" in the
				// command line and mock stdout accordingly.
				regexp.MustCompile(`autoroll/go/repo_manager/.*_test.go`),
				// This doesn't shell out to git; it's referring to a CIPD package with
				// the same name.
				regexp.MustCompile(`infra/bots/gen_tasks.go`),
				// This doesn't shell out to git; it retrieves the path to the Git binary
				// in the corresponding Bazel-downloaded CIPD packages.
				regexp.MustCompile(`bazel/external/cipd/git/git.go`),
				// Our presubmits invoke git directly because git is a necessary
				// executable for all devs, and we do not want our presubmit code to
				// depend on the code it is checking.
				regexp.MustCompile(`cmd/presubmit/.*`),
				// This is the one place where we are allowed to shell out to git; all
				// others should go through here.
				regexp.MustCompile(`go/git/git_common/.*.go`),
			},
		},
	}
	ok := true
	for _, f := range files {
		if filepath.Ext(f.fileName) != ".go" {
			continue
		}
		if f.fileName == "cmd/presubmit/presubmit_test.go" {
			// We don't want our own test cases to trigger any of these.
			continue
		}
		for _, line := range f.touchedLines {
		bannedAPILoop:
			for _, bannedAPI := range bannedAPIs {
				for _, exception := range bannedAPI.exceptions {
					if exception.MatchString(f.fileName) {
						continue bannedAPILoop
					}
				}
				if match := bannedAPI.regex.FindStringSubmatch(line.contents); len(match) > 0 {
					logf(ctx, "%s:%d Instead of %s, please use %s\n", f.fileName, line.num, match[0], bannedAPI.suggestion)
					ok = false
				}
			}
		}
	}
	return ok
}

// checkJSDebugging goes through all touched lines and returns false if any TS or JS files contain
// refinements of debugging that we don't want to check in.
func checkJSDebugging(ctx context.Context, files []fileWithChanges) bool {
	debuggingCalls := []string{"debugger;", "it.only(", "describe.only("}
	targetFileExts := []string{".ts", ".js"}
	ok := true
	for _, f := range files {
		if !contains(targetFileExts, filepath.Ext(f.fileName)) {
			continue
		}
		for _, line := range f.touchedLines {
			for _, call := range debuggingCalls {
				if strings.Contains(line.contents, call) {
					logf(ctx, "%s:%d debugging code found (%s)\n", f.fileName, line.num, call)
					ok = false
				}
			}
		}
	}
	return ok
}

// checkNonASCII goes through all touched lines and returns false if any of them contain non-ASCII
// characters (except for file formats that support things like UTF-8).
func checkNonASCII(ctx context.Context, files []fileWithChanges) bool {
	// This list can grow if other file extensions are OK with non-ascii (UTF-8) characters
	ignoreFileExts := []string{".go"}
	ok := true
	for _, f := range files {
		if contains(ignoreFileExts, filepath.Ext(f.fileName)) {
			continue
		}
		for _, line := range f.touchedLines {
			// https://stackoverflow.com/a/53069799/1447621
			for i := 0; i < len(line.contents); i++ {
				if line.contents[i] > '\u007F' { // unicode.MaxASCII
					// Report both line number and (1-indexed) byte offset
					logf(ctx, "%s:%d:%d Non ASCII character found\n", f.fileName, line.num, i+1)
					ok = false
					break
				}
			}
		}
	}
	return ok
}

// contains returns true if the given slice has the provided element in it.
func contains[T string](a []T, s T) bool {
	for _, x := range a {
		if x == s {
			return true
		}
	}
	return false
}

type contextKeyType string

const outputWriterKey contextKeyType = "outputWriter"

// withOutputWriter registers the given writer to the context. See also logf.
func withOutputWriter(ctx context.Context, w io.Writer) context.Context {
	return context.WithValue(ctx, outputWriterKey, w)
}

// logf takes the writer on the context and writes the given format and arguments to it. This allows
// unit tests to intercept logged output if necessary. It panics if a writer was not registered
// using withOutputWriter.
func logf(ctx context.Context, format string, args ...interface{}) {
	w, ok := ctx.Value(outputWriterKey).(io.Writer)
	if !ok {
		panic("Must set outputWriter on ctx")
	}
	_, err := fmt.Fprintf(w, format, args...)
	if err != nil {
		panic("Error while logging " + err.Error())
	}
}