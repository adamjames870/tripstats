package tripapi

type AuthData struct {
	Api_key string
}

type LocationRequest struct {
	LocationID string
}

type LocationDetails struct {
	LocationID string `json:"location_id"`
	Name       string `json:"name"`
	WebURL     string `json:"web_url"`
	NumReviews string `json:"num_reviews"`
}
