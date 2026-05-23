#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
DOC_DIR="$ROOT_DIR/documentation"
BIN="$ROOT_DIR/bin/adbjson"

mkdir -p "$DOC_DIR"

ADB=$(command -v adb 2>/dev/null || echo "")

generate_doc() {
  local adbjson_cmd="$1"
  local adb_cmd="$2"
  local filename="${adb_cmd// /_}.md"

  echo "Generating $filename ..."

  # shellcheck disable=SC2086
  local json_output
  json_output=$("$BIN" $adb_cmd 2>&1 || true)

  local raw_output=""
  if [ -n "$ADB" ]; then
    # shellcheck disable=SC2086
    raw_output=$("$ADB" $adb_cmd 2>&1 || true)
  else
    raw_output="adb not found in PATH"
  fi

  # Truncate raw output if too long (over 200 lines)
  local line_count
  line_count=$(echo "$raw_output" | wc -l)
  if [ "$line_count" -gt 200 ]; then
    raw_output=$(echo "$raw_output" | head -100)
    raw_output="$raw_output\n\n... (truncated, $line_count lines total)"
  fi

  cat > "$DOC_DIR/$filename" << FILEEOF
# \`$adbjson_cmd\`

## adbjson

**Command:**
\`\`\`bash
$adbjson_cmd
\`\`\`

**Output:**
\`\`\`json
$json_output
\`\`\`

---

## adb

**Command:**
\`\`\`bash
adb $adb_cmd
\`\`\`

**Output:**
\`\`\`
$raw_output
\`\`\`
FILEEOF
}

generate_doc "adbjson version" "version"
generate_doc "adbjson shell getprop" "shell getprop"
generate_doc "adbjson shell dumpsys battery" "shell dumpsys battery"
generate_doc "adbjson shell dumpsys meminfo" "shell dumpsys meminfo"
generate_doc "adbjson shell dumpsys cpuinfo" "shell dumpsys cpuinfo"
generate_doc "adbjson shell dumpsys diskstats" "shell dumpsys diskstats"
generate_doc "adbjson shell dumpsys power" "shell dumpsys power"
generate_doc "adbjson shell dumpsys wifi" "shell dumpsys wifi"
generate_doc "adbjson shell dumpsys deviceidle" "shell dumpsys deviceidle"
generate_doc "adbjson shell dumpsys notification" "shell dumpsys notification"
generate_doc "adbjson shell dumpsys activity activities" "shell dumpsys activity activities"
generate_doc "adbjson shell dumpsys window displays" "shell dumpsys window displays"
generate_doc "adbjson shell wm size" "shell wm size"
generate_doc "adbjson shell wm density" "shell wm density"
generate_doc "adbjson shell service list" "shell service list"
generate_doc "adbjson shell pm list packages" "shell pm list packages"
generate_doc "adbjson shell pm list features" "shell pm list features"
generate_doc "adbjson shell pm list permissions" "shell pm list permissions"
generate_doc "adbjson shell settings list system" "shell settings list system"
generate_doc "adbjson shell settings list secure" "shell settings list secure"
generate_doc "adbjson shell settings list global" "shell settings list global"
generate_doc "adbjson shell dumpsys accessibility" "shell dumpsys accessibility"
generate_doc "adbjson shell dumpsys account" "shell dumpsys account"
generate_doc "adbjson shell dumpsys appwidget" "shell dumpsys appwidget"
generate_doc "adbjson shell dumpsys audio" "shell dumpsys audio"
generate_doc "adbjson shell dumpsys clipboard" "shell dumpsys clipboard"
generate_doc "adbjson shell dumpsys connectivity" "shell dumpsys connectivity"
generate_doc "adbjson shell dumpsys device_policy" "shell dumpsys device_policy"
generate_doc "adbjson shell dumpsys input_method" "shell dumpsys input_method"
generate_doc "adbjson shell dumpsys location" "shell dumpsys location"
generate_doc "adbjson shell dumpsys media" "shell dumpsys media"
generate_doc "adbjson shell dumpsys media_session" "shell dumpsys media_session"
generate_doc "adbjson shell dumpsys mount" "shell dumpsys mount"
generate_doc "adbjson shell dumpsys nfc" "shell dumpsys nfc"
generate_doc "adbjson shell dumpsys permission" "shell dumpsys permission"
generate_doc "adbjson shell dumpsys print" "shell dumpsys print"
generate_doc "adbjson shell dumpsys processinfo" "shell dumpsys processinfo"
generate_doc "adbjson shell dumpsys procstats" "shell dumpsys procstats"
generate_doc "adbjson shell dumpsys role" "shell dumpsys role"
generate_doc "adbjson shell dumpsys sensorservice" "shell dumpsys sensorservice"
generate_doc "adbjson shell dumpsys statusbar" "shell dumpsys statusbar"
generate_doc "adbjson shell dumpsys telecom" "shell dumpsys telecom"
generate_doc "adbjson shell dumpsys telephony" "shell dumpsys telephony"
generate_doc "adbjson shell dumpsys trust" "shell dumpsys trust"
generate_doc "adbjson shell dumpsys uimode" "shell dumpsys uimode"
generate_doc "adbjson shell dumpsys vibrator" "shell dumpsys vibrator"
generate_doc "adbjson shell dumpsys wallpaper" "shell dumpsys wallpaper"
generate_doc "adbjson shell cmd bluetooth_manager enable" "shell cmd bluetooth_manager enable"
generate_doc "adbjson shell cmd bluetooth_manager disable" "shell cmd bluetooth_manager disable"
generate_doc "adbjson shell dumpsys connectivity_native" "shell dumpsys connectivity_native"
generate_doc "adbjson shell dumpsys network_management" "shell dumpsys network_management"
generate_doc "adbjson shell dumpsys netpolicy" "shell dumpsys netpolicy"

echo ""
echo "Documentation generated in: $DOC_DIR"
ls -1 "$DOC_DIR"
