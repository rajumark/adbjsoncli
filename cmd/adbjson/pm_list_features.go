package main

import "strings"

func isPmListFeatures(args []string) bool {
	return len(args) == 3 && args[0] == "pm" && args[1] == "list" && args[2] == "features"
}

func parsePmListFeatures(raw string) any {
	features := make([]string, 0)
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "feature:") {
			features = append(features, strings.TrimPrefix(line, "feature:"))
		}
	}
	return features
}
