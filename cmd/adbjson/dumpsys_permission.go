package main

func isDumpsysPermission(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "permission"
}

func parseDumpsysPermission(raw string) any {
	return parseColonValueLines(raw)
}
