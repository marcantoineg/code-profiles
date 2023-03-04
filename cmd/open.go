package cmd

import (
	"code-profiles/code"
	"code-profiles/config"
	"code-profiles/utils"

	"github.com/marcantoineg/fileutil"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open [profile_name]",
	Short: "open VSCode using a custom profile for extensions",
	Run: func(cmd *cobra.Command, args []string) {
		config.SetConfigPath(fileutil.ReplaceTilde(configPathFlag))

		var profileName = ""
		if len(args) >= 1 {
			profileName = args[0]
		}

		p, err := config.GetProfile(profileName)
		utils.Check(err)

		if installFlag {
			code.InstallExtensions(p, verboseFlag)
		}

		code.LaunchCode(p, verboseFlag)
	},
}
