package main

import (
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func runVersionCommand() {
	adbPath, lookErr := exec.LookPath("adb")
	if lookErr != nil {
		printJSON(cliOutput{
			Status: 1,
			Output: "adb not found in PATH",
		})
		os.Exit(1)
	}

	output, runErr := exec.Command(adbPath, "version").CombinedOutput()
	rawOutput := strings.TrimSpace(string(output))
	if runErr != nil {
		printJSON(cliOutput{
			Status: 1,
			Output: runErr.Error(),
		})
		os.Exit(1)
	}

	printJSON(cliOutput{
		Status: 0,
		Output: map[string]string{
			"version": parseADBVersion(rawOutput),
		},
	})
}

func parseADBVersion(rawOutput string) string {
	matches := regexp.MustCompile(`Version\s+([^\s]+)`).FindStringSubmatch(rawOutput)
	if len(matches) == 2 {
		return matches[1]
	}
	return ""
}
