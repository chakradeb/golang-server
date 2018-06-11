#!/usr/bin/env bash

docker build -t golang-server .
docker tag golang-server debarc/golang-server
docker push debarc/golang-server