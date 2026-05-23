package main

func isDumpsysAlarm(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "alarm"
}

func parseDumpsysAlarm(raw string) any {
	return parseColonValueLines(raw)
}
