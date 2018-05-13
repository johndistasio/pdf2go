package main

import (
	"bytes"
	"encoding/base64"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jung-kurt/gofpdf"
	"log"
)

// Handler processes incoming requests
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr and sent to cloudwatch
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	if len(request.Body) < 1 {
		log.Print("Bad request: no body provided\n")
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
		}, nil
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
