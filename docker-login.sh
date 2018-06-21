#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

echo $DOCKER_USER
docker login -u $DOCKER_USER -p $DOCKER_PASS