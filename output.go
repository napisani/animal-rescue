package main

import (
	"log/slog"
	"os"
	"path"
	"path/filepath"
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

func WriteTempConfigFile(config *Config) (string, error) {
	t, err := config.ToToml()
	if err != nil {
		return "", err
	}

	slog.Debug("Writing config", "file", configPath)
	if err := writeFileAtomically(configPath, []byte(t)); err != nil {
		return "", err
	}
	return configPath, nil
}

func WriteTempSnippetsFile(snips *snippets) (string, error) {
	t, err := snips.ToToml()
	if err != nil {
		return "", err
	}

	slog.Debug("Writing snippets", "file", snippetPath)
	if err := writeFileAtomically(snippetPath, []byte(t)); err != nil {
		return "", err
	}
	return snippetPath, nil
}

func writeFileAtomically(filename string, contents []byte) (err error) {
	dir := filepath.Dir(filename)
	file, err := os.CreateTemp(dir, "."+filepath.Base(filename)+".*")
	if err != nil {
		return err
	}
	tempName := file.Name()
	defer func() {
		if file != nil {
			_ = file.Close()
		}
		_ = os.Remove(tempName)
	}()

	if err := file.Chmod(0600); err != nil {
		return err
	}
	if _, err := file.Write(contents); err != nil {
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	file = nil
	return os.Rename(tempName, filename)
}
