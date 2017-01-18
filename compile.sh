#!/usr/bin/env bash

PROJECT_NAME=domio_public

#buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S'`
buildstamp=`date -u '+%s'`
hash=`git rev-parse --short HEAD`
version=`git tag -l --points-at HEAD`


#export GOPATH=${PWD}

echo
echo ---------------------------
echo "  Buildstamp: ${buildstamp}"
echo "  Hash:       ${hash}"
echo "  Version:    ${version}"
echo ---------------------------
echo

platform='unknown'
unamestr=`uname -s`

if [[ "$unamestr" == 'CYGWIN_NT-10.0' ]]; then
   platform='cygwin'
elif [[ "$unamestr" == 'Darwin' ]]; then
   platform='mac'
elif [[ "$unamestr" == 'Linux' ]]; then
   platform='linux'
fi

if [ $platform == "cygwin" ]
    then
        echo Compiling for Windows...
        export GOPATH=C:/Users/sbasharov/WebstormProjects/DomioPublic
        qtc -dir ${GOPATH}/src/domio_public/templates
        echo Output folder: /usr/local/bin/${PROJECT_NAME}_win.exe
        export GOARCH=amd64
        export GOOS=windows
        go build -o /usr/local/bin/${PROJECT_NAME}_win.exe -ldflags "-X main.Buildstamp=$buildstamp -X main.Hash=$hash  -X main.Version=$version" ${PROJECT_NAME}
fi

if [ $platform == "mac" ]
    then
        echo Compiling for Mac...
        echo Output folder: /usr/local/bin/${PROJECT_NAME}_mac
        export GOARCH=amd64
        export GOOS=darwin
        go build -o /usr/local/bin/${PROJECT_NAME}_mac -ldflags "-X main.Buildstamp=$buildstamp -X main.Hash=$hash  -X main.Version=$version" ${PROJECT_NAME}
fi

if [ $platform == "linux" ]
    then
        echo Compiling for Linux...
        export GOARCH=amd64
        export GOOS=linux
        echo Output folder: /usr/local/bin/${PROJECT_NAME}_linux
        go build -o /usr/local/bin/${PROJECT_NAME}_linux -ldflags "-X main.Buildstamp=$buildstamp -X main.Hash=$hash  -X main.Version=$version" ${PROJECT_NAME}
fi
