#!/bin/bash

ROOTDIR=$(pwd)
V=$(git describe --abbrev=0 --tags)
cd ${ROOTDIR}/server

docker image build -t danielbok/pcr-server:${V} .
docker image tag danielbok/pcr-server:${V} danielbok/pcr-server:latest

cd ${ROOTDIR}/web
docker image build -t danielbok/pcr-web:${V} .
docker image build -t danielbok/pcr-web-ssl:${V} -f ssl.Dockerfile .

docker image tag danielbok/pcr-web:${V} danielbok/pcr-web:latest
docker image tag danielbok/pcr-web-ssl:${V} danielbok/pcr-web-ssl:latest

cd ${ROOTDIR}/cli
echo "Building CLI tool for windows"
GOOS=windows GOARCH=amd64 go build -o pcr.exe .

docker image push danielbok/pcr-server:${V}
docker image push danielbok/pcr-server:latest
docker image push danielbok/pcr-web:${V}
docker image push danielbok/pcr-web:latest
docker image push danielbok/pcr-web-ssl:${V}
docker image push danielbok/pcr-web-ssl:latest
