package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

type Event struct {
	Name string `json:"name"`
}

// Response is a custom struct to hold our JSON response data.
type Response struct {
	Message string `json:"message"`
}


func HandleRequest(ctx context.Context, event Event) (string, error) {
	log.Printf("handle the request runn")

	// event not show up if using apirest invoke method
	log.Printf(fmt.Sprintf("Helooo %s", event.Name))

	return fmt.Sprintf("Helooo %s", event.Name), nil
}

func HttpHandlerRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Print the request body to the console for debugging.
	// Log the incoming request for debugging
		log.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
		log.Printf("Request Path: %s\n", request.Path)

		// Create a response object
		message := Response{
			Message: fmt.Sprintf("Hello, World from Go! 🚀 You hit the path: %s", request.Path),
		}

		// Marshal the Go struct into a JSON string
		responseBody, err := json.Marshal(message)
		if err != nil {
			return events.APIGatewayProxyResponse{Body: "Error marshalling JSON", StatusCode: 500}, err
		}

		// Return a successful response
		return events.APIGatewayProxyResponse{
			Body:       string(responseBody),
			StatusCode: 200,
			Headers:    map[string]string{"Content-Type": "application/json"},
		}, nil
}
