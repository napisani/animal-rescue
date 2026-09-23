package main

import (
	"log/slog"
	"os"
)

type SourcePetLocal struct {
}

var fileVariations = []string{"pet-snippet.toml", ".pet-snippet.toml"}

func (s *SourcePetLocal) GetSnippets(opts *GetSnippetsOptions) (*snippets, error) {
	snips := snippets{}
	f := FindFileVariation(opts.Cwd, fileVariations)

	if _, err := os.Stat(f); os.IsNotExist(err) {
		slog.Debug("Snippet file does not exist", "file", f, ErrAttr(err))
		return &snips, nil
	} else if err != nil {
		slog.Debug("Failed to check snippet file", "file", f, ErrAttr(err))
		return nil, err
	}

	contents, err := os.ReadFile(f)
	if err != nil {
		slog.Debug("Failed to read snippet file", "file", f, ErrAttr(err))
		return nil, err
	}
	slog.Debug("Read snippets", "contents", string(contents))

	return SnippetsFromToml(string(contents))
}
