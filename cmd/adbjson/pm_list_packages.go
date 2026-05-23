package main

import "strings"

type packageInfo struct {
	PackageName string `json:"package_name"`
	APKPath     string `json:"apk_path,omitempty"`
	Installer   string `json:"installer,omitempty"`
	VersionCode string `json:"version_code,omitempty"`
}

func isPmListPackages(args []string) bool {
	return len(args) >= 3 && args[0] == "pm" && args[1] == "list" && args[2] == "packages"
}

func parsePmListPackages(raw string) any {
	packages := make([]packageInfo, 0)
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "package:") {
			continue
		}
		rest := strings.TrimPrefix(line, "package:")

		info := packageInfo{}

		if idx := strings.LastIndex(rest, "  installer="); idx != -1 {
			info.Installer = strings.TrimPrefix(rest[idx+2:], "installer=")
			rest = rest[:idx]
		}

		if idx := strings.LastIndex(rest, " versionCode:"); idx != -1 {
			info.VersionCode = rest[idx+len(" versionCode:"):]
			rest = rest[:idx]
		}

		if idx := strings.LastIndex(rest, "="); idx != -1 {
			info.APKPath = rest[:idx]
			rest = rest[idx+1:]
		}

		info.PackageName = rest
		packages = append(packages, info)
	}
	return packages
}
