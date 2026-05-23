package main

func isDumpsysMedia(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "media"
}

func parseDumpsysMedia(raw string) any {
	return parseColonValueLines(raw)
}
