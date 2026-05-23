package main

func isDumpsysNetstats(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "netstats"
}

func parseDumpsysNetstats(raw string) any {
	return parseColonValueLines(raw)
}
