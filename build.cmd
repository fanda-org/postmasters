@echo off
echo Building application: go build -ldflags="-s -w"
go build -ldflags="-s -w"
echo Building application completed!