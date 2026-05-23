package main

import "strings"

func isDumpsysPower(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "power"
}

func parseDumpsysPower(raw string) any {
	out := make(map[string]string)
	started := false
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "(dumpsys power)") || strings.HasPrefix(line, "Power Manager State:") {
			started = true
			continue
		}
		if started {
			if line == "" && len(out) > 0 {
				break
			}
			if idx := strings.Index(line, "="); idx != -1 {
				out[line[:idx]] = line[idx+1:]
			} else if idx := strings.Index(line, ": "); idx != -1 {
				out[line[:idx]] = line[idx+2:]
			}
		}
	}
	return out
}
