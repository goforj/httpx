#!/usr/bin/env bash
set -euo pipefail
root_dir="$(pwd -P)"
output_dir="${1:?output directory is required}"
mkdir -p "${output_dir}"
module_count=0
lockfile_count=0
while IFS= read -r -d '' module; do
  module_dir="${module%/go.mod}"
  [[ -n "${module_dir}" ]] || module_dir="."
  output="${output_dir}/go-${module_count}.cdx.json"
  output_path="${root_dir}/${output}"
  [[ "${output}" = /* ]] && output_path="${output}"
  (cd "${module_dir}" && GOWORK=off go run github.com/CycloneDX/cyclonedx-gomod/cmd/cyclonedx-gomod@v1.9.0 mod -output "${output_path}")
  module_count=$((module_count + 1))
done < <(find . -type f -name go.mod -not -path './vendor/*' -print0)
while IFS= read -r -d '' lockfile; do lockfile_count=$((lockfile_count + 1)); done < <(find . -type f -name package-lock.json -not -path './node_modules/*' -print0)
[[ "${module_count}" -eq "$(find "${output_dir}" -type f -name 'go-*.cdx.json' -print | wc -l | tr -d ' ')" ]]
printf 'Go module SBOMs: %s\npackage-lock files: %s\n' "${module_count}" "${lockfile_count}"
