package main

import (
	"log/slog"
	"os"
)

type SourcePetLocal struct {
}
var fileVariations = []string{"pet-snippet.toml", ".pet-snippet.toml" }

func (s *SourcePetLocal) GetSnippets(opts *GetSnippetsOptions) (*snippets, error) {
	snips := snippets{}
  f := FindFileVariation(opts.Cwd, fileVariations)  

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


