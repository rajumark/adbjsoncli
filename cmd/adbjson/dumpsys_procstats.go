package main

func isDumpsysProcstats(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "procstats"
}

func parseDumpsysProcstats(raw string) any {
	return parseColonValueLines(raw)
}
