package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/adamjames870/tripstats/internal/database"
	"github.com/adamjames870/tripstats/internal/tripapi"
)

type paramsLocationRequest struct {
	LocationId string `json:"location_id"`
}

// handlerGetLocationInfo godoc
// @Summary      Retrieve location information
// @Description  Fetches location details from the TripAdvisor API using the configured API key.
// @Tags         locations
// @Produce      json
// @Success      200  {object}  tripapi.LocationDetails
// @Failure      400  "Error loading location information"
// @Router       /api/location/info [get]
func (s *apiState) handlerUpdateLocationInfo(w http.ResponseWriter, r *http.Request) {

	auth := tripapi.AuthData{
		Api_key: s.tripApiKey,
	}

	decoder := json.NewDecoder(r.Body)
	params := paramsLocationRequest{}
	errDecode := decoder.Decode(&params)
	if errDecode != nil {
		respondWithError(w, 400, "unable to decode json: "+errDecode.Error())
		return
	}

	data, err := GetLocationInfo(auth, params.LocationId)
	if err != nil {
		respondWithError(w, 400, "error loading data"+err.Error())
		return
	}

	var savedLocation database.LocationInfo
	var errSaveLocation error

	// check if record is in database
	locData, errLocData := s.db.GetLocationFromId(r.Context(), params.LocationId)
	if errLocData != nil {
		// not in database
		dbParams := database.SaveLocationInfoParams{
			ID:         data.LocationID,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			Name:       data.Name,
			WebUrl:     locData.WebUrl,
			Rating:     locData.Rating,
			NumReviews: locData.NumReviews,
		}
		savedLocation, errSaveLocation = s.db.SaveLocationInfo(r.Context(), dbParams)
	} else {
		// only update rating and num_reviews
		dbParams := database.UpdateLocationInfoParams{
			ID:         locData.ID,
			Rating:     locData.Rating,
			NumReviews: locData.NumReviews,
			UpdatedAt:  time.Now(),
		}
		savedLocation, errSaveLocation = s.db.UpdateLocationInfo(r.Context(), dbParams)
	}
	if errSaveLocation != nil {
		respondWithError(w, 400, "unable writing to database"+errSaveLocation.Error())
		return
	}

	respondWithJSON(w, 201, savedLocation)

}

func GetLocationInfo(auth tripapi.AuthData, locId string) (tripapi.LocationDetails, error) {
	params := tripapi.LocationRequest{
		LocationID: locId,
	}
	return tripapi.GetLocationInfo(auth, params)
}
