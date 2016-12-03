#!/bin/bash
set -euo pipefail

HOF_DIR="${HOME}/.hofstadter"

mkdir ${HOF_DIR}
ln -s $(pwd)/dsl ${HOF_DIR}/dsl

go install

