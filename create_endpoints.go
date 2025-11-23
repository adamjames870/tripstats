package main

import (
	"fmt"

	httpSwagger "github.com/swaggo/http-swagger"
)

func (state *apiState) CreateEndpoints() {

	// ------------------- Boilerplate -------------------

	state.mux.HandleFunc("GET /api/healthz", handlerHealthz)
	state.mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// ------------------- Locations -------------------

	state.mux.HandleFunc("POST /api/reset", state.handlerResetDb)
	state.mux.HandleFunc("POST /api/locations", state.handlerSaveLocation)
	state.mux.HandleFunc("GET /api/locations/{LocationId}", state.handlerUpdateLocationInfo)

	fmt.Println("Added Handlers")
}
