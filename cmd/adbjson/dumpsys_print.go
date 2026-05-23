package main

func isDumpsysPrint(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "print"
}

func parseDumpsysPrint(raw string) any {
	return parseColonValueLines(raw)
}
