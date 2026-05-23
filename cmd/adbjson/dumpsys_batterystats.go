package main

import "strings"

func isDumpsysBatterystats(args []string) bool {
	return len(args) >= 2 && args[0] == "dumpsys" && args[1] == "batterystats"
}

func parseDumpsysBatterystats(raw string) any {
	out := parseColonValueLines(raw)
	if len(out) == 0 {
		return strings.TrimSpace(raw)
	}
	return out
}
