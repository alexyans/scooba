#!/bin/bash

mkdir lib
cp /usr/lib/x86_64-linux-gnu/libgit2.* lib/
cp /usr/lib/x86_64-linux-gnu/libhttp_parser.* lib/
CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -a -o scooba .