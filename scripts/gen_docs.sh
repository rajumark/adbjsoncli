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
generate_doc "adbjson shell dumpsys battery" "shell dumpsys battery"
generate_doc "adbjson shell pm list packages" "shell pm list packages"

echo ""
echo "Documentation generated in: $DOC_DIR"
ls -1 "$DOC_DIR"
