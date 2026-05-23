package main

func isDumpsysTelecom(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "telecom"
}

func parseDumpsysTelecom(raw string) any {
	return parseColonValueLines(raw)
}
