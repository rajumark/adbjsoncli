package main

func isDumpsysNetworkScore(args []string) bool {
	return len(args) == 3 && args[0] == "dumpsys" && args[1] == "network_score"
}

func parseDumpsysNetworkScore(raw string) any {
	return parseColonValueLines(raw)
}
