package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSourceMakeUsesRequestedSearchPath(t *testing.T) {
	targetDir := t.TempDir()
	otherDir := t.TempDir()
	writeTestFile(t, filepath.Join(targetDir, "Makefile"), "target-project:\n\t@true\n")
	writeTestFile(t, filepath.Join(otherDir, "Makefile"), "other-project:\n\t@true\n")

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(otherDir); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(originalDir); err != nil {
			t.Errorf("restore working directory: %v", err)
		}
	})

	got, err := (&SourceMake{}).GetSnippets(&GetSnippetsOptions{Cwd: targetDir})
	if err != nil {
		t.Fatalf("GetSnippets() error = %v", err)
	}

	commands := make(map[string]bool)
	for _, snippet := range got.Snippets {
		commands[snippet.Command] = true
	}
	if !commands["make target-project"] {
		t.Fatalf("commands = %#v, want target project command", commands)
	}
	if commands["make other-project"] {
		t.Fatalf("commands = %#v, unexpectedly used process working directory", commands)
	}
}
