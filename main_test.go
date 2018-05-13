package main_test

import (
	"github.com/aws/aws-lambda-go/events"
	main "github.com/johndistasio/pdf2go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	request := events.APIGatewayProxyRequest{
		Body: "Dolph Lundgren",
	}

	response, err := main.Handler(request)

	assert.NotNil(t, response.Body)
	assert.Equal(t, 201, response.StatusCode)
	assert.Nil(t, err)
}

func TestHandlerBadRequest(t *testing.T) {
	request := events.APIGatewayProxyRequest{}

	response, err := main.Handler(request)

	assert.Equal(t, 400, response.StatusCode)
	assert.Nil(t, err)
}
