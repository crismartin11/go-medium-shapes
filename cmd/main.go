package main

import (
	"go-medium-shapes/internal/processor"
	"go-medium-shapes/internal/repository"
	"go-medium-shapes/pkg/handler"
	"go-medium-shapes/pkg/services"
	"go-medium-shapes/pkg/validator"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	r := repository.New()
	s3r := repository.NewS3Repository()
	us := services.NewUserDataService(&http.Client{})
	p := processor.New(r, s3r, us)
	v := validator.NewItemValidator()
	h := handler.New(p, v)

	lambda.Start(h.Handle)
}
