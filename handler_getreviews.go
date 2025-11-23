package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/adamjames870/tripstats/internal/database"
	"github.com/adamjames870/tripstats/internal/tripapi"
	"github.com/google/uuid"
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
		return
	}

	reviewsInDb, errReviewsInDb := getReviewCountInDatabase(r.Context(), *s.db, params.LocationId)
	if errReviewsInDb != nil {
		reviewsInDb = 0
	}

	resultsToFetch := reviewsForLocation - reviewsInDb

	if resultsToFetch <= 0 {
		respondWithJSON(w, 200, "all reviews up to date")
		return
	}

	numPages := resultsToFetch / 5
	var savedReviews []database.Review

	for i := 0; i < numPages; i++ {

		reviews, errReviews := getReviewsFromTripAdvisor(auth, params.LocationId, 5, (i * 5))

		if errReviews != nil {
			respondWithError(w, 400, "unable to get reviews from tripadvisor: "+errReviews.Error())
		}

		for _, review := range reviews.Data {
			rv, errRv := writeReviewToDb(r.Context(), *s.db, review)
			if errRv == nil {
				// ignore erroneous reviews - they will be duplicates
				savedReviews = append(savedReviews, rv)
			}
		}
	}

	if numPages%5 != 0 {
		reviews, errReviews := getReviewsFromTripAdvisor(auth, params.LocationId, numPages%5, numPages*5)

		if errReviews != nil {
			respondWithError(w, 400, "unable to get reviews from tripadvisor: "+errReviews.Error())
		}

		for _, review := range reviews.Data {
			rv, errRv := writeReviewToDb(r.Context(), *s.db, review)
			if errRv == nil {
				// ignore erroneous reviews - they will be duplicates
				savedReviews = append(savedReviews, rv)
			}
		}
	}

	respondWithJSON(w, 200, savedReviews)

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

func getReviewsFromTripAdvisor(auth tripapi.AuthData, locId string, perPage int, offset int) (tripapi.ReviewCollection, error) {
	params := tripapi.ReviewRequest{
		LocationID: locId,
	}

	return tripapi.GetReviews(auth, params, perPage, offset)
}

func writeReviewToDb(ctx context.Context, db database.Queries, review tripapi.ReviewDetails) (database.Review, error) {
	pubDate, _ := time.Parse(time.RFC3339, review.PublishedDate)
	params := database.SaveReviewParams{
		ID:                  uuid.New(),
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
		TripadvisorReviewID: int32(review.ReviewID),
		LocationID:          strconv.Itoa(int(review.LocationID)),
		PublishedDate:       pubDate,
		TripadvisorUrl: sql.NullString{
			String: review.URL,
			Valid:  true,
		},
		TripadvisorTitle: sql.NullString{
			String: review.Title,
			Valid:  true,
		},
		TripadvisorText: sql.NullString{
			String: review.Text,
			Valid:  true,
		},
		Rating: 0,
	}
	rv, err := db.SaveReview(ctx, params)
	if err != nil {
		return database.Review{}, err
	}
	return rv, nil
}
