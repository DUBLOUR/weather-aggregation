package geocode

import (
	"genesis_se/se-school-hw2-DUBLOUR/pkg/generalApiReader"
	"net/http"
	"net/url"
)

type Location struct {
	Lat string
	Lng string
}

func createRequest(city string) (*http.Request, error) {
	//docs: https://geocode.xyz/api
	baseURL, err := url.Parse(ApiEndpoint)
	if err != nil {
		return &http.Request{}, err
	}
	params := url.Values{}
	params.Add("locate", city)
	params.Add("json", "1")
	params.Add("auth", ApiKey)
	baseURL.RawQuery = params.Encode()

	r, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		return &http.Request{}, err
	}
	return r, nil
}

func GetCityLocation(city string) (Location, error) {
	req, err := createRequest(city)
	if err != nil {
		return Location{}, nil
	}

	response := new(struct {
		Lat string `json:"latt"`
		Lng string `json:"longt"`
	})
	if err := generalApiReader.JsonRequest(req, &response); err != nil {
		return Location{}, err
	}

	return Location{
		Lat: response.Lat,
		Lng: response.Lng,
	}, nil
}
