package cmd

import (
	"code-profiles/code"
	"code-profiles/config"
	"code-profiles/utils"

	"github.com/marcantoineg/fileutil"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [profile_name]",
	Short: "install required VSCode extensions for a given profile",
	Run: func(cmd *cobra.Command, args []string) {
		config.SetConfigPath(fileutil.ReplaceTilde(configPathFlag))

		var profileName = ""
		if len(args) >= 1 {
			profileName = args[0]
		}

		p, err := config.GetProfile(profileName)
		utils.Check(err)

		code.InstallExtensions(p, verboseFlag)
	},
}
