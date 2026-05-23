package main

func isDumpsysUimode(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "uimode"
}

func parseDumpsysUimode(raw string) any {
	return parseColonValueLines(raw)
}
