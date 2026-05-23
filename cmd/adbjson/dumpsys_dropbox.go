package main

func isDumpsysDropbox(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "dropbox"
}

func parseDumpsysDropbox(raw string) any {
	return parseColonValueLines(raw)
}
