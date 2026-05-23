package main

import (
	"strconv"
	"strings"
)

func isServiceList(args []string) bool {
	return len(args) == 2 && args[0] == "service" && args[1] == "list"
}

func parseServiceList(raw string) any {
	out := make(map[string]any)
	lines := strings.Split(raw, "\n")
	services := make([]map[string]string, 0)
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if i == 0 && strings.HasPrefix(line, "Found") {
			parts := strings.Split(line, " ")
			if len(parts) > 1 {
				out["service_count"], _ = strconv.Atoi(parts[1])
			}
			continue
		}
		if line == "" || !strings.Contains(line, ":") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			num := fields[0]
			if svcIdx := strings.Index(line, ": ["); svcIdx != -1 {
				name := strings.TrimSpace(line[len(num):svcIdx])
				iface := strings.TrimSpace(line[svcIdx+2:])
				if iface == "[" {
					iface = "[]"
				} else if strings.HasPrefix(iface, "[") {
					iface = iface[1:]
					iface = strings.TrimSuffix(iface, "]")
				}
				services = append(services, map[string]string{
					"id":        num,
					"name":      strings.TrimSpace(name),
					"interface": iface,
				})
			}
		}
	}
	out["services"] = services
	return out
}
