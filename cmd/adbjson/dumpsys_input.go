package main

func isDumpsysInput(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "input"
}

func parseDumpsysInput(raw string) any {
	return parseColonValueLines(raw)
}
