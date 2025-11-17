package main

import (
	"github.com/UmbrellaCrow612/go-zip/cli/args"
	"github.com/UmbrellaCrow612/go-zip/cli/runner"
)

func main() {
	options := args.Parse()
	runner.Run(options)
}
