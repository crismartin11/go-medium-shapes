package test

import (
	"go-medium-shapes/pkg/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFile(t *testing.T) {
	file, err := storage.CreateFile()

	assert.NoError(t, err)
	assert.NotEmpty(t, file, "No debería estar vacío.")
	defer file.Close()
}
