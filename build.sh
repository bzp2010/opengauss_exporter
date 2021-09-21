#!/usr/bin/env bash

set -e

# initialize Var
VERSION=1.0.0
GIT_HASH=$(HASH="ref: HEAD"; while [[ $HASH == ref\:* ]]; do HASH="$(cat ".git/$(echo $HASH | cut -d \  -f 2)")"; done; echo ${HASH:0:7})
DATE=$(date)

GOLDFLAGS="-X 'github.com/prometheus/common/version.Version=${VERSION}' -X 'github.com/prometheus/common/version.Branch=${GIT_HASH}' -X 'github.com/prometheus/common/version.BuildDate=${DATE}'"

set -x

# build
go build -o ./opengauss_exporter -ldflags "${GOLDFLAGS}" ./main.go

echo "Build the openGauss Exporter successful"
