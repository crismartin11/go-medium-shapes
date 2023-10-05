package processor

import (
	"context"
	"fmt"
	"os"
	"time"

	"go-medium-shapes/pkg/constants"
	"go-medium-shapes/pkg/models"
	"go-medium-shapes/pkg/utils"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type Repository interface {
	ListShapesByType(shapeType string) ([]models.Request, error)
	CreateShape(id string, shapeType string, a float64, b float64, creator string) error
}

type RepositoryS3 interface {
	UploadFile(bucketname *string, fileKey *string, file *os.File) error
}

type Processor struct {
	r Repository
}

func New(r Repository) Processor {
	return Processor{
		r,
	}
}

func (p Processor) ProcessCreation(request models.Request) (models.Response, error) {
	log.Info().Msg("Create shape handler started (ProcessCreation).")

	// Get user data from api
	user, err := utils.HttpGetUserData(request.Id)
	if err != nil {
		log.Error().Msg("ProcessCreation.Error getting user data from API.")
		return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
	} else if user.Data.Email == "" {
		log.Error().Msg("ProcessCreation.Error getting email user from data.")
		return models.NewResponseError(400, fmt.Sprintf("ProcessCreation. No se obtuvo el email del usuario con id %s", request.Id))
	}

	// Insert shape in table
	uuid, _ := uuid.NewUUID()
	err = p.r.CreateShape(uuid.String(), request.ShapeType, request.A, request.B, user.Data.Email)
	if err != nil {
		log.Error().Msg("ProcessCreation. Error creating shape in DynamoDB.")
		return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
	}

	log.Info().Msg("Create shape handler finished successfully (ProcessCreation).")
	return models.NewResponseOk("Creation process successful!")
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (p Processor) ProcessGeneration(ctx context.Context, request models.Request) (models.Response, error) {
	log.Info().Msg("ProcessGeneration. Generate shape list file handler started.")

	// Get shapes from table
	log.Info().Msg("ProcessGeneration. Getting list from dynamoDB.")
	listShapes, err := p.r.ListShapesByType(request.ShapeType)
	if err != nil {
		log.Error().Msg("ProcessGeneration. Error getting shape list from DynamoDB.")
		return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
	}

	// Generate Shape list
	log.Info().Msg("ProcessGeneration. Generating shapes.")
	var shapes = []models.IShape{}
	for _, item := range listShapes {
		elem, err := models.ShapeFactory(item.Id, item.ShapeType, item.A, item.B)
		if err != nil {
			log.Error().Msg("ProcessGeneration. Error making shape based on dynamo item.")
			return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
		}
		shapes = append(shapes, elem)
	}

	// Generate file
	log.Info().Msg("ProcessGeneration. Generating file.")
	path, err := utils.GenerateTempFile(shapes)
	if err != nil {
		log.Error().Msg("ProcessGeneration. Error generating file.")
		return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
	} else {
		// Upload file
		log.Info().Msg("ProcessGeneration. Uploading file " + path + ".")
		err = utils.UploadTempFile(path, getFileName(ctx, request))
		if err != nil {
			log.Error().Msg("ProcessGeneration. Error uploading file.")
			return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
		}
	}

	log.Info().Msg("ProcessGeneration. Generate shape list file handler finished successfully.")
	return models.NewResponseOk("Generation file process successful!")
}

func getFileName(ctx context.Context, request models.Request) string {
	lc, ok := lambdacontext.FromContext(ctx)
	awsRequestID := "unknown"
	if ok {
		awsRequestID = lc.AwsRequestID
	}
	return request.ShapeType + "-" + awsRequestID + "-" + time.Now().Format(constants.DATE_FORMAT) + ".txt"
}
