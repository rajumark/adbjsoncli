package main

func isDumpsysSensorservice(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "sensorservice"
}

func parseDumpsysSensorservice(raw string) any {
	return parseColonValueLines(raw)
}
