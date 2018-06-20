#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

go test -v lib/*

cd cmd

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../build/server-linux
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ../build/server-darwin

cd ..
