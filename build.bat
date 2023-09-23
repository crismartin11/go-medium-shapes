@echo off
cls
echo :: Building project.

set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
go build -o cmd/main cmd/main.go
%USERPROFILE%\Go\bin\build-lambda-zip.exe -o cmd/main.zip cmd/main

echo :: Build finished.
