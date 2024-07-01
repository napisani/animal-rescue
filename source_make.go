package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
)

type SourceMake struct {
}

func (s *SourceMake) GetSnippets(opts *GetSnippetsOptions) (*snippets, error) {
	snips := snippets{}
	// results of `make list`
	if !commandExists("make") {
		slog.Debug("make command not found in PATH")
		return &snips, nil
	}
	_, err := os.Stat(path.Join(opts.Cwd, "Makefile"))
	if os.IsNotExist(err) {
		slog.Debug("Makefile does not exist in cwd")
		return &snips, nil
	}

	cmd := exec.Command("make", "-qp")
	output, err := cmd.Output()
	slog.Debug("make err %v", ErrAttr(err))

  if output == nil  && err != nil {
    // only return an error if no output was returned
    // make seems to return a non-zero exit code when printing in question-mode
    return &snips, err
  }

	for _, line := range strings.Split(string(output), "\n") {
		reg := regexp.MustCompile("^[a-zA-Z0-9][^$#\\/\\t=]*:([^=]|$)")
		if reg.MatchString(line) {
			idxColon := strings.Index(line, ":")
			if idxColon == -1 {
				continue
			}
			line := line[:idxColon]
			if line == "Makefile" {
				continue
			}
			slog.Debug("line: %s", line)
			snips.Snippets = append(snips.Snippets, SnippetFromMakeLine(line))
		}
	}
	return &snips, nil
}

func SnippetFromMakeLine(line string) Snippet {
	s := Snippet{}
	s.Description = fmt.Sprintf("Makefile:%s", line)
	s.Tag = []string{"make", "Makefile"}
	s.Command = fmt.Sprintf("make %s", line)
	return s
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
