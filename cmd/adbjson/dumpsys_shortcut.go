package main

func isDumpsysShortcut(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "shortcut"
}

func parseDumpsysShortcut(raw string) any {
	return parseColonValueLines(raw)
}
