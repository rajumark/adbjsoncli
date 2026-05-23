package main

import "strings"

func isDumpsysCpuinfo(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "cpuinfo"
}

func parseDumpsysCpuinfo(raw string) any {
	out := make(map[string]any)
	var processes []string
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "LOAD:") {
			out["load"] = strings.TrimSpace(line[5:])
		} else if line != "" && !strings.HasPrefix(line, "LOAD:") && !strings.HasPrefix(line, "Load:") {
			processes = append(processes, line)
		}
	}
	if len(processes) > 0 {
		out["processes"] = processes
	}
	return out
}
