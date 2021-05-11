#!/bin/bash
set -e

go run cmd/dbmigrate/main.go

go run cmd/dbmigrate/main.go -dbname=goapidesigntest

GO111MODULE=off go get github.com/githubnemo/CompileDaemon

CompileDaemon --build="go build -o main cmd/api/main.go" --command=./main