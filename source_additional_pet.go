package main

import (
	"log/slog"
	"os"
)

type SourcePetAdditional struct {
}

func (s *SourcePetAdditional) GetSnippets(opts *GetSnippetsOptions) (*snippets, error) {
	snips := snippets{}
	additionalSnipsFile := os.Getenv("PET_ADDL_SNIPPETS")
	if additionalSnipsFile != "" {
		s := snippets{}
		return &s, nil
	}

	if _, err := os.Stat(additionalSnipsFile); os.IsNotExist(err) {
		slog.Debug("Additional Snippet file does not exist: %v %v", additionalSnipsFile, err)
		return &snips, nil
	} else if err != nil {
		slog.Debug("Error checking additional snippet file: %v %v", additionalSnipsFile, err)
		return nil, err
	}

	contents, err := os.ReadFile(additionalSnipsFile)
	if err != nil {
		slog.Debug("Error reading snippet additional file: %v %v", additionalSnipsFile, err)
		return nil, err
	}
	slog.Debug("contents: %v", string(contents))

	return SnippetsFromToml(string(contents))
}
