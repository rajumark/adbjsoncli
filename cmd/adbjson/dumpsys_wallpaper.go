package main

func isDumpsysWallpaper(args []string) bool {
	return len(args) == 2 && args[0] == "dumpsys" && args[1] == "wallpaper"
}

func parseDumpsysWallpaper(raw string) any {
	return parseColonValueLines(raw)
}
