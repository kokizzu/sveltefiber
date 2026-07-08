#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."

failures=0

fail() {
  echo "dependency security check failed: $*" >&2
  failures=$((failures + 1))
}

normalize_version() {
  local version="$1"
  version="${version#v}"
  version="${version%%+*}"
  echo "$version"
}

version_at_least() {
  local actual min first
  actual="$(normalize_version "$1")"
  min="$(normalize_version "$2")"
  first="$(printf '%s\n%s\n' "$min" "$actual" | sort -V | head -n 1)"
  [[ "$first" == "$min" ]]
}

check_min_version() {
  local label actual min
  label="$1"
  actual="$2"
  min="$3"
  if ! version_at_least "$actual" "$min"; then
    fail "$label is $actual, expected >= $min"
  fi
}

go_modules="$(mktemp)"
trap 'rm -f "$go_modules"' EXIT
go list -m all > "$go_modules"

go_version() {
  local module="$1"
  awk -v module="$module" '$1 == module { print $2; exit }' "$go_modules"
}

check_go_min() {
  local module min version
  module="$1"
  min="$2"
  version="$(go_version "$module")"
  if [[ -n "$version" ]]; then
    check_min_version "$module" "$version" "$min"
  fi
}

if grep -q '^github.com/docker/docker ' "$go_modules"; then
  fail "github.com/docker/docker is still in selected Go module graph"
fi

check_go_min github.com/docker/cli v29.6.1
check_go_min github.com/gofiber/fiber/v2 v2.52.13
check_go_min github.com/opencontainers/runc v1.3.6
check_go_min golang.org/x/crypto v0.53.0
check_go_min golang.org/x/image v0.18.0

if (( failures > 0 )); then
  exit 1
fi

echo "Dependency security check passed."
