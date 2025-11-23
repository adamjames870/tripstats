package tripapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetLocationInfo(auth AuthData, params LocationRequest) (LocationDetails, error) {
	// "https://api.content.tripadvisor.com/api/v1/location"

	nullReturn := LocationDetails{}

	base := "https://api.content.tripadvisor.com/api/v1/location"
	endpoint := fmt.Sprintf("%s/%s/details?key=%s&language=en&currency=USD",
		base, params.LocationID, auth.Api_key)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nullReturn, err
	}

	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nullReturn, err
	}

	var details LocationDetails

	if err := json.Unmarshal(body, &details); err != nil {
		return nullReturn, err
	}

	return details, nil

}
