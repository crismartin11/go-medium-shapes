@echo off
cls
echo :: Building project.

set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
go build -o build/main cmd/main.go
%USERPROFILE%\Go\bin\build-lambda-zip.exe -o build/main.zip build/main

echo :: Build finished.
