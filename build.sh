#!/bin/bash

export PACKAGE_NAME=$(grep module go.mod | awk '{print $2}' | sed 's/\//_/g')

# Toolchain variables
export TOOLCHAIN_BIN_DIR="<bin dir>"
export GOOS=linux
export GOARCH=arm
export GOARM=7
export GOMIPS=softfloat
export CC="${TOOLCHAIN_BIN_DIR}/arm-openipc-linux-musleabi-gcc"
export STRIP="${TOOLCHAIN_BIN_DIR}/arm-openipc-linux-musleabi-strip"
export CGO_ENABLED=1

go build -o "${PACKAGE_NAME}"

${STRIP} "${PACKAGE_NAME}"
