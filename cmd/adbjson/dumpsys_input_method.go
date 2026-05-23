package main

func isDumpsysInputMethod(args []string) bool {
	return len(args) == 3 && args[0] == "dumpsys" && args[1] == "input_method"
}

func parseDumpsysInputMethod(raw string) any {
	return parseColonValueLines(raw)
}
