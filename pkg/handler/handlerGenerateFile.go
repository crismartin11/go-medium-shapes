package handler

import (
	"context"
	"fmt"
	"time"

	"go-medium-shapes/pkg/constants"
	"go-medium-shapes/pkg/dynamodb"
	"go-medium-shapes/pkg/models"
	"go-medium-shapes/pkg/utils"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/rs/zerolog/log"
)

func HandlerGenerateFile(ctx context.Context, event models.Item) (models.Response, error) {

	log.Info().Msg("(HandlerGenerateFile) Generate shape list file handler started.")

	log.Info().Msg("(HandlerGenerateFile) Getting list from dynamoDB.")
	dynamoDBClient := dynamodb.NewDynamoDBClient()
	listShapes, err := dynamoDBClient.ListShapesByType(event.ShapeType)
	if err != nil {
		log.Error().Msg("(HandlerGenerateFile) Error getting shape list from DynamoDB.")
		return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
	}

	log.Info().Msg("(HandlerGenerateFile) Generating shapes.")
	var shapes = []models.IShape{}
	for _, item := range listShapes {
		elem, err := models.ShapeFactory(item.Id, item.ShapeType, item.A, item.B)
		if err != nil {
			log.Error().Msg("(HandlerGenerateFile) Error making shape based on dynamo item.")
			return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
		}
		shapes = append(shapes, elem)
	}

	log.Info().Msg("(HandlerGenerateFile) Generating file.")
	path, err := utils.GenerateTempFile(shapes)
	if err != nil {
		log.Error().Msg("(HandlerGenerateFile) Error generating file.")
		return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
	} else {
		log.Info().Msg("(HandlerGenerateFile) Uploading file " + path + ".")
		err = utils.UploadTempFile(path, getFileName(ctx, event))
		if err != nil {
			log.Error().Msg("(HandlerGenerateFile) Error uploading file.")
			return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
		}
	}

	log.Info().Msg("(HandlerGenerateFile) Generate shape list file handler finished successfully.")
	return models.NewResponseOk("Generation file process successful!")
}

func getFileName(ctx context.Context, event models.Item) string {
	lc, ok := lambdacontext.FromContext(ctx)
	awsRequestID := "unknown"
	if ok {
		awsRequestID = lc.AwsRequestID
	}
	return event.ShapeType + "-" + awsRequestID + "-" + time.Now().Format(constants.DATE_FORMAT) + ".txt"
}
