package main

func isDumpsysVibrator(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "vibrator"
}

func parseDumpsysVibrator(raw string) any {
	return parseColonValueLines(raw)
}
