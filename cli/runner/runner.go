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
		err := RunZipCmd(options)
		if err != nil {
			utils.PrintStderr(err.Error())
			os.Exit(1)
		}
	case "unzip":
		err := RunUnZipCmd(options)
		if err != nil {
			utils.PrintStderr(err.Error())
			os.Exit(1)
		}
	default:
		utils.PrintStderr("Error: unknown command: " + options.Cmd + ". Valid commands are 'zip' or 'unzip'.")
		os.Exit(1)
	}
}
