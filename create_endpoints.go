package main

func (state *apiState) CreateEndpoints() {

	// ------------------- Boilerplate -------------------

	state.mux.HandleFunc("GET /api/healthz", handlerHealthz)
}
