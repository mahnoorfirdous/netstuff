#!/bin/bash

set -euo pipefail

. env.sh

pushd src/service
	rm -f go.mod
	go mod init service
	go mod tidy
	go build -o $PROJ_ROOT/service service//main.go
popd
