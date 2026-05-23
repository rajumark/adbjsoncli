package main

func isDumpsysNetworkTimeUpdateService(args []string) bool {
	return len(args) == 3 && args[0] == "dumpsys" && args[1] == "network_time_update_service"
}

func parseDumpsysNetworkTimeUpdateService(raw string) any {
	return parseColonValueLines(raw)
}
