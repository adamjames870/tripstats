package main

import (
	"fmt"
)

func (state *apiState) CreateEndpoints() {

	// ------------------- Boilerplate -------------------

	state.mux.HandleFunc("GET /api/healthz", handlerHealthz)

	// ------------------- Locations -------------------

	state.mux.HandleFunc("POST /api/reset", state.handlerResetDb)
	state.mux.HandleFunc("POST /api/locations", state.handlerSaveLocation)

	fmt.Println("Added Handlers")
}
