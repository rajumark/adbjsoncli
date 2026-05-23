# Project Code Guideline

## File Naming Convention

Always create Go files matching the adb command structure using underscore naming.

**Rule:** Take the adb command (everything after `adb`), replace spaces with underscores, and use that as the filename.

### Examples

| adb command | Go file |
|---|---|
| `adb shell dumpsys battery` | `dumpsys_battery.go` |
| `adb shell dumpsys meminfo` | `dumpsys_meminfo.go` |
| `adb shell dumpsys battery set level 50` | `dumpsys_battery_set_level.go` |
| `adb devices -l` | `devices__l.go` |
| `adb shell pm list packages` | `pm_list_packages.go` |
| `adb logcat -b all` | `logcat__b_all.go` |

## JSON Output Structure

Every command must return JSON with the following top-level schema:

```json
{
  "status": 0,
  "output": {}
}
```

- **`status`** (`int`): `0` for success, `1` for error.
- **`output`** (`any`): The parsed/payload data. For known commands (e.g. `dumpsys battery`), this is a structured object/map. For unknown or raw commands, it falls back to a plain string.

### Error example

```json
{
  "status": 1,
  "output": "error message here"
}
```

## No Custom Command Names

Never create new command names. The CLI must always mirror the exact `adb` command structure — only `adb` is replaced with `adbjson`.

**Rule:** If the adb command is `adb shell dumpsys battery`, the adbjson command must be `adbjson shell dumpsys battery`. No custom aliases like `adbjson dumpsys-battery`.

### Correct

```
adbjson shell dumpsys battery
adbjson shell dumpsys meminfo
adbjson devices -l
```

### Incorrect

```
adbjson dumpsys-battery
adbjson show-battery
adbjson battery-status
```

## Build Verification

After creating or modifying any file in the project, always run `build.sh` to ensure there are no build errors.

**Rule:** Run `./scripts/build.sh` after every change. The build must succeed (exit code 0) before considering the work done.
