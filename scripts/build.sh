#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
BIN_DIR="$ROOT_DIR/bin"
CACHE_DIR="$ROOT_DIR/.cache/go-build"

mkdir -p "$BIN_DIR"
mkdir -p "$CACHE_DIR"

cd "$ROOT_DIR"
export GOCACHE="$CACHE_DIR"

GOOS="${GOOS:-$(go env GOOS)}"
GOARCH="${GOARCH:-$(go env GOARCH)}"
output="adbjson"

if [ "$GOOS" = "windows" ]; then
  output="${output}.exe"
fi

go build \
  -o "$BIN_DIR/$output" \
  ./cmd/adbjson

echo "Built $BIN_DIR/$output ($GOOS/$GOARCH)"
