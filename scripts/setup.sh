#!/bin/bash
set -euo pipefail

HOF_DIR="${HOME}/.geb"

mkdir ${HOF_DIR}
ln -s $(pwd)/dsl ${HOF_DIR}/dsl
cp $(pwd)/scripts/dotfolder/geb.yaml ${HOF_DIR}/geb.yaml

go build
sudo ln -s $(pwd)/geb /usr/local/bin/geb


