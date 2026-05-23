package main

import "strings"

func isPmListPermissions(args []string) bool {
	return len(args) == 3 && args[0] == "pm" && args[1] == "list" && args[2] == "permissions"
}

func parsePmListPermissions(raw string) any {
	permissions := make([]map[string]string, 0)
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "permission:") {
			continue
		}
		rest := strings.TrimPrefix(line, "permission:")
		p := make(map[string]string)
		parts := strings.Split(rest, ":")
		if len(parts) > 0 {
			p["name"] = parts[0]
		}
		for i := 1; i < len(parts); i++ {
			if kv := strings.SplitN(parts[i], "=", 2); len(kv) == 2 {
				p[kv[0]] = kv[1]
			}
		}
		permissions = append(permissions, p)
	}
	return permissions
}
