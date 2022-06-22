#!/bin/bash

# set -euo pipefail

# . env.sh

# pushd service
# 	rm -f go.mod
# 	go mod init grsidecar
# 	go mod tidy
# 	go build -o $PROJ_ROOT/main plugin/server/main.go
# popd

CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o server plugin/server/main.go
GOARCH=amd64 GOOS=linux go build -o client plugin/client/main.go
GOARCH=amd64 GOOS=linux go build sanity.go
