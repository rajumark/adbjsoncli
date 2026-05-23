package main

func isDumpsysAccessibility(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "accessibility"
}

func parseDumpsysAccessibility(raw string) any {
	return parseColonValueLines(raw)
}
