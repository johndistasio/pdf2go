package main

import (
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jung-kurt/gofpdf"
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
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
		}, ErrNameNotProvided
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello "+request.Body)

	var buf bytes.Buffer
	pdf.Output(&buf)

	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())

	return events.APIGatewayProxyResponse{
		Body:            encoded,
		StatusCode:      201,
		IsBase64Encoded: true,
		Headers: map[string]string{
			"Content-Type": "application/pdf",
		},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
