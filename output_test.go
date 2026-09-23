package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteTempConfigFileAtomicallyReplacesSymlinkWithPrivateFile(t *testing.T) {
	dir := t.TempDir()
	victim := filepath.Join(dir, "victim")
	if err := os.WriteFile(victim, []byte("do not overwrite"), 0600); err != nil {
		t.Fatal(err)
	}

	originalConfigPath := configPath
	configPath = filepath.Join(dir, "pet-config.toml")
	t.Cleanup(func() { configPath = originalConfigPath })
	if err := os.Symlink(victim, configPath); err != nil {
		t.Fatal(err)
	}

	config := &Config{Gist: GistConfig{AccessToken: "review-secret"}}
	gotPath, err := WriteTempConfigFile(config)
	if err != nil {
		t.Fatalf("WriteTempConfigFile() error = %v", err)
	}
	if gotPath != configPath {
		t.Fatalf("path = %q, want %q", gotPath, configPath)
	}

	victimContents, err := os.ReadFile(victim)
	if err != nil {
		t.Fatal(err)
	}
	if string(victimContents) != "do not overwrite" {
		t.Fatalf("victim contents = %q", victimContents)
	}

	info, err := os.Lstat(configPath)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode()&os.ModeSymlink != 0 {
		t.Fatal("generated config remained a symlink")
	}
	if info.Mode().Perm() != 0600 {
		t.Fatalf("mode = %04o, want 0600", info.Mode().Perm())
	}
}

func TestWriteTempSnippetsFileIsPrivate(t *testing.T) {
	dir := t.TempDir()
	originalSnippetPath := snippetPath
	snippetPath = filepath.Join(dir, "pet-snippets.toml")
	t.Cleanup(func() { snippetPath = originalSnippetPath })

	gotPath, err := WriteTempSnippetsFile(&snippets{})
	if err != nil {
		t.Fatalf("WriteTempSnippetsFile() error = %v", err)
	}
	info, err := os.Stat(gotPath)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm() != 0600 {
		t.Fatalf("mode = %04o, want 0600", info.Mode().Perm())
	}
}
