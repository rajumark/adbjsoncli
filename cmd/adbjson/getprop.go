package main

import "strings"

func isGetprop(args []string) bool {
	return len(args) == 1 && args[0] == "getprop"
}

func parseGetprop(raw string) any {
	m := make(map[string]string)
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if len(line) > 2 && line[0] == '[' {
			end := strings.IndexByte(line, ']')
			if end != -1 && end+3 < len(line) && line[end+1] == ':' && line[end+2] == ' ' && line[end+3] == '[' {
				key := line[1:end]
				valEnd := strings.LastIndexByte(line, ']')
				if valEnd > end+3 {
					m[key] = line[end+4 : valEnd]
				}
			}
		}
	}
	return m
}
