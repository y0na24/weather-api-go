package geo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		return &GeoData{
			City: city,
		}, nil
	}

	response, err := http.Get("https://ipapi.co/json/")

	if err != nil {
		return nil, err
	}

	fmt.Println(response.Status)

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var geoData GeoData
	json.Unmarshal(body, &geoData)

	return &geoData, nil
}
