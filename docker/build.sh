#!/bin/bash
rm DevServer
CGO_ENABLED=0 GOOS=linux go build -o DevServer ../cmd/server

docker build -t dev-server:latest .
rm DevServer