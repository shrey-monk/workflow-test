package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler for the main route
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, welcome to my Go service!")
}

// Health check handler
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Health check request received\n")
	// Respond with 200 OK to indicate that the service is healthy
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Healthy")
}

func main() {
	// Register the routes
	http.HandleFunc("/my-app", handler)
	http.HandleFunc("/health", healthCheckHandler) // Add health check route

	port := "8080"
	fmt.Printf("Starting server on port %s...\n", port)

	// Start the server
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
