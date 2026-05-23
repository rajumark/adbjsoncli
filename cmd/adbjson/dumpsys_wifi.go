package main

import "strings"

func isDumpsysWifi(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "wifi"
}

func parseDumpsysWifi(raw string) any {
	out := make(map[string]string)
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "SSID:") {
			out["ssid"] = strings.TrimSpace(line[5:])
		} else if strings.HasPrefix(line, "BSSID:") {
			out["bssid"] = strings.TrimSpace(line[6:])
		} else if strings.HasPrefix(line, "ipAddress") || strings.HasPrefix(line, "ip address") {
			parts := strings.SplitN(line, " ", 2)
			if len(parts) == 2 {
				out["ip_address"] = strings.TrimSpace(parts[1])
			}
		} else if strings.HasPrefix(line, "mWifiState=") {
			out["wifi_state"] = line[11:]
		} else if strings.HasPrefix(line, "mNetworkInfo") && strings.Contains(line, "state:") {
			if idx := strings.Index(line, "state:"); idx != -1 {
				out["network_state"] = strings.TrimSpace(line[idx+6:])
			}
		}
	}
	return out
}
