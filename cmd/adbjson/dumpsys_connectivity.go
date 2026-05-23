package main

func isDumpsysConnectivity(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "connectivity"
}

func parseDumpsysConnectivity(raw string) any {
	return parseColonValueLines(raw)
}
