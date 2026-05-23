package main

func isDumpsysDevicePolicy(args []string) bool {
	return len(args) == 3 && args[0] == "dumpsys" && args[1] == "device_policy"
}

func parseDumpsysDevicePolicy(raw string) any {
	return parseColonValueLines(raw)
}
