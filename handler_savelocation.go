package main

import (
	"encoding/json"
	"net/http"

	"github.com/adamjames870/tripstats/internal/database"
	"github.com/google/uuid"
)

type paramsCreateLocation struct {
	LocationId string `json:"tripadvisor_id"`
	Name       string `json:"name"`
}

func (s *apiState) handlerSaveLocation(w http.ResponseWriter, r *http.Request) {

	// POST api/locations

	decoder := json.NewDecoder(r.Body)
	params := paramsCreateLocation{}
	errDecode := decoder.Decode(&params)
	if errDecode != nil {
		respondWithError(w, 400, "unable to decode json: "+errDecode.Error())
		return
	}

	dbParams := database.SaveLocationParams{
		ID:         uuid.New(),
		LocationID: params.LocationId,
		Name:       params.Name,
	}

	loc, errLoc := s.db.SaveLocation(r.Context(), dbParams)
	if errLoc != nil {
		respondWithError(w, 400, "Unable to write to database:"+errLoc.Error())
		return
	}

	rv := paramsCreateLocation{
		LocationId: loc.LocationID,
		Name:       loc.Name,
	}

	respondWithJSON(w, 200, rv)

}
