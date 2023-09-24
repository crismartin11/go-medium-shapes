package handler

import (
	"context"
	"fmt"

	"go-medium-shapes/pkg/dynamodb"
	"go-medium-shapes/pkg/models"
	"go-medium-shapes/pkg/utils"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func HandlerCreateShape(ctc context.Context, event models.Item) (models.Response, error) {

	log.Info().Msg("Create shape handler started (HandlerCreateShape).")

	user, err := utils.HttpGetUserData(event.Id)
	if err != nil {
		log.Error().Msg("Error getting user data from API.")
		return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
	} else if user.Data.Email == "" {
		log.Error().Msg("Error getting email user from data.")
		return models.NewResponseError(400, fmt.Sprintf("HandlerCreateShape. No se obtuvo el email del usuario con id %s", event.Id))
	}

	uuid, _ := uuid.NewUUID()
	dynamoDBClient := dynamodb.NewDynamoDBClient()
	err = dynamoDBClient.Create(uuid.String(), event.ShapeType, event.A, event.B, user.Data.Email)
	if err != nil {
		log.Error().Msg("Error creating shape in DynamoDB.")
		return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
	}

	log.Info().Msg("Create shape handler finished successfully (HandlerCreateShape).")
	return models.NewResponseOk("Creation process successful!")
}
