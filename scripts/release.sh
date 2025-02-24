#!/bin/bash

set -eu
tag=$(git describe --tags --abbrev=0)
output="dist/Unclipper_${tag}.zip"
mkdir -p dist
git archive --format=zip "${tag}" --output="${output}" --add-file="unclipper.exe" README.md LICENSE third-party-licenses/
sha256sum "${output}"
