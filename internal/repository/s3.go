package repository

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

type S3Repository struct {
	client *s3.Client
}

type IS3Repository interface {
	UploadFile(bucketName string, objectKey string, fileReader io.Reader) error
}

func NewS3Repository() IS3Repository {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("Configuration error, " + err.Error())
	}

	return S3Repository{client: s3.NewFromConfig(cfg)}
}

func (s3r S3Repository) UploadFile(bucketName string, objectKey string, fileReader io.Reader) error {

	_, err := s3r.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   fileReader,
	})
	if err != nil {
		return fmt.Errorf("UploadFile. No fue posible subir el archivo (%s). %s", objectKey, err)
	}

	return nil
}
