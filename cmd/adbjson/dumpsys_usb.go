package main

func isDumpsysUsb(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "usb"
}

func parseDumpsysUsb(raw string) any {
	return parseColonValueLines(raw)
}
