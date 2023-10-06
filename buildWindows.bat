@echo off
cls
echo :: Building project.

set lambdaName="goMediumShapes"
set directory="package"

set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
go build -o build/%directory%/%lambdaName% cmd/main.go
%USERPROFILE%\Go\bin\build-lambda-zip.exe -o build/%directory%/%lambdaName%.zip build/%directory%/%lambdaName%

echo :: Build finished.
