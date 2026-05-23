package main

import "strings"

func isCmdNetpolicyGetRestrictBackground(args []string) bool {
	return len(args) == 4 && args[0] == "cmd" && args[1] == "netpolicy" && args[2] == "get" && args[3] == "restrict-background"
}

func parseCmdNetpolicy(raw string) any {
	return strings.TrimSpace(raw)
}
