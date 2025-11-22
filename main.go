package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	var state apiState
	errLoadSTate := state.LoadState()
	if errLoadSTate != nil {
		fmt.Println("Error loading state", errLoadSTate)
		os.Exit(1)
	}

	fmt.Println("Loaded state")

	state.mux = http.NewServeMux()
	state.CreateEndpoints()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := http.Server{
		Handler: state.mux,
		Addr:    ":8080", // + port,
	}

	errServe := server.ListenAndServe()
	if errServe != nil {
		fmt.Println("Error loading state", errServe)
		os.Exit(1)
	}

}
