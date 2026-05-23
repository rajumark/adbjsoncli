package main

import "strings"

func isCmdBluetoothManager(args []string) bool {
	return len(args) >= 3 && args[0] == "cmd" && args[1] == "bluetooth_manager"
}

func parseCmdBluetoothManager(raw string) any {
	return strings.TrimSpace(raw)
}
