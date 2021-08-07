package main

import (
	"encoding/json"
	"fmt"
	"genesis_se/se-school-hw2-DUBLOUR/pkg/weatherapi"
)

//type Weather string

type Weather struct {
	Temp float32
	Hum int
	Wind float32
}

func main()  {
	city := "Kyiv"
	city = "Kiev"
	//w, _ := openweathermap.GetWeather(city)
	w, _ := weatherapi.GetWeather(city)
	j, _ := json.Marshal(w)
	fmt.Println(string(j))
	//fmt.Println(w.temp, ' ', w.hum, ' ', w.wind)
}
