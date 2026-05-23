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

## CI/CD (GitHub Actions)

The project uses GitHub Actions for continuous integration and release.

### CI (push / pull request)

- **Lint**: `go vet ./...` on Ubuntu
- **Build**: Compiles on `ubuntu-latest`, `macos-latest`, and `windows-latest` to verify the code compiles on all platforms

### Release (tag push `v*`)

When a tag matching `v*` (e.g., `v1.0.0`) is pushed, the release workflow:

1. Cross-compiles binaries for all platforms:
   - `linux/amd64`, `linux/arm64`
   - `darwin/amd64`, `darwin/arm64`
   - `windows/amd64`
2. Generates SHA256 checksums
3. Creates a GitHub Release with the binaries and auto-generated release notes

### Creating a release

```bash
git tag v1.0.0
git push origin v1.0.0
```

### Triggering CI

CI only runs on push when the commit message contains `#go`.

```bash
git commit -m "add new feature #go"
git push
```

**Tag release** (`v*` tags) always runs CI regardless of commit message.

**Rule:** Always verify CI passes before pushing a release tag.
