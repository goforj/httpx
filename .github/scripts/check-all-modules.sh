#!/usr/bin/env bash
set -euo pipefail
module_count=0
while IFS= read -r -d '' module; do
  module_dir="${module%/go.mod}"
  [[ -n "${module_dir}" ]] || module_dir="."
  (
    cd "${module_dir}"
    mapfile -t packages < <(GOWORK=off go list ./...)
    [[ "${#packages[@]}" -eq 0 ]] || GOWORK=off go run golang.org/x/vuln/cmd/govulncheck@v1.7.0 -test "${packages[@]}"
  ) < /dev/null
  module_count=$((module_count + 1))
done < <(find . -type f -name go.mod -not -path './.git/*' -not -path './node_modules/*' -not -path './vendor/*' -print0)

[[ "${module_count}" -gt 0 ]]
