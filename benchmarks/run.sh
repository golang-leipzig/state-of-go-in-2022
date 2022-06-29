#!/usr/bin/env bash

set -euo pipefail

for x in $(seq 9 18); do
    version=1.$x
    sed "s/VERSION/$version-alpine/" Dockerfile.tmpl > Dockerfile
    docker build --quiet -t "bench:$version" . >/dev/null
    echo Benchmarking Go $version
    docker run --rm "bench:$version" go test -v -bench BenchmarkImageResizing . 2>&1 | grep BenchmarkImageResizing-
done