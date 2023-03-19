// Package cmd exposes function to interact with the cobra command
package cmd

import (
	"code-profiles/utils"

	"github.com/spf13/cobra"
)

var (
	configPathFlag string
	verboseFlag    bool
	installFlag    bool
)

// Execute runs the main cobra command for cli `code-profiles`.
func Execute() {
	var rootCmd = &cobra.Command{Use: "code-profiles"}

	openCmd.Flags().StringVarP(&configPathFlag, "config-path", "c", "./code-profiles.yml", "Path to code-profiles config")
	openCmd.Flags().BoolVarP(&installFlag, "install", "i", false, "should install extensions before opening VSCode")
	openCmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "prints additional logs")

	installCmd.Flags().StringVarP(&configPathFlag, "config-path", "c", "./code-profiles.yml", "Path to code-profiles config")
	installCmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "prints additional logs")

	rootCmd.AddCommand(openCmd)
	rootCmd.AddCommand(installCmd)

	utils.Check(rootCmd.Execute())
}
