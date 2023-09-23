package storage

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Client struct{}

func NewS3Client() S3Client {
	return S3Client{}
}

func (s3c S3Client) UploadFile(bucketname *string, fileKey *string, file *os.File) error {
	session, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	service := s3manager.NewUploader(session)

	_, err := service.Upload(&s3manager.UploadInput{
		Bucket: bucketname,
		Key:    fileKey,
		Body:   file,
	})
	if err != nil {
		return fmt.Errorf("UploadFile. No fue posible subir el archivo (%s). %s", *fileKey, err)
	}

	return nil
}
