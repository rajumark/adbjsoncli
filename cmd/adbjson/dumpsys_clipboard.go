package main

func isDumpsysClipboard(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "clipboard"
}

func parseDumpsysClipboard(raw string) any {
	return parseColonValueLines(raw)
}
