#!/usr/bin/env bash
set -euo pipefail
while IFS= read -r -d '' module; do
  module_dir="${module%/go.mod}"
  [[ -n "${module_dir}" ]] || module_dir="."
  (
    cd "${module_dir}"
    packages="$(GOWORK=off go list ./...)"
    [[ -z "${packages}" ]] || GOWORK=off go run golang.org/x/vuln/cmd/govulncheck@v1.7.0 -test ${packages}
  )
done < <(find . -type f -name go.mod -not -path './.git/*' -not -path './node_modules/*' -not -path './vendor/*' -print0)
