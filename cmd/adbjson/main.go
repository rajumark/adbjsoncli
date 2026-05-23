package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printJSON(cliOutput{
			Status: 1,
			Output: "missing command",
		})
		os.Exit(1)
	}

	switch os.Args[1] {
	case "version":
		runVersionCommand()
	case "shell":
		runShellCommand()
	case "all":
		runAllCommand()
	default:
		printJSON(cliOutput{
			Status: 1,
			Output: fmt.Sprintf("unknown command: %s", os.Args[1]),
		})
		os.Exit(1)
	}
}

func printJSON(value any) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)
	_ = encoder.Encode(value)
}
