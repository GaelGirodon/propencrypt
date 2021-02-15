#!/bin/sh

#
# package.sh
#
# Build and package propencrypt for Linux
#

mkdir -p dist

# Build
export GOARCH=amd64
export GOOS=linux
go build -ldflags="-s -w" -o dist/propencrypt cmd/propencrypt.go

# Package
pushd dist
archive=propencrypt.tar.gz
tar zcvf "$archive" propencrypt
sha256sum "$archive" | cut -d' ' -f 1 > "$archive.sha256"
popd
