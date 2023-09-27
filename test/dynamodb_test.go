package test

import (
	"go-medium-shapes/pkg/dynamodb"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamoDB(t *testing.T) {
	t.Run("ListShapesByType", func(t *testing.T) {
		dynamoDBClient := dynamodb.NewDynamoDBClient()
		listShapes, err := dynamoDBClient.ListShapesByType("RECTANGLE")

		assert.NoError(t, err)
		assert.NotEmpty(t, listShapes, "Lista no debería estar vacía.")
	})

	t.Run("ListShapesByType empty", func(t *testing.T) {
		dynamoDBClient := dynamodb.NewDynamoDBClient()
		listShapes, err := dynamoDBClient.ListShapesByType("AA")

		assert.NoError(t, err)
		assert.Empty(t, listShapes, "La lista debería estar vacía.")
	})
}
