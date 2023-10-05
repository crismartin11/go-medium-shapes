package handler

import (
	"context"
	"go-medium-shapes/pkg/models"
	"go-medium-shapes/pkg/validator"

	"github.com/rs/zerolog/log"
)

type Processor interface {
	ProcessCreation(request models.Request) (models.Response, error)
	ProcessGeneration(ctx context.Context, request models.Request) (models.Response, error)
}

type Handler interface {
	Handle(ctc context.Context, request models.Request) (models.Response, error)
}

type GenerateFileHandler struct {
	p Processor
	v validator.Validator
}

func New(p Processor, v validator.Validator) Handler {
	return &GenerateFileHandler{
		p,
		v,
	}
}

func (h GenerateFileHandler) Handle(ctx context.Context, request models.Request) (models.Response, error) {

	log.Info().Msg("Handle started")
	err := h.v.ValidateRequest(request)
	if err != nil {
		log.Error().Msg("Handle. Error in ValidateRequest.")
		return models.NewResponseError(400, err.Error())
	}

	if request.IsValidData() {
		return h.p.ProcessCreation(request)
	}

	return h.p.ProcessGeneration(ctx, request)
}
