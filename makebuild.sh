#!/bin/sh

go env -w GOPROXY="https://goproxy.cn,direct"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app
