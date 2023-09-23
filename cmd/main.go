package main

import (
	"go-medium-shapes/pkg/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	handler := handler.New()
	lambda.Start(handler.Handle)
}
