package main

import (
	"flag"
	"fmt"
	"weather-api/geo"
	"weather-api/weather"
)

func main() {
	city := flag.String("city", "Nizhniy Novgorod", "Город пользователя")
	format := flag.String("format", "1", "Формат вывода погоды")

	flag.Parse()

	geoData, err := geo.GetMyLocation(*city)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(geoData)
	weatherData := weather.GetWeather(*geoData, *format)

	fmt.Println(weatherData)
}
