package app

import (
	"context"
	"fmt"
	"log"
	"github.com/aws/aws-lambda-go/events"
)

type Event struct {
	Name string `json:"name"`
}


func HandleRequest(ctx context.Context, event Event) (string, error) {
	log.Printf("handle the request runn")

	// event not show up if using apirest invoke method
	log.Printf(fmt.Sprintf("Helooo %s", event.Name))

	return fmt.Sprintf("Helooo %s", event.Name), nil
}

func HttpHandlerRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Print the request body to the console for debugging.
	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	fmt.Printf("Body size = %d.\n", len(request.Body))

	return events.APIGatewayProxyResponse{
		Body: "heloo world from api",
		StatusCode: 200,
	}, nil
}
