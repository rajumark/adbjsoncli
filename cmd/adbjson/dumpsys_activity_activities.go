package main

import (
	"strconv"
	"strings"
)

func isDumpsysActivityActivities(args []string) bool {
	return len(args) == 3 && args[0] == "dumpsys" && args[1] == "activity" && args[2] == "activities"
}

func parseDumpsysActivityActivities(raw string) any {
	out := make(map[string]any)
	lines := strings.Split(raw, "\n")
	stacks := 0
	tasks := 0
	topActivity := ""
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Stack #") {
			stacks++
		}
		if strings.HasPrefix(line, "Task id #") {
			tasks++
		}
		if strings.HasPrefix(line, "mCurrentFocus=") {
			out["current_focus"] = line[14:]
		} else if strings.HasPrefix(line, "mFocusedApp=") {
			out["focused_app"] = line[12:]
		} else if strings.HasPrefix(line, "mResumedActivity:") && topActivity == "" {
			topActivity = strings.TrimSpace(line[18:])
		}
		_ = i
	}
	out["stack_count"], _ = strconv.Atoi(strconv.Itoa(stacks))
	out["task_count"], _ = strconv.Atoi(strconv.Itoa(tasks))
	if topActivity != "" {
		out["top_activity"] = topActivity
	}
	return out
}
