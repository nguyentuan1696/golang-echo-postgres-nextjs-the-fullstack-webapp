#!/bin/sh

go mod tidy

set -e

BIN_DIR=${BIN_DIR:-/app/bin}
mkdir -p "$BIN_DIR"

files=`ls *.go`

echo "****************************************"
echo "******** building applications **********"
echo "****************************************"

for file in $files; do
	echo building $file
	go build -o "$BIN_DIR"/${file%.go} $file
done