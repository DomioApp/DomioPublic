#!/usr/bin/env bash

#buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S'`
buildstamp=`date -u '+%s'`
hash=`git rev-parse --short HEAD`
version=`git tag -l --points-at HEAD`

export GOPATH=$PWD

echo
echo ---------------------------
echo "  Buildstamp: ${buildstamp}"
echo "  Hash:       ${hash}"
echo "  Version:    ${version}"
echo ---------------------------

echo
echo Compiling for Windows...
export GOARCH=amd64
export GOOS=windows
PROJECT_NAME=domio_public
go build -o /usr/local/bin/$PROJECT_NAME.exe -ldflags "-X main.Buildstamp=$buildstamp -X main.Hash=$hash  -X main.Version=$version" $PROJECT_NAME


echo Compiling for Linux...
export GOARCH=amd64
export GOOS=linux
go build -o /usr/local/bin/$PROJECT_NAME -ldflags "-X main.Buildstamp=$buildstamp -X main.Hash=$hash  -X main.Version=$version" $PROJECT_NAME
