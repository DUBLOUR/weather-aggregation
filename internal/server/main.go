package server

import (
	"genesis_se/se-school-hw2-DUBLOUR/pkg/openweathermap"
	"genesis_se/se-school-hw2-DUBLOUR/pkg/stormglass"
	"genesis_se/se-school-hw2-DUBLOUR/pkg/weatherMaster"
	"genesis_se/se-school-hw2-DUBLOUR/pkg/weatherapi"
	"log"
)

func CreateMaster() weatherMaster.Master {
	m := weatherMaster.Master{}
	m.InitSource(
		weatherMaster.WeatherapiAdapter{weatherapi.WeatherReport{}},
		weatherMaster.StormglassAdapter{stormglass.WeatherReport{}},
		weatherMaster.OpenweathermapAdapter{openweathermap.WeatherReport{}},
	)

	m.SetLogger(Logger{LogFile})
	log.Println("Set `" + LogFile + "` as LogFile")
	m.SetMetricHandler(Metric{MetricDbFile})
	log.Println("Set `" + MetricDbFile + "` as MetricDbFile")
	return m
}

func HandleCity(city string) (weatherMaster.Weather, error) {
	m := CreateMaster()
	return m.WeatherInCity(city)
}
