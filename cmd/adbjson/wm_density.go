package main

import "strings"

func isWmDensity(args []string) bool {
	return len(args) == 2 && args[0] == "wm" && args[1] == "density"
}

func parseWmDensity(raw string) any {
	out := make(map[string]string)
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Physical density:") {
			out["physical"] = strings.TrimSpace(line[17:])
		} else if strings.HasPrefix(line, "Override density:") {
			out["override"] = strings.TrimSpace(line[16:])
		}
	}
	return out
}
