package openweathermap

import (
	"genesis_se/se-school-hw2-DUBLOUR/pkg/generalApiReader"
	"net/http"
	"net/url"
)

type Weather struct {
	Temp float32
	Hum  int
	Wind float32
}

func createRequest(city string) (*http.Request, error) {
	//docs: https://openweathermap.org/current
	baseURL, err := url.Parse(ApiEndpoint)
	if err != nil {
		return &http.Request{}, err
	}
	params := url.Values{}
	params.Add("q", city)
	params.Add("appid", ApiKey)
	params.Add("units", "metric")
	baseURL.RawQuery = params.Encode()

	r, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		return &http.Request{}, err
	}
	return r, nil
}

type WeatherReport struct{}

func (w WeatherReport) InCity(city string) (Weather, error) {
	req, err := createRequest(city)
	if err != nil {
		return Weather{}, nil
	}

	response := new(struct {
		Core struct {
			Temp float32 `json:"temp"`
			Hum  int     `json:"humidity"`
		} `json:"main"`
		Wind struct {
			Speed float32 `json:"speed"`
		} `json:"wind"`
	})

	if err := generalApiReader.JsonRequest(req, &response); err != nil {
		return Weather{}, err
	}

	return Weather{
		Temp: response.Core.Temp,
		Hum:  response.Core.Hum,
		Wind: response.Wind.Speed,
	}, nil
}
