package cmd

import (
	"code-profiles/code"
	"code-profiles/config"
	"code-profiles/utils"

	"github.com/spf13/cobra"
)

var (
	profileName string
	installFlag bool

	openCmd = &cobra.Command{
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

			base_p, err := config.BaseProfile()
			base_p.Path = p.Path

			if installFlag {
				if err == nil {
					code.InstallExtensions(base_p)
				}
				code.InstallExtensions(p)
			}

			code.LaunchCode(p)
		},
	}
)

func Execute() {
	var rootCmd = &cobra.Command{Use: "code-profiles"}

	openCmd.Flags().BoolVarP(&installFlag, "install", "i", false, "should install extensions before opening vscode")

	rootCmd.AddCommand(openCmd)

	utils.Check(rootCmd.Execute())
}
