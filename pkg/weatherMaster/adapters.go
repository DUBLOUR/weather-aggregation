package weatherMaster

import (
	"github.com/DUBLOUR/weather-aggregation/pkg/openweathermap"
	"github.com/DUBLOUR/weather-aggregation/pkg/stormglass"
	"github.com/DUBLOUR/weather-aggregation/pkg/weatherapi"
)

type OpenweathermapAdapter struct {
	WR openweathermap.WeatherReport
}

func (a OpenweathermapAdapter) InCity(city string) (Weather, error) {
	rawWeather, err := a.WR.InCity(city)
	if err != nil {
		return Weather{}, err
	}
	return Weather{
		Temp: rawWeather.Temp,
		Hum:  rawWeather.Hum,
		Wind: rawWeather.Wind,
	}, nil
}

func (a OpenweathermapAdapter) Name() string {
	return "OpenWeatherMap"
}

type StormglassAdapter struct {
	WR stormglass.WeatherReport
}

func (a StormglassAdapter) InCity(city string) (Weather, error) {
	rawWeather, err := a.WR.InCity(city)
	if err != nil {
		return Weather{}, err
	}
	return Weather{
		Temp: rawWeather.Temp,
		Hum:  int(rawWeather.Hum),
		Wind: rawWeather.Wind,
	}, nil
}
func (a StormglassAdapter) Name() string {
	return "StormGlass"
}

type WeatherapiAdapter struct {
	WR weatherapi.WeatherReport
}

func (a WeatherapiAdapter) InCity(city string) (Weather, error) {
	rawWeather, err := a.WR.InCity(city)
	if err != nil {
		return Weather{}, err
	}
	return Weather{
		Temp: rawWeather.Temp,
		Hum:  rawWeather.Hum,
		Wind: rawWeather.Wind,
	}, nil
}

func (a WeatherapiAdapter) Name() string {
	return "WeatherApi"
}
