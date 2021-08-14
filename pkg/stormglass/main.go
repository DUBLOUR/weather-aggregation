package stormglass

import (
	"fmt"
	"genesis_se/se-school-hw2-DUBLOUR/pkg/generalApiReader"
	"genesis_se/se-school-hw2-DUBLOUR/pkg/geocode"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Weather struct {
	Temp float32
	Hum  float32
	Wind float32
}

func createRequest(city string) (*http.Request, error) {
	loc, err := geocode.GetCityLocation(city)
	//fmt.Println(loc.Lng, loc.Lat)
	if err != nil {
		return &http.Request{}, err
	}

	nowTime := strconv.FormatInt(time.Now().UTC().Unix(), 10)

	//docs: https://docs.stormglass.io/#/weather
	baseURL, err := url.Parse(ApiEndpoint)
	if err != nil {
		return &http.Request{}, err
	}
	params := url.Values{}
	params.Add("lat", loc.Lat)
	params.Add("lng", loc.Lng)
	params.Add("start", nowTime)
	params.Add("end", nowTime)
	params.Add("source", StormglassSource)
	params.Add("params", "airTemperature,windSpeed,humidity")
	baseURL.RawQuery = params.Encode()

	r, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		return &http.Request{}, err
	}
	r.Header.Set("Authorization", ApiKey)
	return r, nil
}

type WeatherReport struct{}

func (w WeatherReport) InCity(city string) (Weather, error) {
	req, err := createRequest(city)
	if err != nil {
		return Weather{}, nil
	}

	response := new(struct {
		Core []struct {
			Temp map[string]float32 `json:"airTemperature"`
			Hum  map[string]float32 `json:"humidity"`
			Wind map[string]float32 `json:"windSpeed"`
		} `json:"hours"`
	})

	if err := generalApiReader.JsonRequest(req, &response); err != nil {
		return Weather{}, err
	}

	if _, hasSource := response.Core[0].Temp[StormglassSource]; !hasSource {
		return Weather{}, fmt.Errorf("invalid Stormglass source")
	}

	return Weather{
		Temp: response.Core[0].Temp[StormglassSource],
		Hum:  response.Core[0].Hum[StormglassSource],
		Wind: response.Core[0].Wind[StormglassSource],
	}, nil
}
