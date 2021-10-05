#!/usr/bin/env sh

set -e

# initialize build variables
VERSION=1.0.0
GIT_HASH=$(git log -n1 --format=format:"%h")
DATE=$(date)

GOLDFLAGS="-X 'github.com/prometheus/common/version.Version=${VERSION}' -X 'github.com/prometheus/common/version.Branch=${GIT_HASH}' -X 'github.com/prometheus/common/version.BuildDate=${DATE}'"

set -x

# build
go build -o opengauss_exporter -ldflags "${GOLDFLAGS}" ./main.go
