package main

import (
	"github.com/pelletier/go-toml"
)

type snippets struct {
  Snippets []Snippet `json:"snippets"`
}

type Snippet struct {
	Description string   `json:"description"`
	Command     string   `json:"command"`
	Tag         []string `json:"tag"`
	Output      string   `json:"output"`
}



func (s *snippets) ToToml() (string, error) {
	b, err := toml.Marshal(s)
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
