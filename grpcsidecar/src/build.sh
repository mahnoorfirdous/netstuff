#!/bin/bash

# set -euo pipefail

# . env.sh

# pushd service
# 	rm -f go.mod
# 	go mod init grsidecar
# 	go mod tidy
# 	go build -o $PROJ_ROOT/main plugin/server/main.go
# popd

GOARCH=amd64 GOOS=linux go build plugin/server/main.go
