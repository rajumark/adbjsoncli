package main

func isDumpsysVoiceinteraction(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "voiceinteraction"
}

func parseDumpsysVoiceinteraction(raw string) any {
	return parseColonValueLines(raw)
}
