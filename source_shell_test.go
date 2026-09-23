package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestSourceShellIsOptIn(t *testing.T) {
	source := &SourceShell{}
	got, err := source.GetSnippets(&GetSnippetsOptions{})
	if err != nil {
		t.Fatalf("GetSnippets() error = %v", err)
	}
	if len(got.Snippets) != 0 {
		t.Fatalf("GetSnippets() returned %d snippets without a shell path", len(got.Snippets))
	}
}

func TestSourceShellReadsAnnotatedAliasesAndFunctions(t *testing.T) {
	dir := t.TempDir()
	writeTestFile(t, filepath.Join(dir, "10-first.sh"), `# pet: List files with details
alias ll='ls -alF'

# This unrelated comment breaks the annotation.
# pet: ignored

# pet: Run the current project's tests
run-tests() {
  make test
}

# pet: Function declaration with the keyword form
function deploy() {
  ./deploy
}
`)
	writeTestFile(t, filepath.Join(dir, "20-second.sh"), `# pet: Last definition wins
alias ll='ls -la'

# An annotation must be directly associated with a declaration.
# pet: not a snippet
export VALUE=1
`)

	got, err := (&SourceShell{}).GetSnippets(&GetSnippetsOptions{ShellPaths: []string{dir}})
	if err != nil {
		t.Fatalf("GetSnippets() error = %v", err)
	}

	want := []Snippet{
		{Description: "Last definition wins", Command: "ll", Tag: []string{"shell", "alias"}},
		{Description: "Run the current project's tests", Command: "run-tests", Tag: []string{"shell", "function"}},
		{Description: "Function declaration with the keyword form", Command: "deploy", Tag: []string{"shell", "function"}},
	}
	if !reflect.DeepEqual(got.Snippets, want) {
		t.Fatalf("snippets = %#v, want %#v", got.Snippets, want)
	}
}

func TestSourceShellReadsSingleFileAndExpandsHome(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "aliases.sh")
	t.Setenv("HOME", dir)
	writeTestFile(t, file, `# pet: Say hello
hello() { echo hello; }
`)

	got, err := (&SourceShell{}).GetSnippets(&GetSnippetsOptions{ShellPaths: []string{"~/aliases.sh"}})
	if err != nil {
		t.Fatalf("GetSnippets() error = %v", err)
	}
	if len(got.Snippets) != 1 || got.Snippets[0].Command != "hello" {
		t.Fatalf("snippets = %#v, want one hello snippet", got.Snippets)
	}
}

func writeTestFile(t *testing.T, filename, contents string) {
	t.Helper()
	if err := os.WriteFile(filename, []byte(contents), 0644); err != nil {
		t.Fatalf("WriteFile(%q) error = %v", filename, err)
	}
}
