@echo off
cls
echo :: Building project.

set lambdaName="goMediumShapes"

set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
go build -o build/%lambdaName% cmd/main.go
%USERPROFILE%\Go\bin\build-lambda-zip.exe -o build/%lambdaName%.zip build/%lambdaName%

echo :: Build finished.
