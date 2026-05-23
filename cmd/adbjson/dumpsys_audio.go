package main

func isDumpsysAudio(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "audio"
}

func parseDumpsysAudio(raw string) any {
	return parseColonValueLines(raw)
}
