package main

func isDumpsysTranslation(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "translation"
}

func parseDumpsysTranslation(raw string) any {
	return parseColonValueLines(raw)
}
