// Package code is a go interface with VS Code's CLI: `code`.
// It offers various go functions to execute subcommands using `code`, mainly to manage/install extensions and launch code with `Profiles`
package code

import (
	"code-profiles/config"
	"code-profiles/utils"
	"fmt"
	"os/exec"
)

// InstallExtensions install all of profile's extensions and dependencies (DependsOn) to the profile's path.
func InstallExtensions(profile config.Profile) {
	println("installing extensions for profile '" + profile.Name + "'")
	cmd_args := []string{".", "--extensions-dir", profile.Path}
	cmd_args = append(cmd_args, getInstallArgs(profile.Extensions)...)

	if len(profile.DependsOn) > 0 {
		for i, dep_exts := range profile.DependsOn {
			fmt.Printf("installing dependency %d/%d for profile '%s'\n", i+1, len(profile.DependsOn), profile.Name)
			cmd_args = append(cmd_args, getInstallArgs(dep_exts)...)
		}
	}

	cmd_args = append(cmd_args, "--force")

	runCmd(cmd_args)
}

// LaunchCode takes a profile and launch code loading only extensions in the profile's path
func LaunchCode(profile config.Profile) {
	println("launching code with profile '" + profile.Name + "'")
	cmd_args := []string{".", "--extensions-dir", profile.Path}

	runCmd(cmd_args)
}

// getInstallArgs returns an array of command arguments to install extensions in the input array.
func getInstallArgs(exts []string) []string {
	var cmd_args = []string{}
	for _, ext := range exts {
		cmd_args = append(cmd_args, "--install-extension", ext)
	}
	return cmd_args
}

// runCmd runs `code` cli with the given command arguments and prints it's output once finished.
func runCmd(cmd_args []string) {
	cmd := exec.Command("code", cmd_args...)

	out, err := cmd.CombinedOutput()
	fmt.Printf("%s", out)
	utils.Check(err)
}
