package shared

import "regexp"

// Options from args mapppped here
type Options struct {
	// The path to the file to folder to zip - resolves to abs path
	Path string

	// The path to area to output - resolves to abs path
	OutPath string

	// Regular expression for include only specific file names
	IncludeFiles *regexp.Regexp

	// Regular expression to only include specific folders
	IncludeFolders *regexp.Regexp

	// The command to run
	Cmd string
}
