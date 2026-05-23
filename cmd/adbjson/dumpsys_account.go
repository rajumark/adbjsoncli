package main

func isDumpsysAccount(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "account"
}

func parseDumpsysAccount(raw string) any {
	return parseColonValueLines(raw)
}
