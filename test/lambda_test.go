package lambda_test

import (
	"context"
	"go-medium-shapes/pkg/handler"
	"go-medium-shapes/pkg/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandle(t *testing.T) {

	ctx := context.Background()

	t.Run("Test handle", func(t *testing.T) {
		hand := handler.New()
		//response, err := hand.Handle(ctx, models.Item{ShapeType: "ELLIPSE"})
		response, err := hand.Handle(ctx, models.Item{ShapeType: "ELLIPSE", A: 1, B: 2, Id: "2"})

		assert.NoError(t, err)
		assert.Equal(t, "ShapeType ELLIPSE!", response.Body)

	})

}
