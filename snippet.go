package main

import (
	"encoding/json"

	"github.com/pelletier/go-toml"
)

type snippets struct {
	Snippets []Snippet `json:"snippets" toml:"snippets"`
}

type Snippet struct {
	Description string   `json:"description" toml:"description"`
	Command     string   `json:"command" toml:"command"`
	Tag         []string `json:"tag" toml:"tag"`
	Output      string   `json:"output" toml:"output"`
}

func (s *snippets) ToToml() (string, error) {
	b, err := toml.Marshal(s)
	if err != nil {
		return "", err

	}
	return string(b), nil
}

func (s *snippets) ToJson() (string, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func SnippetsFromToml(data string) (*snippets, error) {
	s := &snippets{}
	err := toml.Unmarshal([]byte(data), s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
