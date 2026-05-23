package main

import (
	"os"
	"os/exec"
	"strings"
)

type cliOutput struct {
	Status int `json:"status"`
	Output any `json:"output"`
}

func runShellCommand() {
	adbPath, lookErr := exec.LookPath("adb")
	if lookErr != nil {
		printJSON(cliOutput{
			Status: 1,
			Output: "adb not found in PATH",
		})
		os.Exit(1)
	}

	shellArgs := os.Args[2:]
	args := append([]string{"shell"}, shellArgs...)

	output, runErr := exec.Command(adbPath, args...).CombinedOutput()
	rawOutput := strings.TrimSpace(string(output))
	if runErr != nil {
		printJSON(cliOutput{
			Status: 1,
			Output: runErr.Error(),
		})
		os.Exit(1)
	}

	out := cliOutput{Status: 0}

	switch {
	case isDumpsysBattery(shellArgs):
		out.Output = parseBatteryProps(rawOutput)
	case isPmListPackages(shellArgs):
		out.Output = parsePmListPackages(rawOutput)
	case isGetprop(shellArgs):
		out.Output = parseGetprop(rawOutput)
	case isDumpsysMeminfo(shellArgs):
		out.Output = parseDumpsysMeminfo(rawOutput)
	case isDumpsysCpuinfo(shellArgs):
		out.Output = parseDumpsysCpuinfo(rawOutput)
	case isDumpsysDiskstats(shellArgs):
		out.Output = parseDumpsysDiskstats(rawOutput)
	case isDumpsysPower(shellArgs):
		out.Output = parseDumpsysPower(rawOutput)
	case isDumpsysWifi(shellArgs):
		out.Output = parseDumpsysWifi(rawOutput)
	case isDumpsysDeviceidle(shellArgs):
		out.Output = parseDumpsysDeviceidle(rawOutput)
	case isDumpsysNotification(shellArgs):
		out.Output = parseDumpsysNotification(rawOutput)
	case isDumpsysActivityActivities(shellArgs):
		out.Output = parseDumpsysActivityActivities(rawOutput)
	case isDumpsysWindowDisplays(shellArgs):
		out.Output = parseDumpsysWindowDisplays(rawOutput)
	case isWmSize(shellArgs):
		out.Output = parseWmSize(rawOutput)
	case isWmDensity(shellArgs):
		out.Output = parseWmDensity(rawOutput)
	case isServiceList(shellArgs):
		out.Output = parseServiceList(rawOutput)
	case isPmListFeatures(shellArgs):
		out.Output = parsePmListFeatures(rawOutput)
	case isPmListPermissions(shellArgs):
		out.Output = parsePmListPermissions(rawOutput)
	case isSettingsList(shellArgs):
		out.Output = parseSettingsList(rawOutput)
	case isDate(shellArgs):
		out.Output = parseDate(rawOutput)
	case isDumpsysBatterystats(shellArgs):
		out.Output = parseDumpsysBatterystats(rawOutput)
	case isDumpsysAlarm(shellArgs):
		out.Output = parseDumpsysAlarm(rawOutput)
	case isDumpsysJobscheduler(shellArgs):
		out.Output = parseDumpsysJobscheduler(rawOutput)
	case isDumpsysNetstats(shellArgs):
		out.Output = parseDumpsysNetstats(rawOutput)
	case isDumpsysUsb(shellArgs):
		out.Output = parseDumpsysUsb(rawOutput)
	case isDumpsysInput(shellArgs):
		out.Output = parseDumpsysInput(rawOutput)
	case isDumpsysGraphicsstats(shellArgs):
		out.Output = parseDumpsysGraphicsstats(rawOutput)
	case isDumpsysAppops(shellArgs):
		out.Output = parseDumpsysAppops(rawOutput)
	case isDumpsysBackup(shellArgs):
		out.Output = parseDumpsysBackup(rawOutput)
	case isDumpsysDropbox(shellArgs):
		out.Output = parseDumpsysDropbox(rawOutput)
	case isPmListLibraries(shellArgs):
		out.Output = parsePmListLibraries(rawOutput)
	case isCmdWifiStatus(shellArgs), isCmdWifiGetCountryCode(shellArgs), isCmdWifiListNetworks(shellArgs), isCmdWifiListScanResults(shellArgs):
		out.Output = parseCmdWifi(rawOutput)
	case isCmdNotificationList(shellArgs):
		out.Output = parseCmdNotificationList(rawOutput)
	case isCmdUimodeNight(shellArgs):
		out.Output = parseCmdUimode(rawOutput)
	case isBmgrEnabled(shellArgs), isBmgrListTransports(shellArgs):
		out.Output = parseBmgr(rawOutput)
	case isCmdNetpolicyGetRestrictBackground(shellArgs):
		out.Output = parseCmdNetpolicy(rawOutput)
	case isSvcPower(shellArgs):
		out.Output = parseSvcPower(rawOutput)
	case isCmdStatusbarGetStatusIcons(shellArgs):
		out.Output = parseCmdStatusbarGetStatusIcons(rawOutput)
	default:
		out.Output = rawOutput
	}

	printJSON(out)
}

func isDumpsysBattery(args []string) bool {
	return len(args) >= 2 && args[0] == "dumpsys" && args[1] == "battery"
}

func parseBatteryProps(raw string) map[string]string {
	props := make(map[string]string)
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if idx := strings.Index(line, ": "); idx != -1 {
			props[strings.TrimSpace(line[:idx])] = strings.TrimSpace(line[idx+2:])
		}
	}
	return props
}
