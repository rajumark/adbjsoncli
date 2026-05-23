package main

func isDumpsysWebviewupdate(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "webviewupdate"
}

func parseDumpsysWebviewupdate(raw string) any {
	return parseColonValueLines(raw)
}
