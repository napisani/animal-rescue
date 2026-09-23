package main

import (
	"log/slog"
	"os"
)

type SourceBase struct {
}

func (s *SourceBase) GetSnippets(opts *GetSnippetsOptions) (*snippets, error) {
	snips := snippets{}

	f := ExpandPath(opts.InputConfig.General.SnippetFile)

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
