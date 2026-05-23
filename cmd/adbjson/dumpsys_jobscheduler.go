package main

func isDumpsysJobscheduler(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "jobscheduler"
}

func parseDumpsysJobscheduler(raw string) any {
	return parseColonValueLines(raw)
}
