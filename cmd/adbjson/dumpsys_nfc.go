package main

func isDumpsysNfc(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "nfc"
}

func parseDumpsysNfc(raw string) any {
	return parseColonValueLines(raw)
}
