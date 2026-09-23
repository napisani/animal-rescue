package main

import (
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func ErrAttr(err error) slog.Attr {
	return slog.Any("error", err)
}

func ExpandPath(s string) string {
	if len(s) >= 2 && s[0] == '~' && os.IsPathSeparator(s[1]) {
		if runtime.GOOS == "windows" {
			s = filepath.Join(os.Getenv("USERPROFILE"), s[2:])
		} else {
			s = filepath.Join(os.Getenv("HOME"), s[2:])
		}
	}
	return os.Expand(s, os.Getenv)
}

func DoesFileExist(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func DoesFileExistAtRoot(cwd string, filename string) bool {
	root := FindProjectRoot(cwd)
	if root == "" {
		return false
	}
	return DoesFileExist(path.Join(root, filename))
}

func FindFileVariation(cwd string, fileVariations []string) string {
	for _, f := range fileVariations {
		f = ExpandPath(path.Join(cwd, f))
		if DoesFileExist(f) {
			return f
		}
	}
	return ""
}

func FindProjectRoot(cwd string) string {
	for {
		slog.Debug("Checking project root", "directory", cwd)
		if DoesFileExist(path.Join(cwd, ".git")) {
			return cwd
		}
		if cwd == "/" {
			return ""
		}
		cwd = path.Dir(cwd)
	}
}
