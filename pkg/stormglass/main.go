package stormglass

import (
	"genesis_se/se-school-hw2-DUBLOUR/pkg/generalApiReader"
	"genesis_se/se-school-hw2-DUBLOUR/pkg/geocode"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Weather struct {
	Temp float32
	Hum  int
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
	params.Add("source", "sg")
	params.Add("params", "airTemperature,windSpeed,humidity")
	baseURL.RawQuery = params.Encode()

	r, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		return &http.Request{}, err
	}
	r.Header.Set("Authorization", ApiKey)
	return r, nil
}

func GetWeather(city string) (Weather, error) {
	req, err := createRequest(city)
	if err != nil {
		return Weather{}, nil
	}

	response := new(struct {
		Core []struct {
			Temp struct {
				Value float32 `json:"sg"`
			} `json:"airTemperature"`
			Hum struct {
				Value float32 `json:"sg"`
			} `json:"humidity"`
			Wind struct {
				Value float32 `json:"sg"`
			} `json:"windSpeed"`
		} `json:"hours"`
	})

	if err := generalApiReader.JsonRequest(req, &response); err != nil {
		return Weather{}, err
	}

	return Weather{
		Temp: response.Core[0].Temp.Value,
		Hum:  int(response.Core[0].Hum.Value),
		Wind: response.Core[0].Wind.Value,
	}, nil
}
