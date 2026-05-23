package main

import (
	"os"
	"os/exec"
	"strings"
)

type cliOutput struct {
	Status int `json:"status"`
	Output any `json:"output"`
}

func runShellCommand() {
	adbPath, lookErr := exec.LookPath("adb")
	if lookErr != nil {
		printJSON(cliOutput{
			Status: 1,
			Output: "adb not found in PATH",
		})
		os.Exit(1)
	}

	shellArgs := os.Args[2:]
	args := append([]string{"shell"}, shellArgs...)

	output, runErr := exec.Command(adbPath, args...).CombinedOutput()
	rawOutput := strings.TrimSpace(string(output))
	if runErr != nil {
		printJSON(cliOutput{
			Status: 1,
			Output: runErr.Error(),
		})
		os.Exit(1)
	}

	out := cliOutput{Status: 0}

	switch {
	case isDumpsysBattery(shellArgs):
		out.Output = parseBatteryProps(rawOutput)
	case isPmListPackages(shellArgs):
		out.Output = parsePmListPackages(rawOutput)
	default:
		out.Output = rawOutput
	}

	printJSON(out)
}

func isDumpsysBattery(args []string) bool {
	return len(args) >= 2 && args[0] == "dumpsys" && args[1] == "battery"
}

func parseBatteryProps(raw string) map[string]string {
	props := make(map[string]string)
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if idx := strings.Index(line, ": "); idx != -1 {
			props[strings.TrimSpace(line[:idx])] = strings.TrimSpace(line[idx+2:])
		}
	}
	return props
}
