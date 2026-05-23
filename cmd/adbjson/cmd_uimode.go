package main

import "strings"

func isCmdUimodeNight(args []string) bool {
	return len(args) >= 3 && args[0] == "cmd" && args[1] == "uimode" && args[2] == "night"
}

func parseCmdUimode(raw string) any {
	return strings.TrimSpace(raw)
}
