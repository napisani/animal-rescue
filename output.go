package main

import (
	"log/slog"
	"os"
	"path"
)

var snippetPath = path.Join(os.TempDir(), "pet-snippets.toml")
var configPath = path.Join(os.TempDir(), "pet-config.toml")

func DeleteTempSnippetsFile() error {
  slog.Debug("Deleting temp snippets file")
	err := os.Remove(snippetPath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

func DeleteTempConfigFile() error {
  slog.Debug("Deleting temp config file")
  err := os.Remove(configPath)
  if err != nil && !os.IsNotExist(err) {
    return err
  }
  return nil
}

func WriteTempConfigFile(config *Config) (string ,error) {
  f, err := CreateFileIfNotExist(configPath)
  if err != nil {
    return "", err
  }

  defer f.Close()
  slog.Debug("Writing config to %v", f.Name())

  t, err := config.ToToml()
  if err != nil {
    return "", err
  }

  _, err = f.WriteString(t)
  if err != nil {
    return "", err
  }

  return f.Name(), nil
}

func WriteTempSnippetsFile(snips *snippets) (string, error) {

	f, err := CreateFileIfNotExist(snippetPath)

	if err != nil {
		return "", err
	}

	defer f.Close()
	slog.Debug("Writing snippets to %v", f.Name())

	t, err := snips.ToToml()
	if err != nil {
		return "", err
	}

	_, err = f.WriteString(t)
	if err != nil {
		return "", err
	}

	return f.Name(), nil
}
