package main

import "strings"

func isShellArgs(args []string, cmd ...string) bool {
	if len(args) < len(cmd) {
		return false
	}
	for i, c := range cmd {
		if args[i] != c {
			return false
		}
	}
	return true
}

func parseColonValueLines(raw string) map[string]string {
	out := make(map[string]string)
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if idx := strings.Index(line, ": "); idx != -1 {
			out[strings.TrimSpace(line[:idx])] = strings.TrimSpace(line[idx+2:])
		}
	}
	return out
}

func parseKeyValueLines(raw, sep string) map[string]string {
	out := make(map[string]string)
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if idx := strings.Index(line, sep); idx != -1 {
			out[strings.TrimSpace(line[:idx])] = strings.TrimSpace(line[idx+len(sep):])
		}
	}
	return out
}

func parsePrefixedLines(raw, prefix string) []string {
	out := make([]string, 0)
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, prefix) {
			out = append(out, strings.TrimPrefix(line, prefix))
		}
	}
	return out
}

func firstMatch(lines []string, prefix string) string {
	for _, l := range lines {
		if strings.HasPrefix(l, prefix) {
			return strings.TrimPrefix(l, prefix)
		}
	}
	return ""
}
