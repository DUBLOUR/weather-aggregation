package main

import (
	"encoding/json"
	"fmt"
	"genesis_se/se-school-hw2-DUBLOUR/pkg/openweathermap"
	"genesis_se/se-school-hw2-DUBLOUR/pkg/stormglass"
	"genesis_se/se-school-hw2-DUBLOUR/pkg/weatherapi"
)

//type Weather string

type Weather struct {
	Temp float32
	Hum  int
	Wind float32
}

func main() {
	city := "Kyiv"
	city = "Kiev"
	w1, err := openweathermap.GetWeather(city)
	w2, err := weatherapi.GetWeather(city)
	w3, err := stormglass.GetWeather(city)

	j, _ := json.Marshal([]Weather{Weather(w1), Weather(w2), Weather(w3)})
	fmt.Println(string(j))
	fmt.Println(err)
	//fmt.Println(w.temp, ' ', w.hum, ' ', w.wind)
}
