package main

import "strings"

func isDate(args []string) bool {
	return len(args) == 1 && args[0] == "date"
}

func parseDate(raw string) any {
	return strings.TrimSpace(raw)
}
