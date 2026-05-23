package main

func isDumpsysNetpolicy(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "netpolicy"
}

func parseDumpsysNetpolicy(raw string) any {
	return parseColonValueLines(raw)
}
