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
	Rating     string `json:"rating"`
	NumReviews string `json:"num_reviews"`
}

type ReviewRequest struct {
	LocationID string
}

type ReviewDetails struct {
	ReviewID      int64  `json:"id"`
	LocationID    int64  `json:"location_id"`
	PublishedDate string `json:"published_date"`
	URL           string `json:"url"`
	Title         string `json:"title"`
	Text          string `json:"text"`
	Rating        int64  `json:"rating"`
}
