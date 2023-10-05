package main

import (
	"go-medium-shapes/internal/processor"
	"go-medium-shapes/internal/repository"
	"go-medium-shapes/pkg/handler"
	"go-medium-shapes/pkg/validator"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	r := repository.New()
	p := processor.New(r)
	v := validator.NewItemValidator()
	h := handler.New(p, v)

	lambda.Start(h.Handle)
}
