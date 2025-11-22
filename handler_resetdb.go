package main

import "net/http"

func (s *apiState) handlerResetDb(w http.ResponseWriter, r *http.Request) {

	// POST api/reset

	s.db.ResetLocations(r.Context())
	respondWithJSON(w, 200, nil)
}
