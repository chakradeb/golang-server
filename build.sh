#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

cd app
GOOS=linux GOARCH=amd64 go build -o ../build/server-linux
GOOS=darwin GOARCH=amd64 go build -o ../build/server-darwin
sh ./docker-publish.sh