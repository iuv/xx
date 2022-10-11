#!/bin/bash
# linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s'
upx xx
mv xx build/linux/x86/xx
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags '-w -s'
upx xx
mv xx build/linux/arm64/xx
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags '-w -s'
upx xx
mv xx build/linux/arm/xx
# mac
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags '-w -s'
upx xx
mv xx build/mac/x86/xx
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags '-w -s'
upx xx
mv xx build/mac/arm64/xx
