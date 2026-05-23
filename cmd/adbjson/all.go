package main

type cmdEntry struct {
	Command string `json:"command"`
	Parsed  bool   `json:"parsed"`
}

func runAllCommand() {
	commands := []cmdEntry{
		{Command: "adbjson version", Parsed: true},
		{Command: "adbjson shell <args>", Parsed: false},
	}
	commands = append(commands, parsedShellCommands()...)
	printJSON(cliOutput{Status: 0, Output: commands})
}

func parsedShellCommands() []cmdEntry {
	return []cmdEntry{
		{Command: "adbjson shell date", Parsed: true},
		{Command: "adbjson shell getprop", Parsed: true},
		{Command: "adbjson shell dumpsys battery", Parsed: true},
		{Command: "adbjson shell dumpsys batterystats", Parsed: true},
		{Command: "adbjson shell dumpsys meminfo", Parsed: true},
		{Command: "adbjson shell dumpsys cpuinfo", Parsed: true},
		{Command: "adbjson shell dumpsys diskstats", Parsed: true},
		{Command: "adbjson shell dumpsys power", Parsed: true},
		{Command: "adbjson shell dumpsys wifi", Parsed: true},
		{Command: "adbjson shell dumpsys deviceidle", Parsed: true},
		{Command: "adbjson shell dumpsys notification", Parsed: true},
		{Command: "adbjson shell dumpsys alarm", Parsed: true},
		{Command: "adbjson shell dumpsys jobscheduler", Parsed: true},
		{Command: "adbjson shell dumpsys netstats", Parsed: true},
		{Command: "adbjson shell dumpsys usb", Parsed: true},
		{Command: "adbjson shell dumpsys input", Parsed: true},
		{Command: "adbjson shell dumpsys graphicsstats", Parsed: true},
		{Command: "adbjson shell dumpsys appops", Parsed: true},
		{Command: "adbjson shell dumpsys backup", Parsed: true},
		{Command: "adbjson shell dumpsys dropbox", Parsed: true},
		{Command: "adbjson shell dumpsys activity activities", Parsed: true},
		{Command: "adbjson shell dumpsys window displays", Parsed: true},
		{Command: "adbjson shell wm size", Parsed: true},
		{Command: "adbjson shell wm density", Parsed: true},
		{Command: "adbjson shell service list", Parsed: true},
		{Command: "adbjson shell pm list packages", Parsed: true},
		{Command: "adbjson shell pm list features", Parsed: true},
		{Command: "adbjson shell pm list permissions", Parsed: true},
		{Command: "adbjson shell pm list libraries", Parsed: true},
		{Command: "adbjson shell cmd wifi status", Parsed: true},
		{Command: "adbjson shell cmd wifi get-country-code", Parsed: true},
		{Command: "adbjson shell cmd wifi list-networks", Parsed: true},
		{Command: "adbjson shell cmd notification list", Parsed: true},
		{Command: "adbjson shell cmd uimode night", Parsed: true},
		{Command: "adbjson shell bmgr enabled", Parsed: true},
		{Command: "adbjson shell bmgr list transports", Parsed: true},
		{Command: "adbjson shell cmd netpolicy get restrict-background", Parsed: true},
		{Command: "adbjson shell svc power", Parsed: true},
		{Command: "adbjson shell cmd statusbar get-status-icons", Parsed: true},
		{Command: "adbjson shell settings list system", Parsed: true},
		{Command: "adbjson shell settings list secure", Parsed: true},
		{Command: "adbjson shell settings list global", Parsed: true},
	}
}
