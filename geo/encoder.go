package geo

import (
	"io"
	"net/http"
)

func OpenMeteoAPI(city string) string {
	var base_api string = "https://geocoding-api.open-meteo.com/v1/search?name="
	var city_api string = base_api + city
	resp, err := http.Get(city_api)

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		panic("Geo-Encoding (Open Meteo) API Not Available.")
	}

	defer resp.Body.Close()

	return string(body)
}

func GetLatLong(body string) string {
	return "1"
}
