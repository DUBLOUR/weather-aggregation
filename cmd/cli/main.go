package main

import (
	"encoding/json"
	"fmt"
	"github.com/DUBLOUR/weather-aggregation/internal/server"
	"os"
)

func main() {
	fmt.Println(os.Args)
	if !(len(os.Args) == 2) {
		fmt.Println("Use as `go run main.go Kyiv`")
		os.Exit(1)
	}
	city := os.Args[1]

	w, _ := server.HandleCity(city)
	j, _ := json.Marshal(w)
	fmt.Println(string(j))
}
