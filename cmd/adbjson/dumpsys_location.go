package main

func isDumpsysLocation(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "location"
}

func parseDumpsysLocation(raw string) any {
	return parseColonValueLines(raw)
}
