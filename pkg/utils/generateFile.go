package utils

import (
	"fmt"
	"go-medium-shapes/pkg/constants"
	"go-medium-shapes/pkg/models"
	"go-medium-shapes/pkg/storage"
	"os"

	"github.com/rs/zerolog/log"
)

func GenerateTempFile(shapes []models.IShape) (string, error) {

	file, err := storage.CreateFile()
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}
	defer file.Close()

	for _, sh := range shapes {
		err = storage.WriteInFile(file, sh.Detail())
		if err != nil {
			return "", fmt.Errorf("%s", err)
		}
	}

	return file.Name(), nil
}

func UploadTempFile(path string, fileName string) error {
	log.Info().Msg("(UploadTempFile) Started.")
	file, err := storage.OpenFile(path)
	if err != nil {
		return fmt.Errorf("OpenFile. No fue posible abrir el archivo. %s", err)
	}
	defer os.Remove(file.Name())
	defer file.Close()

	bucketname := constants.BUCKET_NAME
	key := constants.S3_DIRECTORY + "/" + fileName

	s3Client := storage.NewS3Client()
	err = s3Client.UploadFile(&bucketname, &key, file)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
}
