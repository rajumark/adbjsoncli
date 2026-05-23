package main

import "strings"

func isBmgrEnabled(args []string) bool {
	return len(args) == 2 && args[0] == "bmgr" && args[1] == "enabled"
}

func isBmgrListTransports(args []string) bool {
	return len(args) == 3 && args[0] == "bmgr" && args[1] == "list" && args[2] == "transports"
}

func parseBmgr(raw string) any {
	return strings.TrimSpace(raw)
}
