package main

import "strings"

func isDumpsysWindowDisplays(args []string) bool {
	return len(args) == 3 && args[0] == "dumpsys" && args[1] == "window" && args[2] == "displays"
}

func parseDumpsysWindowDisplays(raw string) any {
	out := make(map[string]any)
	displays := make([]map[string]string, 0)
	var cur map[string]string
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Display ") {
			if cur != nil {
				displays = append(displays, cur)
			}
			cur = make(map[string]string)
			cur["id"] = strings.TrimPrefix(line, "Display ")
			continue
		}
		if cur != nil {
			if strings.HasPrefix(line, "mDisplayId=") {
				cur["display_id"] = line[11:]
			} else if strings.HasPrefix(line, "mBounds=") {
				cur["bounds"] = line[8:]
			} else if strings.HasPrefix(line, "mCurrentFocus=") {
				cur["current_focus"] = line[14:]
			} else if strings.HasPrefix(line, "mFocusedApp=") {
				cur["focused_app"] = line[12:]
			} else if idx := strings.Index(line, "="); idx != -1 {
				cur[line[:idx]] = line[idx+1:]
			}
		}
	}
	if cur != nil {
		displays = append(displays, cur)
	}
	out["display_count"] = len(displays)
	out["displays"] = displays
	return out
}
