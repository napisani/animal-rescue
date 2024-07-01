package main
//  this file is from the https://github.com/knqyf263/pet repository.

import (
	"github.com/pelletier/go-toml"
)

// Config is a struct of config
type Config struct {
	General GeneralConfig
	Gist    GistConfig
	GitLab  GitLabConfig
	GHEGist GHEGistConfig
}

// GeneralConfig is a struct of general config
type GeneralConfig struct {
	SnippetFile string
	Editor      string
	Column      int
	SelectCmd   string
	Backend     string
	SortBy      string
	Color       bool
	Cmd         []string
}

// GistConfig is a struct of config for Gist
type GistConfig struct {
	FileName    string `toml:"file_name"`
	AccessToken string `toml:"access_token"`
	GistID      string `toml:"gist_id"`
	Public      bool
	AutoSync    bool `toml:"auto_sync"`
}

// GitLabConfig is a struct of config for GitLabSnippet
type GitLabConfig struct {
	FileName    string `toml:"file_name"`
	AccessToken string `toml:"access_token"`
	Url         string
	ID          string
	Visibility  string
	AutoSync    bool `toml:"auto_sync"`
	SkipSsl     bool `toml:"skip_ssl"`
}

// GHEGistConfig is a struct of config for Gist of Github Enterprise
type GHEGistConfig struct {
	BaseUrl     string `toml:"base_url"`
	UploadUrl   string `toml:"upload_url"`
	FileName    string `toml:"file_name"`
	AccessToken string `toml:"access_token"`
	GistID      string `toml:"gist_id"`
	Public      bool
	AutoSync    bool `toml:"auto_sync"`
}

// Flag is global flag variable
var Flag FlagConfig

// FlagConfig is a struct of flag
type FlagConfig struct {
	Debug        bool
	Query        string
	FilterTag    string
	Command      bool
	Delimiter    string
	OneLine      bool
	Color        bool
	Tag          bool
	UseMultiLine bool
	UseEditor    bool
}

func ConfigFromToml(data string) (*Config, error) {
	c := &Config{}
	err := toml.Unmarshal([]byte(data), c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Config) ToToml() (string, error) {
  b, err := toml.Marshal(c)
  if err != nil {
    return "", err
  }
  return string(b), nil
}
