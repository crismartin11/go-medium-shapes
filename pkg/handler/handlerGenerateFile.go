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
)

func HandlerGenerateFile(ctx context.Context, event models.Item) (models.Response, error) {

	dynamoDBClient := dynamodb.NewDynamoDBClient()
	listShapes, err := dynamoDBClient.ListShapesByType(event.ShapeType)
	if err != nil {
		return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
	}

	var shapes = []models.IShape{}
	for _, item := range listShapes {
		elem, err := models.ShapeFactory(item.Id, item.ShapeType, item.A, item.B)
		if err != nil {
			return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
		}
		shapes = append(shapes, elem)
	}

	path, err := utils.GenerateTempFile(shapes)
	if err != nil {
		return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
	} else {
		err = utils.UploadTempFile(path, getFileName(ctx, event))
		if err != nil {
			return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
		}
	}

	return models.NewResponseOk("Generation file process succesful!")
}

func getFileName(ctx context.Context, event models.Item) string {
	lc, _ := lambdacontext.FromContext(ctx)
	return event.ShapeType + "-" + lc.AwsRequestID + "-" + time.Now().Format(constants.DATE_FORMAT)
}
