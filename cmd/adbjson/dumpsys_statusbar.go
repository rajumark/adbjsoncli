package main

func isDumpsysStatusbar(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "statusbar"
}

func parseDumpsysStatusbar(raw string) any {
	return parseColonValueLines(raw)
}
