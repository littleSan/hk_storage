#! /bin/bash

PWD="$(cd `dirname $0`; pwd)"
NAME="hk_storageServer"


CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./$NAME

echo "Success!"