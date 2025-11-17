package shared

import "regexp"

// Options from args mapppped here
type Options struct {
	// The path to the file to folder to zip - resolves to abs path
	Path string

	// The path to area to output - resolves to abs path
	OutPath string

	// New name for the unziped folder or file if empty not provided
	Name string

	// If it should recursive copy
	Recursive bool

	// Regular expression for exclude files or folders or nil im not provided
	Exclude *regexp.Regexp

	// The command to run
	Cmd string
}
