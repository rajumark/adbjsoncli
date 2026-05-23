package main

func isDumpsysTrust(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "trust"
}

func parseDumpsysTrust(raw string) any {
	return parseColonValueLines(raw)
}
