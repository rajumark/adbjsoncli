package main

func isCmdNotificationList(args []string) bool {
	return len(args) == 3 && args[0] == "cmd" && args[1] == "notification" && args[2] == "list"
}

func parseCmdNotificationList(raw string) any {
	return parseColonValueLines(raw)
}
