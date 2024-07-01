package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path"
)

type Args struct {
	InputConfig string
	Clean       bool
	Debug       bool
}

var tmpFile = path.Join(os.TempDir(), ".animal-rescue.json.tmp")

func parseArgs() Args {
	var args Args
	flag.StringVar(&args.InputConfig, "config", "", "Input config file")
	flag.BoolVar(&args.Clean, "clean", false, "Clean any cached `pet` config files")
	flag.BoolVar(&args.Debug, "debug", false, "debug")
	flag.Parse()
	return args
}

func main() {
	args := parseArgs()
	if args.Debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	err := DeleteTempSnippetsFile()
	if err != nil {
		slog.Error("Failed to delete temp snippets file %v", ErrAttr(err))
		panic(err)
	}
	err = DeleteTempConfigFile()
	if err != nil {
		slog.Error("Failed to delete temp config file %v", ErrAttr(err))
		panic(err)
	}
	if args.Clean {
		return
	}

	if args.InputConfig == "" {
		slog.Error("Input config file is required")
		os.Exit(1)
	}

	cwd, err := os.Getwd()
	if err != nil {
		slog.Error("Failed to get current working directory %v", ErrAttr(err))
		panic(err)
	}

	inputConfigPath := args.InputConfig
	contents, err := os.ReadFile(inputConfigPath)
	if err != nil {
		slog.Error("Failed to read input config file %v", ErrAttr(err))
		panic(err)
	}

	inputConfig, err := ConfigFromToml(string(contents))
	if err != nil {
		slog.Error("Failed to parse input config file %v", ErrAttr(err))
		panic(err)
	}

	opts := GetSnippetsOptions{
		Cwd:         cwd,
		InputConfig: inputConfig,
	}

	allSnips := snippets{}
	sources := []SnippetSource{&SourceBase{}, &SourceMake{}, &SourceNpm{}, &SourcePetLocal{}}

	for _, src := range sources {
		snips, err := src.GetSnippets(&opts)
		if err != nil {
			slog.Error("Failed to get snippets %v", ErrAttr(err))
			panic(err)
		}
		allSnips.Snippets = append(allSnips.Snippets, snips.Snippets...)
	}

	tempSnipsFile, err := WriteTempSnippetsFile(&allSnips)
	if err != nil {
		slog.Error("Failed to write temp snippets file %v", ErrAttr(err))
		panic(err)
	}

	inputConfig.General.SnippetFile = tempSnipsFile

	tempConfigFile, err := WriteTempConfigFile(inputConfig)
	if err != nil {
		slog.Error("Failed to write temp config file %v", ErrAttr(err))
		panic(err)
	}

	fmt.Println(tempConfigFile)
}
