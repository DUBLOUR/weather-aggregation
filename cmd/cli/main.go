package main

import (
	"encoding/json"
	"fmt"
	"genesis_se/se-school-hw2-DUBLOUR/pkg/openweathermap"
)

//type Weather string

type Weather struct {
	Temp float32
	Hum int
	Wind float32
}

func main()  {
	city := "Kyiv"
	w, _ := openweathermap.GetWeather(city)
	j, _ := json.Marshal(w)
	fmt.Println(string(j))
	//fmt.Println(w.temp, ' ', w.hum, ' ', w.wind)
}
