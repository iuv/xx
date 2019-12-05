#!/bin/bash
# linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s'
upx xx
mv xx build/linux/xx
# mac
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags '-w -s'
upx xx
mv xx build/mac/xx
