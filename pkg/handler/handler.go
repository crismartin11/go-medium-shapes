package handler

import (
	"context"
	"fmt"
	"go-medium-shapes/pkg/models"

	"github.com/rs/zerolog/log"
)

type Handler struct{}

func New() Handler {
	return Handler{}
}

func (h Handler) Handle(ctc context.Context, event models.Item) (models.Response, error) {

	log.Info().Msg("Handle Shape started")
	if !event.IsValidShapeType() {
		log.Error().Msg("Handle Shape. Invalid shape type.")
		return models.NewResponseError(400, fmt.Sprintf("ERROR: Tipo de figura %s inválido.", event.ShapeType))
	}

	if event.IsValidData() {
		return HandlerCreateShape(ctc, event)
	}

	return HandlerGenerateFile(ctc, event)
}
