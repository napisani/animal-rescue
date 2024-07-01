package main

type GetSnippetsOptions struct {
  Cwd string
  InputConfig *Config 
} 



type SnippetSource interface {
  GetSnippets(*GetSnippetsOptions) (*snippets, error)
}
