package args

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/UmbrellaCrow612/go-zip/cli/shared"
	"github.com/UmbrellaCrow612/go-zip/cli/utils"
)

// Parses args array and gets all flag options
func Parse() *shared.Options {
	opts := &shared.Options{}
	args := os.Args[1:]

	if len(args) < 3 {
		utils.PrintStderr("Error: missing required arguments. Usage: <cmd> <path> <outPath> [flags]")
		os.Exit(1)
	}

	// Assign mandatory args
	opts.Cmd = args[0]
	rawPath := args[1]
	rawOutPath := args[2]

	// Validate and resolve input path
	absPath, err := filepath.Abs(rawPath)
	if err != nil {
		utils.PrintStderr("Error resolving input path: " + err.Error())
		os.Exit(1)
	}
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		utils.PrintStderr("Error: input path does not exist: " + absPath)
		os.Exit(1)
	}
	opts.Path = absPath

	// Resolve output path to absolute
	absOutPath, err := filepath.Abs(rawOutPath)
	if err != nil {
		utils.PrintStderr("Error resolving output path: " + err.Error())
		os.Exit(1)
	}
	opts.OutPath = absOutPath

	// Parse optional flags
	for _, arg := range args[3:] {
		switch {
		case strings.HasPrefix(arg, "--include-files="):
			pattern := strings.TrimPrefix(arg, "--include-files=")
			re, err := regexp.Compile(pattern)
			if err != nil {
				utils.PrintStderr("Error compiling exclude regex: " + err.Error())
				os.Exit(1)
			}
			opts.IncludeFiles = re
		case strings.HasPrefix(arg, "--include-folders="):
			pattern := strings.TrimPrefix(arg, "--include-folders=")
			re, err := regexp.Compile(pattern)
			if err != nil {
				utils.PrintStderr("Error compiling exclude regex: " + err.Error())
				os.Exit(1)
			}
			opts.IncludeFolders = re
		default:
			utils.PrintStderr("Error: unknown flag: " + arg)
			os.Exit(1)
		}
	}

	return opts
}
