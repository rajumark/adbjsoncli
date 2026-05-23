package main

import "strings"

func isSettingsList(args []string) bool {
	return len(args) >= 2 && args[0] == "settings" && args[1] == "list"
}

func parseSettingsList(raw string) any {
	out := make(map[string]string)
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if idx := strings.Index(line, "="); idx != -1 {
			out[line[:idx]] = line[idx+1:]
		}
	}
	return out
}
