package weather

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"weather-api/geo"
)

func GetWeather(geoData geo.GeoData, format string) string {
	baseUrl, err := url.Parse("https://wttr.in/" + geoData.City)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	params := url.Values{}
	params.Add("format", format)

	baseUrl.RawQuery = params.Encode()

	response, err := http.Get(baseUrl.String())

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	body, err := io.ReadAll(response.Body)

	return string(body)
}
