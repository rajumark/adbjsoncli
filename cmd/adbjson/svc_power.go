package main

import "strings"

func isSvcPower(args []string) bool {
	return len(args) >= 2 && args[0] == "svc" && args[1] == "power"
}

func parseSvcPower(raw string) any {
	return strings.TrimSpace(raw)
}
