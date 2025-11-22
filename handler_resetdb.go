package main

import "net/http"

// ResetDB godoc
// @Summary      Resets DB
// @Description  Removes all data in DB, Returns 200 if successful
// @Tags         reset
// @Success      200 {object} map[string]interface{}
// @Router       /api/reset [post]
func (s *apiState) handlerResetDb(w http.ResponseWriter, r *http.Request) {

	// POST api/reset

	s.db.ResetLocations(r.Context())
	respondWithJSON(w, 200, nil)
}
