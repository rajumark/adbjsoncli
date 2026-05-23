package main

func isDumpsysNetworkManagement(args []string) bool {
	return len(args) == 3 && args[0] == "dumpsys" && args[1] == "network_management"
}

func parseDumpsysNetworkManagement(raw string) any {
	return parseColonValueLines(raw)
}
