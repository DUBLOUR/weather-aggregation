package geocode

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Location struct {
	Lat string
	Lng string
}

func GetCityLocation(city string) (Location, error) {
	//docs: https://geocode.xyz/api
	baseURL, err := url.Parse(ApiEndpoint)
	if err != nil {
		return Location{}, nil
	}
	params := url.Values{}
	params.Add("locate", city)
	params.Add("json", "1")
	params.Add("auth", ApiKey)

	baseURL.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		return Location{}, nil
	}

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Location{}, nil
	}

	type Response struct {
		Lat string `json:"latt"`
		Lng string `json:"longt"`
	}

	response := new(Response)
	if err := json.Unmarshal(body, &response); err != nil {
		return Location{}, err
	}

	return Location{
		Lat: response.Lat,
		Lng: response.Lng,
	}, nil

}
