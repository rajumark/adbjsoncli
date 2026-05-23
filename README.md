# adbjsoncli

Simple Go CLI that wraps `adb` commands and prints JSON.

## Build

```zsh
cd /Users/raju/Documents/allproj/adbjsoncli
./scripts/build.sh
```

The binary is created here:

```text
/Users/raju/Documents/allproj/adbjsoncli/bin/adbjson
```

## Run

After building, run:

```zsh
adbjson version
```

Example output:

```json
{
  "command": "version",
  "adb_path": "/Users/raju/Desktop/tools/platform-tools/adb",
  "version": "37.0.0-14910828",
  "raw_output": "Android Debug Bridge version 1.0.41...",
  "successful": true
}
```

The build folder is added to `~/.zshrc`, so new Terminal windows can run `adbjson`
from any path. In the current Terminal, run this once if needed:

```zsh
source ~/.zshrc
```

## Add More Commands

Add a new Go file inside:

```text
/Users/raju/Documents/allproj/adbjsoncli/cmd/adbjson
```

For example, create `devices.go` with a `runDevicesCommand()` function, then add a new case in `main.go`:

```go
case "devices":
	runDevicesCommand()
```

Keep each command in its own file so the project stays easy to grow.
