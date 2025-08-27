package app

import (
	"context"
	"fmt"
	"log"
)

type Event struct {
	Name string `json:"name"`
}


func HandleRequest(ctx context.Context, event Event) (string, error) {
	log.Printf("handle the request runn")
	log.Printf(fmt.Sprintf("Helooo %s", event.Name))

	return fmt.Sprintf("Helooo %s", event.Name), nil
}
