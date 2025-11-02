package main

import (
	"learn-go/crud-mongo-serverless/internal/app"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(app.HttpHandlerRequest)
}
