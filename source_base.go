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
		slog.Debug("Snippet file does not exist: %v %v", f, err)
		return &snips, nil 
	} else if err != nil {
		slog.Debug("Error checking snippet file: %v %v", f, err)
		return nil, err
	}

	contents, err := os.ReadFile(f)
	if err != nil {
    slog.Debug("Error reading snippet file: %v %v", f, err)
		return nil, err
	}
  slog.Debug("contents: %v", string(contents))

	return SnippetsFromToml(string(contents))
}


