#!/usr/bin/env bash
set -u

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
  json_output=$(timeout 30 "$BIN" $adb_cmd 2>&1 || true)

  local raw_output=""
  if [ -n "$ADB" ]; then
    # shellcheck disable=SC2086
    raw_output=$(timeout 30 "$ADB" $adb_cmd 2>&1 || true)
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
generate_doc "adbjson shell cmd connectivity airplane-mode" "shell cmd connectivity airplane-mode"
generate_doc "adbjson shell dumpsys device_config" "shell dumpsys device_config"
generate_doc "adbjson shell dumpsys shortcut" "shell dumpsys shortcut"
generate_doc "adbjson shell dumpsys settings" "shell dumpsys settings"
generate_doc "adbjson shell dumpsys network_time_update_service" "shell dumpsys network_time_update_service"
generate_doc "adbjson shell dumpsys safety_center" "shell dumpsys safety_center"
generate_doc "adbjson shell dumpsys voiceinteraction" "shell dumpsys voiceinteraction"
generate_doc "adbjson shell dumpsys webviewupdate" "shell dumpsys webviewupdate"
generate_doc "adbjson shell dumpsys network_score" "shell dumpsys network_score"
generate_doc "adbjson shell dumpsys translation" "shell dumpsys translation"
generate_doc "adbjson shell dumpsys search" "shell dumpsys search"
generate_doc "adbjson shell uptime" "shell uptime"
generate_doc "adbjson shell uname -a" "shell uname -a"
generate_doc "adbjson shell id" "shell id"
generate_doc "adbjson shell df" "shell df"
generate_doc "adbjson shell free" "shell free"
generate_doc "adbjson shell ps" "shell ps"
generate_doc "adbjson shell printenv" "shell printenv"
generate_doc "adbjson shell ifconfig" "shell ifconfig"
generate_doc "adbjson shell netstat" "shell netstat"
generate_doc "adbjson shell vmstat" "shell vmstat"
generate_doc "adbjson shell groups" "shell groups"
generate_doc "adbjson shell whoami" "shell whoami"
generate_doc "adbjson shell env" "shell env"
generate_doc "adbjson shell ls" "shell ls -l /"
generate_doc "adbjson shell cat /proc/version" "shell cat /proc/version"
generate_doc "adbjson shell cat /proc/cpuinfo" "shell cat /proc/cpuinfo"
generate_doc "adbjson shell cmd battery reset" "shell cmd battery reset"
generate_doc "adbjson shell cmd deviceidle step" "shell cmd deviceidle step"
generate_doc "adbjson shell cmd deviceidle force-idle" "shell cmd deviceidle force-idle"
generate_doc "adbjson shell cmd deviceidle unforce" "shell cmd deviceidle unforce"
generate_doc "adbjson shell dumpsys app_binding" "shell dumpsys app_binding"
generate_doc "adbjson shell dumpsys content" "shell dumpsys content"
generate_doc "adbjson shell dumpsys dbinfo" "shell dumpsys dbinfo"
generate_doc "adbjson shell dumpsys hardware_properties" "shell dumpsys hardware_properties"
generate_doc "adbjson shell dumpsys launcher" "shell dumpsys launcher"
generate_doc "adbjson shell dumpsys locksettings" "shell dumpsys locksettings"
generate_doc "adbjson shell dumpsys media_metrics" "shell dumpsys media_metrics"
generate_doc "adbjson shell dumpsys media_resource_monitor" "shell dumpsys media_resource_monitor"
generate_doc "adbjson shell dumpsys media_projection" "shell dumpsys media_projection"
generate_doc "adbjson shell dumpsys media_communication" "shell dumpsys media_communication"
generate_doc "adbjson shell dumpsys package_native" "shell dumpsys package_native"
generate_doc "adbjson shell dumpsys permission_checker" "shell dumpsys permission_checker"
generate_doc "adbjson shell dumpsys permissionmgr" "shell dumpsys permissionmgr"
generate_doc "adbjson shell dumpsys pinner" "shell dumpsys pinner"
generate_doc "adbjson shell dumpsys platform_compat" "shell dumpsys platform_compat"
generate_doc "adbjson shell dumpsys platform_compat_native" "shell dumpsys platform_compat_native"
generate_doc "adbjson shell dumpsys recovery" "shell dumpsys recovery"
generate_doc "adbjson shell dumpsys runtime" "shell dumpsys runtime"
generate_doc "adbjson shell dumpsys scheduling_policy" "shell dumpsys scheduling_policy"
generate_doc "adbjson shell dumpsys sdk_sandbox" "shell dumpsys sdk_sandbox"
generate_doc "adbjson shell dumpsys serial" "shell dumpsys serial"
generate_doc "adbjson shell dumpsys storaged" "shell dumpsys storaged"
generate_doc "adbjson shell dumpsys storaged_pri" "shell dumpsys storaged_pri"
generate_doc "adbjson shell dumpsys storagestats" "shell dumpsys storagestats"
generate_doc "adbjson shell dumpsys system_config" "shell dumpsys system_config"
generate_doc "adbjson shell dumpsys testing" "shell dumpsys testing"
generate_doc "adbjson shell dumpsys textclassification" "shell dumpsys textclassification"
generate_doc "adbjson shell dumpsys textservices" "shell dumpsys textservices"
generate_doc "adbjson shell dumpsys thermalservice" "shell dumpsys thermalservice"
generate_doc "adbjson shell dumpsys updatelock" "shell dumpsys updatelock"
generate_doc "adbjson shell dumpsys usagestats" "shell dumpsys usagestats"
generate_doc "adbjson shell dumpsys virtualdevice" "shell dumpsys virtualdevice"
generate_doc "adbjson shell dumpsys wifi_aware" "shell dumpsys wifi_aware"
generate_doc "adbjson shell dumpsys wifi_rtt" "shell dumpsys wifi_rtt"
generate_doc "adbjson shell dumpsys wifi_scanner" "shell dumpsys wifi_scanner"
generate_doc "adbjson shell dumpsys batteryproperties" "shell dumpsys batteryproperties"

echo ""
echo "Documentation generated in: $DOC_DIR"
ls -1 "$DOC_DIR"
