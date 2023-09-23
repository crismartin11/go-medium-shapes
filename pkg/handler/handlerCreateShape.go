package handler

import (
	"context"
	"fmt"

	"go-medium-shapes/pkg/dynamodb"
	"go-medium-shapes/pkg/models"
	"go-medium-shapes/pkg/utils"

	"github.com/google/uuid"
)

func HandlerCreateShape(ctc context.Context, event models.Item) (models.Response, error) {

	user, err := utils.HttpGetUserData(event.Id)
	if err != nil {
		return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
	} else if user.Data.Email == "" {
		return models.NewResponseError(400, fmt.Sprintf("HandlerCreateShape. No se obtuvo el email del usuario con id %s", event.Id))
	}

	uuid, _ := uuid.NewUUID()
	dynamoDBClient := dynamodb.NewDynamoDBClient()
	err = dynamoDBClient.Create(uuid.String(), event.ShapeType, event.A, event.B, user.Data.Email)
	if err != nil {
		return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
	}

	return models.NewResponseOk("Creation process succesful!")
}
