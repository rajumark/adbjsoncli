package main

func isDumpsysConnectivityNative(args []string) bool {
	return len(args) == 3 && args[0] == "dumpsys" && args[1] == "connectivity_native"
}

func parseDumpsysConnectivityNative(raw string) any {
	return parseColonValueLines(raw)
}
