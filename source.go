package main

type GetSnippetsOptions struct {
	Cwd         string
	InputConfig *Config
	ShellPaths  []string
}

type SnippetSource interface {
	GetSnippets(*GetSnippetsOptions) (*snippets, error)
}
