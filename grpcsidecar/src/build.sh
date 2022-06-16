#!/bin/bash

# set -euo pipefail

# . env.sh

# pushd service
# 	rm -f go.mod
# 	go mod init grsidecar
# 	go mod tidy
# 	go build -o $PROJ_ROOT/main plugin/server/main.go
# popd

go build plugin/server/main.go 