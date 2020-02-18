#!/usr/bin/env bash

env GOOS=windows GOARCH=386 go build -o ./bin/386/service.exe main.go

env GOOS=windows GOARCH=amd64 go build -o ./bin/amd64/service.exe main.go
