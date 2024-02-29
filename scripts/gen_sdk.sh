#!/bin/bash

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &>/dev/null && pwd)"
PROJECT_DIR="$(dirname "${SCRIPT_DIR}")"

# generate protobuf related .go files
cd ${PROJECT_DIR}
rm -r doc/ || true
rm -r go/ || true
rm -r pycityproto/ || true
rm -r es/ || true
buf generate --include-imports --path city
buf generate --template buf.gen.tag.yaml --include-imports \
    --path city/agent/v2 \
    --path city/comm/input/v1 \
    --path city/economy/v1 \
    --path city/elec/input/v1 \
    --path city/event/v1 \
    --path city/map/v2 \
    --path city/routing/v2 \
    --path city/person/v1 \
    --path city/water/input/v1
protol --create-package --in-place --python-out pycityproto/ \
    protoc --proto-path=./ $(find . -name "*.proto")
