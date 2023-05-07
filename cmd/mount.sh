#!/bin/bash

rm main main.zip

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go

zip main.zip main

aws lambda update-function-code --function-name main --zip-file fileb://main.zip
