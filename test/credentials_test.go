package test

import (
	"go-medium-shapes/pkg/credentials"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCredentials(t *testing.T) {
	t.Run("GetClientDynamo", func(t *testing.T) {
		service, err := credentials.GetClientDynamo()

		assert.NoError(t, err)
		assert.Equal(t, service.ClientInfo.ServiceName, "dynamodb")
	})
}
