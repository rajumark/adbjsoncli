package main

func isDumpsysMount(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "mount"
}

func parseDumpsysMount(raw string) any {
	return parseColonValueLines(raw)
}
