package credentials

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
)

func getSession() (*session.Session, error) {
	// Creo sesión tomando credenciales y región de ~/.aws/credentials y ~/.aws/config
	return session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})), nil

	/*aws_access_key_id := os.Getenv("AWS_ACCESS_KEY_ID")
	aws_secret_access_key := os.Getenv("AWS_SECRET_ACCESS_KEY")
	aws_default_region := os.Getenv("AWS_DEFAULT_REGION")

	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, "")
	_, err := creds.Get()
	if err != nil {
		//fmt.Println(err)
		//os.Exit(1)
		return nil, err
	}
	fmt.Println("SSSS")
	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(aws_default_region),
		Credentials: creds,
	})
	return session, err*/
}

func GetClientS3() (*s3.S3, error) {
	// Instancio cliente de servicio S3
	session, err := getSession()
	if err != nil {
		return nil, err
	}
	return s3.New(session), nil
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
