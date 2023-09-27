package storage

import (
	"fmt"
	"go-medium-shapes/pkg/credentials"
	"os"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Client struct{}

func NewS3Client() S3Client {
	return S3Client{}
}

func (s3c S3Client) UploadFile(bucketname *string, fileKey *string, file *os.File) error {
	service, err := credentials.GetClientS3Uploader()
	if err != nil {
		return fmt.Errorf("UploadFile. Error obteniendo credenciales de AWS. %s", err)
	}

	_, err = service.Upload(&s3manager.UploadInput{
		Bucket: bucketname,
		Key:    fileKey,
		Body:   file,
	})
	if err != nil {
		return fmt.Errorf("UploadFile. No fue posible subir el archivo (%s). %s", *fileKey, err)
	}

	return nil
}
