package main

func isDumpsysProcessinfo(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "processinfo"
}

func parseDumpsysProcessinfo(raw string) any {
	return parseColonValueLines(raw)
}
