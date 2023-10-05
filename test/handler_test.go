package test

import (
	"context"
	"go-medium-shapes/internal/processor"
	"go-medium-shapes/internal/repository"
	"go-medium-shapes/pkg/handler"
	"go-medium-shapes/pkg/models"
	"go-medium-shapes/pkg/validator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandle(t *testing.T) {

	ctx := context.Background()

	t.Run("HandlerGenerateFile", func(t *testing.T) {

		r := repository.New()
		p := processor.New(r)
		v := validator.NewItemValidator()
		h := handler.New(p, v)
		response, err := h.Handle(ctx, models.Request{ShapeType: "RECTANGLE"})

		assert.NoError(t, err)
		assert.Equal(t, "Generation file process successful!", response.Body)
	})

	t.Run("HandlerCreateShape", func(t *testing.T) {
		r := repository.New()
		p := processor.New(r)
		v := validator.NewItemValidator()
		h := handler.New(p, v)
		response, err := h.Handle(ctx, models.Request{ShapeType: "RECTANGLE", A: 1, B: 2, Id: "7"})

		assert.NoError(t, err)
		assert.Equal(t, "Creation process successful!", response.Body)
	})

}
