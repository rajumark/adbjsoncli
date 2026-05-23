package main

func isPmListLibraries(args []string) bool {
	return len(args) == 3 && args[0] == "pm" && args[1] == "list" && args[2] == "libraries"
}

func parsePmListLibraries(raw string) any {
	return parsePrefixedLines(raw, "library:")
}
