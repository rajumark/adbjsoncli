package main

func isDumpsysAppwidget(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "appwidget"
}

func parseDumpsysAppwidget(raw string) any {
	return parseColonValueLines(raw)
}
