package runner

import (
	"os"

	"github.com/UmbrellaCrow612/go-zip/cli/shared"
	"github.com/UmbrellaCrow612/go-zip/cli/utils"
)

// Runs the main loop for the given cmd and it's options
func Run(options *shared.Options) {
	switch options.Cmd {
	case "zip":
		RunZipCmd(options)
	case "unzip":
		RunUnZipCmd(options)
	default:
		utils.PrintStderr("Error: unknown command: " + options.Cmd + ". Valid commands are 'zip' or 'unzip'.")
		os.Exit(1)
	}
}
