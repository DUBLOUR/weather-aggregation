package openweathermap

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
	//docs: https://openweathermap.org/current
	baseURL, err := url.Parse(ApiEndpoint)
	if err != nil {
		return Weather{}, nil
	}
	params := url.Values{}
	params.Add("q", city)
	params.Add("appid", ApiKey)
	params.Add("units", "metric")

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
			Temp float32 `json:"temp"`
			Hum int `json:"humidity"`
		} `json:"main"`
		Wind struct{
			Speed int `json:"speed"`
		} `json:"wind"`
	}

	response := new(Response)
	if err := json.Unmarshal(body, &response); err != nil {
		return Weather{}, err
	}

	return Weather{
		Temp: response.Core.Temp,
		Hum: response.Core.Hum,
		Wind: float32(response.Wind.Speed),
	}, nil

}
