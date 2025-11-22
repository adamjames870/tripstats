package main

import "net/http"

// HealthCheck godoc
// @Summary      Health check
// @Description  Returns 200 if service is healthy
// @Tags         health
// @Success      200 {object} map[string]interface{}
// @Router       /api/healthz [get]
func handlerHealthz(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, nil)
}
