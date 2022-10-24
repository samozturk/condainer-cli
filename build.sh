#! /bin/bash
GOOS=linux go build -o bin/rte-cli main.go;
GOOS=darwin go build -o bin/rte-cli-mac main.go;