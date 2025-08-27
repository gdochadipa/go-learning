package main

import (
	"fmt"
	"learn-go/crud-mongo-serverless/internal/app"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var client *mongo.Client

func main() {

	log.Printf("runing handler")

	lambda.Start(app.HandleRequest)
	// mongoClient, db, err := db.Connect(client, "", false)
	fmt.Println("test go")
}
