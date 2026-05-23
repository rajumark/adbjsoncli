package main

import "strings"

func isDumpsysDeviceConfig(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "device_config"
}

func parseDumpsysDeviceConfig(raw string) map[string]string {
	out := make(map[string]string)
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if idx := strings.Index(line, "="); idx != -1 {
			out[strings.TrimSpace(line[:idx])] = strings.TrimSpace(line[idx+1:])
		}
	}
	return out
}
