package main

import "strings"

func isLs(args []string) bool {
	return len(args) >= 1 && args[0] == "ls"
}

func parseLs(raw string) string {
	return strings.TrimSpace(raw)
}

func isCatProcVersion(args []string) bool {
	return len(args) == 3 && args[0] == "cat" && args[1] == "/proc/version"
}

func parseCatProcVersion(raw string) string {
	return strings.TrimSpace(raw)
}

func isCatProcCpuinfo(args []string) bool {
	return len(args) == 3 && args[0] == "cat" && args[1] == "/proc/cpuinfo"
}

func parseCatProcCpuinfo(raw string) any {
	return parseColonValueLines(raw)
}

func isGetpropSingle(args []string) bool {
	return len(args) == 2 && args[0] == "getprop"
}

func parseGetpropSingle(raw string) string {
	return strings.TrimSpace(raw)
}

func isSettingsGet(args []string) bool {
	return len(args) == 4 && args[0] == "settings" && args[1] == "get"
}

func parseSettingsGet(raw string) string {
	return strings.TrimSpace(raw)
}

func isCmdBatteryReset(args []string) bool {
	return len(args) == 3 && args[0] == "cmd" && args[1] == "battery" && args[2] == "reset"
}

func parseCmdBatteryReset(raw string) string {
	return strings.TrimSpace(raw)
}

func isCmdDeviceidle(args []string) bool {
	if len(args) < 3 || args[0] != "cmd" || args[1] != "deviceidle" {
		return false
	}
	switch args[2] {
	case "step", "force-idle", "unforce":
		return true
	}
	return false
}

func parseCmdDeviceidle(raw string) string {
	return strings.TrimSpace(raw)
}
