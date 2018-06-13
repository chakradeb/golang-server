#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

docker build -t golang-server .
docker tag golang-server debarc/golang-server
docker push debarc/golang-server