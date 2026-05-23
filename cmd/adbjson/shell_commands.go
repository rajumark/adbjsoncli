package main

import "strings"

func isUptime(args []string) bool {
	return len(args) == 1 && args[0] == "uptime"
}

func parseUptime(raw string) any {
	return parseColonValueLines(raw)
}

func isUname(args []string) bool {
	return len(args) == 2 && args[0] == "uname" && args[1] == "-a"
}

func parseUname(raw string) string {
	return strings.TrimSpace(raw)
}

func isId(args []string) bool {
	return len(args) == 1 && args[0] == "id"
}

func parseId(raw string) any {
	return parseColonValueLines(raw)
}

func isDf(args []string) bool {
	return len(args) == 1 && args[0] == "df"
}

func parseDf(raw string) any {
	return parseColonValueLines(raw)
}

func isFree(args []string) bool {
	return len(args) == 1 && args[0] == "free"
}

func parseFree(raw string) any {
	return parseColonValueLines(raw)
}

func isPs(args []string) bool {
	return len(args) == 1 && args[0] == "ps"
}

func parsePs(raw string) string {
	return strings.TrimSpace(raw)
}

func isPrintenv(args []string) bool {
	return len(args) == 1 && args[0] == "printenv"
}

func parsePrintenv(raw string) any {
	return parseKeyValueLines(raw, "=")
}

func isIfconfig(args []string) bool {
	return len(args) == 1 && args[0] == "ifconfig"
}

func parseIfconfig(raw string) string {
	return strings.TrimSpace(raw)
}

func isNetstat(args []string) bool {
	return len(args) == 1 && args[0] == "netstat"
}

func parseNetstat(raw string) string {
	return strings.TrimSpace(raw)
}

func isVmstat(args []string) bool {
	return len(args) == 1 && args[0] == "vmstat"
}

func parseVmstat(raw string) any {
	return parseColonValueLines(raw)
}

func isGroups(args []string) bool {
	return len(args) == 1 && args[0] == "groups"
}

func parseGroups(raw string) string {
	return strings.TrimSpace(raw)
}

func isWhoami(args []string) bool {
	return len(args) == 1 && args[0] == "whoami"
}

func parseWhoami(raw string) string {
	return strings.TrimSpace(raw)
}

func isEnv(args []string) bool {
	return len(args) == 1 && args[0] == "env"
}

func parseEnv(raw string) any {
	return parseKeyValueLines(raw, "=")
}
