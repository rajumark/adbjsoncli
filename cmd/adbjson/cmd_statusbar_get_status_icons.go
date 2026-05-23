package main

func isCmdStatusbarGetStatusIcons(args []string) bool {
	return len(args) == 4 && args[0] == "cmd" && args[1] == "statusbar" && args[2] == "get-status-icons"
}

func parseCmdStatusbarGetStatusIcons(raw string) any {
	return parseColonValueLines(raw)
}
