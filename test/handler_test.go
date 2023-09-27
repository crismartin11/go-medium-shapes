package test

import (
	"context"
	"go-medium-shapes/pkg/handler"
	"go-medium-shapes/pkg/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandle(t *testing.T) {

	ctx := context.Background()

	t.Run("HandlerGenerateFile", func(t *testing.T) {
		hand := handler.New()
		response, err := hand.Handle(ctx, models.Item{ShapeType: "ELLIPSE"})

		assert.NoError(t, err)
		assert.Equal(t, "Generation file process successful!", response.Body)
	})

	t.Run("HandlerCreateShape", func(t *testing.T) {
		hand := handler.New()
		response, err := hand.Handle(ctx, models.Item{ShapeType: "RECTANGLE", A: 1, B: 2, Id: "2"})

		assert.NoError(t, err)
		assert.Equal(t, "Creation process successful!", response.Body)
	})

}
