package main

func isDumpsysSafetyCenter(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "safety_center"
}

func parseDumpsysSafetyCenter(raw string) any {
	return parseColonValueLines(raw)
}
