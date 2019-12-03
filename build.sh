#!/bin/bash
go build -ldflags '-w -s'
upx xx
