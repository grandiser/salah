package apis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type LocationAPIResponse struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

type Location struct {
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
}

type GeoEncodingApiResponse struct {
	Results []struct {
		Id          int     `json:"id"`
		Name        string  `json:"name"`
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitude"`
		Elevation   float64 `json:"elevation"`
		FeatureCode string  `json:"feature_code"`
		CountryCode string  `json:"country_code"`
		Admin1Id    int     `json:"admin1_id"`
		Admin2Id    int     `json:"admin2_id"`
		Admin3Id    int     `json:"admin3_id"`
		Timezone    string  `json:"timezone"`
		Population  int     `json:"population"`
		CountryId   int     `json:"country_id"`
		Country     string  `json:"country"`
		Admin1      string  `json:"admin1"`
		Admin2      string  `json:"admin2"`
		Admin3      string  `json:"admin3"`
	} `json:"results"`
	GenerationtimeMs float64 `json:"generationtime_ms"`
}

type GeoEncoding struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Country   string  `json:"country"`
}

func FixCityName(city string) string {
	var corrected_city string = strings.ReplaceAll(city, " ", "+")

	return corrected_city
}

func OpenMeteoAPI(city string) (GeoEncoding, error) {
	fixedCityName := FixCityName(city)

	var base_api = "https://geocoding-api.open-meteo.com/v1/search?name="
	var city_api = base_api + fixedCityName + "&count=1"
	resp, err := http.Get(city_api)

	if err != nil {
		return GeoEncoding{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return GeoEncoding{}, fmt.Errorf("Geo-Encoding (Open Meteo) API Not Available.")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GeoEncoding{}, err
	}

	var geo GeoEncodingApiResponse
	err = json.Unmarshal(body, &geo)
	if err != nil {
		return GeoEncoding{}, err
	}

	if len(geo.Results) == 0 {
		fmt.Printf("Could not geo-encode city provided.")
		return GeoEncoding{}, fmt.Errorf("Could not geo-encode city provided.")
	}

	result := geo.Results[0]

	return GeoEncoding{
		Latitude:  result.Latitude,
		Longitude: result.Longitude,
		Country:   result.Country,
	}, nil
}

func LocationAPI(ip string) (Location, error) {
	var apiCall = "http://ip-api.com/json/" + ip

	resp, err := http.Get(apiCall)

	if err != nil {
		return Location{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Location{}, fmt.Errorf("IP to Location API Not Available.")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	var loc LocationAPIResponse
	err = json.Unmarshal(body, &loc)
	if err != nil {
		return Location{}, err
	}

	return Location{
		Country:     loc.Country,
		CountryCode: loc.CountryCode,
		Region:      loc.Region,
		RegionName:  loc.RegionName,
		City:        loc.City,
		Zip:         loc.Zip,
		Lat:         loc.Lat,
		Lon:         loc.Lon,
		Timezone:    loc.Timezone,
	}, nil

}
