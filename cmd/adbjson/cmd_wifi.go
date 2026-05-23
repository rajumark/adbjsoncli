package main

import "strings"

func isCmdWifiStatus(args []string) bool {
	return len(args) == 3 && args[0] == "cmd" && args[1] == "wifi" && args[2] == "status"
}

func isCmdWifiGetCountryCode(args []string) bool {
	return len(args) == 4 && args[0] == "cmd" && args[1] == "wifi" && args[2] == "get-country-code"
}

func isCmdWifiListNetworks(args []string) bool {
	return len(args) == 4 && args[0] == "cmd" && args[1] == "wifi" && args[2] == "list-networks"
}

func isCmdWifiListScanResults(args []string) bool {
	return len(args) == 5 && args[0] == "cmd" && args[1] == "wifi" && args[2] == "list-scan-results"
}

func parseCmdWifi(raw string) any {
	return strings.TrimSpace(raw)
}
