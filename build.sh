#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

cd app

go test main_test.go main.go

GOOS=linux GOARCH=amd64 go build -o ../build/server-linux
GOOS=darwin GOARCH=amd64 go build -o ../build/server-darwin
cd ..
sh ./docker-publish.sh