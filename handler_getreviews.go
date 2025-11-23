package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/adamjames870/tripstats/internal/database"
	"github.com/adamjames870/tripstats/internal/tripapi"
)

type paramsReviewRequest struct {
	LocationId string `json:"location_id"`
}

func (s *apiState) handlerGetReviews(w http.ResponseWriter, r *http.Request) {

	auth := tripapi.AuthData{
		Api_key: s.tripApiKey,
	}

	decoder := json.NewDecoder(r.Body)
	params := paramsReviewRequest{}
	errDecode := decoder.Decode(&params)
	if errDecode != nil {
		respondWithError(w, 400, "unable to decode json: "+errDecode.Error())
		return
	}

	reviewsForLocation, errReviewsForLoc := getReviewNumsInLocation(r.Context(), *s.db, params.LocationId)
	if errReviewsForLoc != nil {
		respondWithError(w, 400, "unable to read location info in db: "+errReviewsForLoc.Error())
	}

	reviewsInDb, errReviewsInDb := getReviewCountInDatabase(r.Context(), *s.db, params.LocationId)
	if errReviewsInDb != nil {
		reviewsInDb = 0
	}

	if reviewsForLocation <= reviewsInDb {
		respondWithJSON(w, 200, "all reviews up to date")
	}

	// TODO edit to get most recent x reviews

	reviews, errReviews := getReviewsFromTripAdvisor(auth, params.LocationId)
	if errReviews != nil {
		respondWithError(w, 400, "unable to get reviews from tripadvisor: "+errReviewsForLoc.Error())
	}

	respondWithJSON(w, 200, reviews)

}

func getReviewNumsInLocation(ctx context.Context, db database.Queries, loc_id string) (int, error) {
	data, err := db.GetLocationFromId(ctx, loc_id)
	if err != nil {
		return 0, err
	}
	if data.NumReviews.Valid {
		return int(data.NumReviews.Int32), nil // Convert Int32 to int
	} else {
		return 0, nil
	}
}

func getReviewCountInDatabase(ctx context.Context, db database.Queries, loc_id string) (int, error) {
	reviewCount, err := db.GetReviewCount(ctx, loc_id)
	if err != nil {
		return 0, err
	}
	return int(reviewCount), nil
}

func getReviewsFromTripAdvisor(auth tripapi.AuthData, locId string) (tripapi.ReviewDetails, error) {
	params := tripapi.ReviewRequest{
		LocationID: locId,
	}
	return tripapi.GetReviews(auth, params)
}
