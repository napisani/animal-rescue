package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var (
	shellPetComment      = regexp.MustCompile(`^\s*#\s*pet:\s*(\S.*)$`)
	shellAlias           = regexp.MustCompile(`^\s*alias\s+([A-Za-z_][A-Za-z0-9_-]*)\s*=`)
	shellFunction        = regexp.MustCompile(`^\s*([A-Za-z_][A-Za-z0-9_-]*)\s*\(\s*\)\s*\{?`)
	shellKeywordFunction = regexp.MustCompile(`^\s*function\s+([A-Za-z_][A-Za-z0-9_-]*)\s*(?:\(\s*\))?\s*\{?`)
)

type SourceShell struct{}

func (s *SourceShell) GetSnippets(opts *GetSnippetsOptions) (*snippets, error) {
	result := &snippets{}
	if len(opts.ShellPaths) == 0 {
		return result, nil
	}

	seen := make(map[string]int)
	for _, shellPath := range opts.ShellPaths {
		files, err := shellFiles(ExpandPath(shellPath))
		if err != nil {
			return nil, err
		}
		for _, file := range files {
			found, err := snippetsFromShellFile(file)
			if err != nil {
				return nil, err
			}
			for _, snippet := range found {
				if index, ok := seen[snippet.Command]; ok {
					result.Snippets[index] = snippet
					continue
				}
				seen[snippet.Command] = len(result.Snippets)
				result.Snippets = append(result.Snippets, snippet)
			}
		}
	}

	return result, nil
}

func shellFiles(shellPath string) ([]string, error) {
	info, err := os.Stat(shellPath)
	if err != nil {
		return nil, fmt.Errorf("stat shell path %q: %w", shellPath, err)
	}
	if !info.IsDir() {
		return []string{shellPath}, nil
	}

	entries, err := os.ReadDir(shellPath)
	if err != nil {
		return nil, fmt.Errorf("read shell path %q: %w", shellPath, err)
	}

	files := make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		files = append(files, filepath.Join(shellPath, entry.Name()))
	}
	sort.Strings(files)
	return files, nil
}

func snippetsFromShellFile(filename string) ([]Snippet, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("open shell file %q: %w", filename, err)
	}
	defer file.Close()

	var result []Snippet
	var description string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if match := shellPetComment.FindStringSubmatch(line); match != nil {
			description = strings.TrimSpace(match[1])
			continue
		}
		if strings.TrimSpace(line) == "" {
			continue
		}

		name, kind := shellDeclaration(line)
		if name == "" {
			description = ""
			continue
		}
		if description != "" {
			result = append(result, Snippet{
				Description: description,
				Command:     name,
				Tag:         []string{"shell", kind},
			})
		}
		description = ""
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read shell file %q: %w", filename, err)
	}
	return result, nil
}

func shellDeclaration(line string) (name, kind string) {
	if match := shellAlias.FindStringSubmatch(line); match != nil {
		return match[1], "alias"
	}
	if match := shellKeywordFunction.FindStringSubmatch(line); match != nil {
		return match[1], "function"
	}
	if match := shellFunction.FindStringSubmatch(line); match != nil {
		return match[1], "function"
	}
	return "", ""
}
