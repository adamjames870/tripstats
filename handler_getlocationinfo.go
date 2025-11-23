package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
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
// @Router       /api/locationinfo [get]
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
		respondWithError(w, 400, "error loading data: "+err.Error())
		return
	}

	var savedLocation database.LocationInfo
	var errSaveLocation error

	// check if record is in database
	_, errLocData := s.db.GetLocationFromId(r.Context(), params.LocationId)
	if errLocData != nil {
		// not in database
		rating, _ := strconv.ParseFloat(data.Rating, 64)
		numReviews, _ := strconv.Atoi(data.NumReviews)
		dbParams := database.SaveLocationInfoParams{
			ID:        data.LocationID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      data.Name,
			WebUrl: sql.NullString{
				String: data.WebURL,
				Valid:  true,
			},
			Rating: sql.NullFloat64{
				Float64: rating,
				Valid:   true,
			},
			NumReviews: sql.NullInt32{
				Int32: int32(numReviews),
				Valid: true,
			},
		}
		savedLocation, errSaveLocation = s.db.SaveLocationInfo(r.Context(), dbParams)
	} else {
		// only update rating and num_reviews
		rating, _ := strconv.ParseFloat(data.Rating, 64)
		numReviews, _ := strconv.Atoi(data.NumReviews)
		dbParams := database.UpdateLocationInfoParams{
			ID: data.LocationID,
			Rating: sql.NullFloat64{
				Float64: rating,
				Valid:   true,
			},
			NumReviews: sql.NullInt32{
				Int32: int32(numReviews),
				Valid: true,
			},
			UpdatedAt: time.Now(),
		}
		savedLocation, errSaveLocation = s.db.UpdateLocationInfo(r.Context(), dbParams)
	}
	if errSaveLocation != nil {
		respondWithError(w, 400, "unable to write to database: "+errSaveLocation.Error())
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
