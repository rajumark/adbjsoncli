package main

import (
	"strings"
)

func isDumpsysDiskstats(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "diskstats"
}

func parseDumpsysDiskstats(raw string) any {
	out := make(map[string]string)
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(line, "Data-Free:"):
			out["data_free"] = strings.TrimSpace(strings.TrimPrefix(line, "Data-Free:"))
		case strings.HasPrefix(line, "Cache-Free:"):
			out["cache_free"] = strings.TrimSpace(strings.TrimPrefix(line, "Cache-Free:"))
		case strings.HasPrefix(line, "System-Free:"):
			out["system_free"] = strings.TrimSpace(strings.TrimPrefix(line, "System-Free:"))
		case strings.HasPrefix(line, "Metadata-Free:"):
			out["metadata_free"] = strings.TrimSpace(strings.TrimPrefix(line, "Metadata-Free:"))
		case strings.HasPrefix(line, "Recent Disk Write Speed (kB/s) ="):
			out["disk_write_speed_kbps"] = strings.TrimSpace(strings.TrimPrefix(line, "Recent Disk Write Speed (kB/s) ="))
		case strings.HasPrefix(line, "Latency:"):
			out["latency"] = strings.TrimSpace(strings.TrimPrefix(line, "Latency:"))
		case strings.HasPrefix(line, "File-based Encryption:"):
			out["file_based_encryption"] = strings.TrimSpace(strings.TrimPrefix(line, "File-based Encryption:"))
		}
	}
	return out
}
