package repository

import (
	"fmt"
	"go-medium-shapes/pkg/constants"
	"go-medium-shapes/pkg/credentials"
	"go-medium-shapes/pkg/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/rs/zerolog/log"
)

type Repository struct {
	service *dynamodb.DynamoDB
}

func New() Repository {

	log.Info().Msg("New Repository. Getting credentials.")
	service, err := credentials.GetClientDynamo()
	if err != nil {
		log.Fatal().Msg("Error creating repository." + err.Error())
	}

	return Repository{
		service: service,
	}
}

func (r Repository) ListShapesByType(shapeType string) ([]models.Request, error) {
	shapes := []models.Request{}

	// Con la proyección obtengo el id, tipo, a, b y creador de cada elemento recuperado. Impotante: name de la DB, no del modelo (por eso en minúscula)
	proj := expression.NamesList(expression.Name("id"), expression.Name("tipo"), expression.Name("a"), expression.Name("b"), expression.Name("creador"))

	// Creo la expresión
	expr, err := expression.NewBuilder().WithFilter(
		expression.Contains(expression.Name("tipo"), shapeType),
	).WithProjection(proj).Build()
	if err != nil {
		return shapes, fmt.Errorf("ListShapesByType. Error creando expression. %s", err)
	}

	// Creo el objeto de parámetros de entrada
	params := &dynamodb.ScanInput{
		TableName:                 aws.String(constants.SHAPE_TABLE_NAME),
		FilterExpression:          expr.Filter(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ProjectionExpression:      expr.Projection(),
	}

	// Invoco DynamoDB Query API
	result, err := r.service.Scan(params)
	if err != nil {
		return shapes, fmt.Errorf("ListShapesByType. Error al invocar API Query. %s", err)
	}

	// Recorro los items obtenidos
	for _, i := range result.Items {
		shape := models.Request{}

		err = dynamodbattribute.UnmarshalMap(i, &shape) // Parseo y almaceno en user
		if err != nil {
			return shapes, fmt.Errorf("ListShapesByType. Error al parsear item (%s). %s", i, err)
		}
		shapes = append(shapes, shape)
	}

	return shapes, nil
}

func (r Repository) CreateShape(id string, shapeType string, a float64, b float64, creator string) error {
	shape := models.Request{Id: id, ShapeType: shapeType, A: a, B: b, Creator: creator}

	// Parseo cada ítems de Go Types a DynamoDB attributes values
	sh, err := dynamodbattribute.MarshalMap(shape)
	if err != nil {
		return fmt.Errorf("Create. Error de parseo. %s", err)
	}

	// Creo el ítem a insertar en la tabla Users
	input := &dynamodb.PutItemInput{
		Item:      sh,
		TableName: aws.String(constants.SHAPE_TABLE_NAME),
	}

	// Inserto
	_, err = r.service.PutItem(input)
	if err != nil {
		return fmt.Errorf("Create. Error insertando figura (%s). %s", id, err)
	}

	return nil
}
