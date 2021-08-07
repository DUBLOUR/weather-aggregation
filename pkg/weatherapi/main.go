package weatherapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Weather struct {
	Temp float32
	Hum int
	Wind float32
}

func GetWeather(city string) (Weather, error) {
	//docs: https://www.weatherapi.com/docs/#apis-realtime
	baseURL, err := url.Parse(ApiEndpoint)
	if err != nil {
		return Weather{}, nil
	}
	params := url.Values{}
	params.Add("q", city)
	params.Add("key", ApiKey)
	params.Add("aqi", "no")

	baseURL.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		return Weather{}, nil
	}

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Weather{}, nil
	}

	type Response struct {
		Core struct{
			Temp float32 `json:"temp_c"`
			Hum int `json:"humidity"`
			Wind float32 `json:"wind_kph"`
		} `json:"current"`
	}

	response := new(Response)
	if err := json.Unmarshal(body, &response); err != nil {
		return Weather{}, err
	}

	return Weather{
		Temp: response.Core.Temp,
		Hum: response.Core.Hum,
		Wind: response.Core.Wind,
	}, nil

}
