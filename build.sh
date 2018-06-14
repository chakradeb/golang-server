#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

cd cmd

go test ./...

GOOS=linux GOARCH=amd64 go build -o ../build/server-linux
GOOS=darwin GOARCH=amd64 go build -o ../build/server-darwin
cd ..
sh ./docker-publish.sh