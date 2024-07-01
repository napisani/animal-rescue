package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
)

type SourceNpm struct {
}

type PackageJson struct {
	Scripts map[string]string `json:"scripts"`
}

func (s *SourceNpm) GetSnippets(opts *GetSnippetsOptions) (*snippets, error) {
	snips := snippets{}
	f := FindFileVariation(opts.Cwd, []string{"package.json"})
	if f == "" {
		return &snips, nil
	}

	isPnpm := DoesFileExistAtRoot(opts.Cwd, "pnpm-lock.yaml")

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
	p := PackageJson{}
	err = json.Unmarshal(contents, &p)
	if err != nil {
		slog.Debug("Error unmarshalling package.json: %v", err)
		return nil, err
	}
	return SnippetsFromPackageJson(p, isPnpm), nil
}

func SnippetsFromPackageJson(p PackageJson, isPnpm bool) *snippets {
	cmd := "npm"
	if isPnpm {
		cmd = "pnpm"
	}
	snips := snippets{}
	for k, _ := range p.Scripts {
		snips.Snippets = append(snips.Snippets, Snippet{
			Description: fmt.Sprintf("package.json: %s run %s", cmd, k),
			Command:     fmt.Sprintf("%s run %s", cmd, k),
			Tag:         []string{cmd, "package.json", k},
		})
	}
	return &snips
}
