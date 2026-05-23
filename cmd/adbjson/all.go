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

	printJSON(cliOutput{
		Status: 0,
		Output: commands,
	})
}

func parsedShellCommands() []cmdEntry {
	return []cmdEntry{
		{Command: "adbjson shell getprop", Parsed: true},
		{Command: "adbjson shell dumpsys battery", Parsed: true},
		{Command: "adbjson shell dumpsys meminfo", Parsed: true},
		{Command: "adbjson shell dumpsys cpuinfo", Parsed: true},
		{Command: "adbjson shell dumpsys diskstats", Parsed: true},
		{Command: "adbjson shell dumpsys power", Parsed: true},
		{Command: "adbjson shell dumpsys wifi", Parsed: true},
		{Command: "adbjson shell dumpsys deviceidle", Parsed: true},
		{Command: "adbjson shell dumpsys notification", Parsed: true},
		{Command: "adbjson shell dumpsys activity activities", Parsed: true},
		{Command: "adbjson shell dumpsys window displays", Parsed: true},
		{Command: "adbjson shell wm size", Parsed: true},
		{Command: "adbjson shell wm density", Parsed: true},
		{Command: "adbjson shell service list", Parsed: true},
		{Command: "adbjson shell pm list packages", Parsed: true},
		{Command: "adbjson shell pm list features", Parsed: true},
		{Command: "adbjson shell pm list permissions", Parsed: true},
		{Command: "adbjson shell settings list system", Parsed: true},
		{Command: "adbjson shell settings list secure", Parsed: true},
		{Command: "adbjson shell settings list global", Parsed: true},
	}
}
