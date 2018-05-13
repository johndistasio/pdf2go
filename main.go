package main

import (
	"encoding/base64"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/jung-kurt/gofpdf"
	"log"
)

var (
	// ErrNameNotProvided is thrown when a name is not provided
	ErrNameNotProvided = errors.New("no named was provided in the HTTP body")
)

// Handler processes incoming requests
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr and sent to cloudwatch
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	// If no name is provided in HTTP body, throw an error
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, ErrNameNotProvided
	}

	return events.APIGatewayProxyResponse{
		Body:            base64.StdEncoding.EncodeToString([]byte(request.Body)),
		StatusCode:      201,
		IsBase64Encoded: true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
