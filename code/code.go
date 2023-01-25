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
	for _, ext := range profile.Extensions {
		cmd_args = append(cmd_args, "--install-extension", ext)
	}
	cmd_args = append(cmd_args, "--force")

	runCmd(cmd_args)
}

func LaunchCode(profile config.Profile) {
	println("launching code with profile '" + profile.Name + "'")
	cmd_args := []string{".", "--extensions-dir", profile.Path}

	runCmd(cmd_args)
}

func runCmd(cmd_args []string) {
	cmd := exec.Command("code", cmd_args...)
	// cmd.Dir = ""

	out, err := cmd.CombinedOutput()
	fmt.Printf("%s", out)
	utils.Check(err)
}
