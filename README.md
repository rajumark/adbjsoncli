# adbjson

[![CI](https://github.com/raju/adbjsoncli/actions/workflows/ci.yml/badge.svg)](https://github.com/raju/adbjsoncli/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/github/go-mod/go-version/raju/adbjsoncli)](https://go.dev/)
[![Release](https://img.shields.io/github/v/release/raju/adbjsoncli)](https://github.com/raju/adbjsoncli/releases)
[![License](https://img.shields.io/github/license/raju/adbjsoncli)](LICENSE)

`adbjson` is a Go CLI that wraps `adb` commands and prints structured JSON output.

Use the exact same `adb` command structure — just replace `adb` with `adbjson`.

```bash
adb version          →  adbjson version
adb shell dumpsys battery  →  adbjson shell dumpsys battery
adb shell pm list packages →  adbjson shell pm list packages
```

## Installation

### Download (pre-built binary)

Download the latest binary for your platform from the [releases page](https://github.com/raju/adbjsoncli/releases).

| Platform | Download |
|---|---|
| macOS (Intel) | `adbjson-darwin-amd64` |
| macOS (Apple Silicon) | `adbjson-darwin-arm64` |
| Linux (x86_64) | `adbjson-linux-amd64` |
| Linux (ARM64) | `adbjson-linux-arm64` |
| Windows (x86_64) | `adbjson-windows-amd64.exe` |

```bash
# macOS / Linux
chmod +x adbjson-darwin-arm64
sudo mv adbjson-darwin-arm64 /usr/local/bin/adbjson

# Windows (PowerShell)
# Rename to adbjson.exe and add to PATH
```

### Build from source

```bash
git clone https://github.com/raju/adbjsoncli.git
cd adbjsoncli
./scripts/build.sh
```

The binary is created at `bin/adbjson`.

### Go install

```bash
go install github.com/raju/adbjsoncli/cmd/adbjson@latest
```

## Usage

```bash
adbjson version
adbjson shell dumpsys battery
adbjson shell pm list packages
```

### Output format

```json
{
  "status": 0,
  "output": {}
}
```

- `status`: `0` for success, `1` for error
- `output`: parsed JSON data or raw fallback

### Examples

```bash
adbjson version
```

```json
{
  "status": 0,
  "output": {
    "version": "37.0.0-14910828"
  }
}
```

```bash
adbjson shell dumpsys battery
```

```json
{
  "status": 0,
  "output": {
    "level": "100",
    "status": "4",
    "temperature": "250",
    "technology": "Li-ion"
  }
}
```

```bash
adbjson shell pm list packages -3
```

```json
{
  "status": 0,
  "output": [
    {
      "package_name": "com.example.app"
    }
  ]
}
```

## Documentation

- **📁 Markdown files** — [docs/](./docs) — per-command and reference docs in markdown
- **🌐 Website** — [rajumark.github.io/adbjsoncli](https://rajumark.github.io/adbjsoncli) — interactive browser with sidebar

## Development

### Prerequisites

- Go 1.26+
- [adb](https://developer.android.com/studio/command-line/adb) (Android SDK platform-tools)

### Build

```bash
./scripts/build.sh
```

Cross-compile for a specific platform:

```bash
GOOS=linux GOARCH=arm64 ./scripts/build.sh
```

### Add a new command

1. Create a new Go file in `cmd/adbjson/` named after the adb subcommand (e.g., `dumpsys_meminfo.go`)
2. Implement the parsing logic
3. Register it in the shell dispatch inside `dumpsys_battery.go`
4. Run `./scripts/build.sh` to verify

### Regenerate documentation

```bash
./scripts/gen_docs.sh
```

## Project guidelines

See [project_code_guidline.md](./project_code_guidline.md) for project conventions.
