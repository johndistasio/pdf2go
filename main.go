package main

import (
	"fmt"
	_ "github.com/aws/aws-lambda-go/events"
	_ "github.com/aws/aws-lambda-go/lambda"
	_ "github.com/jung-kurt/gofpdf"
)

func main() {
	fmt.Println("pdf2go")
}
