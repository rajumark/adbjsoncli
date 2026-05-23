package main

import "strings"

func isCmdConnectivityAirplaneMode(args []string) bool {
	return len(args) == 4 && args[0] == "cmd" && args[1] == "connectivity" && args[2] == "airplane-mode"
}

func parseCmdConnectivityAirplaneMode(raw string) string {
	return strings.TrimSpace(raw)
}
