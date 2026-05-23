package main

import "strings"

func isWmSize(args []string) bool {
	return len(args) == 2 && args[0] == "wm" && args[1] == "size"
}

func parseWmSize(raw string) any {
	out := make(map[string]string)
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Physical size:") {
			out["physical"] = strings.TrimSpace(line[15:])
		} else if strings.HasPrefix(line, "Override size:") {
			out["override"] = strings.TrimSpace(line[14:])
		}
	}
	return out
}
