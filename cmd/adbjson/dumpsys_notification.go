package main

import (
	"strconv"
	"strings"
)

func isDumpsysNotification(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "notification"
}

func parseDumpsysNotification(raw string) any {
	out := make(map[string]any)
	listeners := 0
	notifications := 0
	inListeners := false
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Notification listeners:") {
			inListeners = true
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				listeners, _ = strconv.Atoi(strings.TrimSpace(parts[1]))
			}
			continue
		}
		if inListeners {
			if line == "" {
				inListeners = false
				continue
			}
			listeners++
		}
		if strings.HasPrefix(line, "mNotificationsByKey:") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				notifications, _ = strconv.Atoi(strings.TrimSpace(parts[1]))
			}
		}
	}
	out["listener_count"] = listeners
	out["notification_count"] = notifications
	return out
}
