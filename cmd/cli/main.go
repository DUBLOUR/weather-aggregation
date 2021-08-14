package main

import (
	"encoding/json"
	"fmt"
	"genesis_se/se-school-hw2-DUBLOUR/internal/server"
	"os"
)

func main() {
	fmt.Println(os.Args)
	if !(len(os.Args) == 2) {
		fmt.Println("Use as `go run main.go Kyiv`")
		os.Exit(1)
	}
	city := os.Args[1]
	//city = "Kyiv"
	//city = "Kiev"

	w, err := server.HandleCity(city)
	j, _ := json.Marshal(w)
	fmt.Println(string(j))
	fmt.Println(err)
	//fmt.Println(w.temp, ' ', w.hum, ' ', w.wind)
}
