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
	state.mux.HandleFunc("POST /api/location", state.handlerSaveLocation)
	state.mux.HandleFunc("GET /api/locationinfo", state.handlerUpdateLocationInfo)

	// ------------------- Reviews -------------------

	state.mux.HandleFunc("GET /api/reviewinfo", state.handlerGetReviews)

	fmt.Println("Added Handlers")
}
