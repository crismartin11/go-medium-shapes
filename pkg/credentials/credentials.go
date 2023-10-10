package credentials

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func getSession() (*session.Session, error) {
	// Creo sesión tomando credenciales y región de ~/.aws/credentials y ~/.aws/config
	return session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})), nil
}

func GetClientS3() (*s3.S3, error) {
	// Instancio cliente de servicio S3
	session, err := getSession()
	if err != nil {
		return nil, err
	}
	return s3.New(session), nil
}

func GetClientS3Uploader() (*s3manager.Uploader, error) {
	// Instancio cliente de servicio s3manager
	session, err := getSession()
	if err != nil {
		return nil, err
	}
	return s3manager.NewUploader(session), nil
}

func GetClientLambda() (*lambda.Lambda, error) {
	// Create Lambda service client
	session, err := getSession()
	if err != nil {
		return nil, err
	}
	return lambda.New(session), nil
}

func GetClientDynamo() (*dynamodb.DynamoDB, error) {
	// Creo cliente de DynamoDB
	session, err := getSession()
	if err != nil {
		return nil, err
	}
	return dynamodb.New(session), nil
}
