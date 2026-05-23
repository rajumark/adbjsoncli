package main

func isDumpsysBackup(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "backup"
}

func parseDumpsysBackup(raw string) any {
	return parseColonValueLines(raw)
}
