package main

func isDumpsysRole(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "role"
}

func parseDumpsysRole(raw string) any {
	return parseColonValueLines(raw)
}
