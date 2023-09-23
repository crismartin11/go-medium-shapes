package handler

import (
	"context"
	"fmt"
	"go-medium-shapes/pkg/models"
)

type Handler struct{}

func New() Handler {
	return Handler{}
}

func (h Handler) Handle(ctc context.Context, event models.Item) (models.Response, error) {

	if !event.IsValidShapeType() {
		return models.NewResponseError(400, fmt.Sprintf("ERROR: Tipo de figura %s inválido.", event.ShapeType))
	}

	if event.IsValidData() {
		return HandlerCreateShape(ctc, event)
	}

	return HandlerGenerateFile(ctc, event)
}
