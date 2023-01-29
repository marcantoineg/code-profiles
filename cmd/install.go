package cmd

import (
	"code-profiles/code"
	"code-profiles/config"
	"code-profiles/utils"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [profile_name]",
	Short: "install required VSCode extensions for a given profile",
	Run: func(cmd *cobra.Command, args []string) {
		var p config.Profile
		var err error
		if len(args) < 1 {
			p, err = config.GetProfileFromFile()
		} else {
			p, err = config.GetProfile(args[0])
		}
		utils.Check(err)

		code.InstallExtensions(p, verboseFlag)
	},
}
