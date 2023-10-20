package repository

import (
	"context"
	"fmt"
	"go-medium-shapes/pkg/constants"
	"go-medium-shapes/pkg/models"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

type DynamoDB struct {
	client *dynamodb.Client
}

type IDynamoDB interface {
	ListShapesByType(shapeType string) ([]models.Request, error)
	CreateShape(id string, shapeType string, a float64, b float64, creator string) error
}

func NewDynamoDB() IDynamoDB {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("Configuration error, " + err.Error())
	}

	return DynamoDB{client: dynamodb.NewFromConfig(cfg)}
}

func (d DynamoDB) ListShapesByType(shapeType string) ([]models.Request, error) {
	shapes := []models.Request{}

	output, err := d.client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(constants.SHAPE_TABLE_NAME),
		IndexName:              aws.String("tipo-index"),
		KeyConditionExpression: aws.String("tipo = :shapeType"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":shapeType": &types.AttributeValueMemberS{
				Value: shapeType,
			},
		},
	})

	if err != nil {
		return shapes, fmt.Errorf("ListShapesByType. Error al invocar API Query. %s", err)
	}
	if len(output.Items) == 0 {
		fmt.Println("ListShapesByType. No se encontraron shapes")
		return shapes, nil
	}

	err = attributevalue.UnmarshalListOfMaps(output.Items, &shapes)
	if err != nil {
		return shapes, fmt.Errorf("ListShapesByType. Error al parsear items. %s", err)
	}

	return shapes, nil
}

func (d DynamoDB) CreateShape(id string, shapeType string, a float64, b float64, creator string) error {
	shape := models.Request{Id: id, ShapeType: shapeType, A: a, B: b, Creator: creator}

	av, err := attributevalue.MarshalMap(shape)
	if err != nil {
		return fmt.Errorf("Create. Error de parseo. %s", err)
	}

	// Creo el ítem a insertar en la tabla Users
	input := &dynamodb.PutItemInput{
		TableName: aws.String(constants.SHAPE_TABLE_NAME),
		Item:      av,
	}

	// Inserto
	_, err = d.client.PutItem(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("Create. Error insertando figura (%s). %s", id, err)
	}

	return nil
}
