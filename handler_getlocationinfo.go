package main

import (
	"net/http"

	"github.com/adamjames870/tripstats/internal/tripapi"
)

func (s *apiState) handlerGetLocationInfo(w http.ResponseWriter, r *http.Request) {

	auth := tripapi.AuthData{
		Api_key: s.tripApiKey,
	}

	params := tripapi.LocationRequest{
		LocationID: "25443143",
	}

	data, err := tripapi.GetLocationInfo(auth, params)
	if err != nil {
		respondWithError(w, 400, "error loading"+err.Error())
		return
	}

	respondWithJSON(w, 200, data)

}
