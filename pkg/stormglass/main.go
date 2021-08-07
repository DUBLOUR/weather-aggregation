package stormglass

import (
	"encoding/json"
	"genesis_se/se-school-hw2-DUBLOUR/pkg/geocode"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Weather struct {
	Temp float32
	Hum int
	Wind float32
}

func GetWeather(city string) (Weather, error) {
	loc, err := geocode.GetCityLocation(city)
	if err != nil {
		return Weather{}, nil
	}

	nowTime := strconv.FormatInt(time.Now().UTC().Unix(), 10)

	//docs: https://docs.stormglass.io/#/weather
	baseURL, err := url.Parse(ApiEndpoint)
	if err != nil {
		return Weather{}, nil
	}
	params := url.Values{}
	params.Add("lat", loc.Lat)
	params.Add("lng", loc.Lng)
	params.Add("start", nowTime)
	params.Add("end", nowTime)
	params.Add("source", "sg")
	params.Add("params", "airTemperature,windSpeed,humidity")
	baseURL.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		return Weather{}, nil
	}
	req.Header.Set(
		"Authorization",
		ApiKey,
	)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Weather{}, nil
	}

	type Response struct {
		Core[] struct{
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
	}

	response := new(Response)
	if err := json.Unmarshal(body, &response); err != nil {
		return Weather{}, err
	}

	return Weather{
		Temp: response.Core[0].Temp.Value,
		Hum: int(response.Core[0].Hum.Value),
		Wind: response.Core[0].Wind.Value,
	}, nil

}
