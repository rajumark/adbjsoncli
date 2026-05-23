package main

func isDumpsysAppops(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "appops"
}

func parseDumpsysAppops(raw string) any {
	return parseColonValueLines(raw)
}
