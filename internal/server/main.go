package server

import (
	"genesis_se/se-school-hw2-DUBLOUR/pkg/openweathermap"
	"genesis_se/se-school-hw2-DUBLOUR/pkg/stormglass"
	"genesis_se/se-school-hw2-DUBLOUR/pkg/weatherMaster"
	"genesis_se/se-school-hw2-DUBLOUR/pkg/weatherapi"
)

func CreateMaster() weatherMaster.Master {
	m := weatherMaster.Master{}
	m.InitSource(
		weatherMaster.OpenweathermapAdapter{openweathermap.WeatherReport{}},
		weatherMaster.WeatherapiAdapter{weatherapi.WeatherReport{}},
		weatherMaster.StormglassAdapter{stormglass.WeatherReport{}},
	)
	m.SetLogger(new(Logger))
	m.SetMetricHandler(Metric{"data/city.json"})
	return m
}

func HandleCity(city string) (weatherMaster.Weather, error) {
	m := CreateMaster()
	return m.WeatherInCity(city)
}
