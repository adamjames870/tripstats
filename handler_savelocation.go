package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/adamjames870/tripstats/internal/database"
	"github.com/google/uuid"
)

type paramsCreateLocation struct {
	LocationId string `json:"tripadvisor_id"`
	Name       string `json:"name"`
}

// handlerSaveLocation godoc
// @Summary      Create a new location
// @Description  Accepts a JSON payload containing a location's external ID and name, stores it, and returns the saved record.
// @Tags         locations
// @Accept       json
// @Produce      json
// @Param        request  body     paramsCreateLocation  true  "Location details"
// @Success      200      {object} paramsCreateLocation
// @Failure      400      {object} map[string]string "Invalid request or database error"
// @Router       /api/locations [post]
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
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
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

	respondWithJSON(w, 201, rv)

}
