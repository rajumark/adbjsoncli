package main

import "strings"

func isDumpsysDeviceidle(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "deviceidle"
}

func parseDumpsysDeviceidle(raw string) any {
	out := make(map[string]string)
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if idx := strings.Index(line, "="); idx != -1 {
			key := strings.TrimSpace(line[:idx])
			val := strings.TrimSpace(line[idx+1:])
			if strings.HasPrefix(key, "m") || strings.HasPrefix(key, "Settings") || strings.HasPrefix(key, "State:") {
				out[key] = val
			}
		}
	}
	if v, ok := out["State:"]; ok {
		out["state"] = v
		delete(out, "State:")
	}
	return out
}
