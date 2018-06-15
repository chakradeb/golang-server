#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin

docker push debarc/golang-server