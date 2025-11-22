package main

import (
	"fmt"
)

func (state *apiState) CreateEndpoints() {

	// ------------------- Boilerplate -------------------

	state.mux.HandleFunc("GET /api/healthz", handlerHealthz)
	fmt.Println("Added Handlers")
}
