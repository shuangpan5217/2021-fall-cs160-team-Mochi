#!/bin/bash
set -ex
scriptdir="$( cd "$(dirname "$0")" ; pwd -P )"
echo "Building go-swagger"
swagger_bin_dir="${scriptdir}/../bin"

mkdir -p $swagger_bin_dir
swagger_bin_path="${swagger_bin_dir}/go-swagger"
rm -rf $swagger_bin_path

download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | \
  jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url')
curl -o $swagger_bin_path -L'#' "$download_url"
chmod +x $swagger_bin_path