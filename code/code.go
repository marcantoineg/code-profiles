package code

import (
	"code-profiles/config"
	"code-profiles/utils"
	"fmt"
	"os/exec"
)

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

func LaunchCode(profile config.Profile) {
	println("launching code with profile '" + profile.Name + "'")
	cmd_args := []string{".", "--extensions-dir", profile.Path}

	runCmd(cmd_args)
}

func getInstallArgs(exts []string) []string {
	var cmd_args = []string{}
	for _, ext := range exts {
		cmd_args = append(cmd_args, "--install-extension", ext)
	}
	return cmd_args
}

func runCmd(cmd_args []string) {
	cmd := exec.Command("code", cmd_args...)
	// cmd.Dir = ""

	out, err := cmd.CombinedOutput()
	fmt.Printf("%s", out)
	utils.Check(err)
}
