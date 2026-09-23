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
	slog.Debug("Resolved additional snippets file", "file", additionalSnipsFile)
	if additionalSnipsFile == "" {
		s := snippets{}
		return &s, nil
	}

	if _, err := os.Stat(additionalSnipsFile); os.IsNotExist(err) {
		slog.Debug("Additional snippets file does not exist", "file", additionalSnipsFile, ErrAttr(err))
		return &snips, nil
	} else if err != nil {
		slog.Debug("Failed to check additional snippets file", "file", additionalSnipsFile, ErrAttr(err))
		return nil, err
	}

	contents, err := os.ReadFile(additionalSnipsFile)
	if err != nil {
		slog.Debug("Failed to read additional snippets file", "file", additionalSnipsFile, ErrAttr(err))
		return nil, err
	}
	slog.Debug("Read additional snippets", "contents", string(contents))

	return SnippetsFromToml(string(contents))
}
