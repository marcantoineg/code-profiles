package cmd

import (
	"code-profiles/code"
	"code-profiles/config"
	"code-profiles/utils"

	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open [profile_name]",
	Short: "open VSCode using a custom profile for extensions",
	Run: func(cmd *cobra.Command, args []string) {
		var p config.Profile
		var err error
		if len(args) < 1 {
			p, err = config.GetProfileFromFile()
		} else {
			p, err = config.GetProfile(args[0])
		}
		utils.Check(err)

		if installFlag {
			code.InstallExtensions(p, verboseFlag)
		}

		code.LaunchCode(p, verboseFlag)
	},
}
