package main

func isDumpsysMediaSession(args []string) bool {
	return len(args) == 3 && args[0] == "dumpsys" && args[1] == "media_session"
}

func parseDumpsysMediaSession(raw string) any {
	return parseColonValueLines(raw)
}
