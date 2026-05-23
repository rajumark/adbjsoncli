package main

func isDumpsysTelephony(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "telephony"
}

func parseDumpsysTelephony(raw string) any {
	return parseColonValueLines(raw)
}
