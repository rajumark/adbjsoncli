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
		{Command: "adbjson shell dumpsys battery", Parsed: true},
		{Command: "adbjson shell pm list packages", Parsed: true},
	}
}
