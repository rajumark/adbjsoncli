package main

import (
	"strings"
)

func isDumpsysMeminfo(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "meminfo"
}

func parseDumpsysMeminfo(raw string) any {
	out := make(map[string]string)
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		for _, prefix := range []string{"Total RAM:", "Free RAM:", "Used RAM:", "Lost RAM:", "ZRAM:"} {
			if strings.HasPrefix(line, prefix) {
				out[snakeKey(prefix)] = strings.TrimSpace(line[len(prefix):])
			}
		}
	}
	return out
}

func snakeKey(s string) string {
	return strings.ReplaceAll(strings.ToLower(strings.TrimSuffix(s, ":")), " ", "_")
}
