#!/bin/sh

echo Building project.

lambdaName="goMediumShapes"
directory="package"

cd ./cmd
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $lambdaName main.go
zip -rm ../build/$directory/$lambdaName.zip $lambdaName

echo Build finished.