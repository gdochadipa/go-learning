package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// create a server
func createServer() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		log.Println("slow request started...")
		time.Sleep(5 * time.Second)
		fmt.Fprintf(w, "slow Request completed at %v\n", time.Now())
	})

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}

// running the server
func runServer(
	ctx context.Context,
	server *http.Server,
	shutdownTimeout time.Duration,
) error {
	// start listen and serve
	serverErr := make(chan error, 1)
	go func() {
		log.Println("Server running on :8080....")
		if err := server.ListenAndServe(); err != nil {
			serverErr <- err
		}
		// permited to receive any value after close channel,
		// but the channel still contain the value
		close(serverErr)
	}()

	// signal notify
	// this for notify the os, the app are going to interupted
	stop := make(chan os.Signal, 1)

	// this code is for we receive any signal interupt or terminate from os
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	// waits and receive from channel
	case err := <-serverErr:
		// for case error from server, we will not terminate the server
		return err
	// wait value from "stop" channel
	case <-stop:
		// receive signal from os, than we shut down it
		log.Println("shutdown signal received")
	// waiting if the context already done, or there is terminate from outside function
	// done is use channel
	case <-ctx.Done():
		log.Print("context canceled")
	}

	// this will start shuting down properly
	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)

	defer cancel()

	// for the start let shutdown the server
	// Shutdown gracefully shuts down the server without interrupting any active connections
	if err := server.Shutdown(shutdownCtx); err != nil {
		// server close is like terminate ly server
		if closeErr := server.Close(); closeErr != nil {
			return errors.Join(err, closeErr)
		}
	}

	log.Println("server exited gracefully")

	return nil
}

func main() {
	server := createServer()

	if err := runServer(context.Background(), server, 3*time.Second); err != nil {
		log.Fatalf("Server err: %v", err)
	}
}
