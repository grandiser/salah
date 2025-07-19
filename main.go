package main

import (
	"encoding/json"
	"fmt"
	"github.com/grandiser/salah/ip"
	"io"
	"net/http"
	"strings"
)
import "flag"

type IslamicFinder struct {
	Results struct {
		Fajr    string `json:"Fajr"`
		Duha    string `json:"Duha"`
		Dhuhr   string `json:"Dhuhr"`
		Asr     string `json:"Asr"`
		Maghrib string `json:"Maghrib"`
		Isha    string `json:"Isha"`
	} `json:"results"`
	Settings struct {
		Name     string `json:"name"`
		Location struct {
			City    string `json:"city"`
			State   string `json:"state"`
			Country string `json:"country"`
		} `json:"location"`
		Latitude     string `json:"latitude"`
		Longitude    string `json:"longitude"`
		Timezone     string `json:"timezone"`
		Method       int    `json:"method"`
		Juristic     int    `json:"juristic"`
		HighLatitude int    `json:"high_latitude"`
		FajirRule    struct {
			Type  int `json:"type"`
			Value int `json:"value"`
		} `json:"fajir_rule"`
		MaghribRule struct {
			Type  int `json:"type"`
			Value int `json:"value"`
		} `json:"maghrib_rule"`
		IshaRule struct {
			Type  int `json:"type"`
			Value int `json:"value"`
		} `json:"isha_rule"`
		TimeFormat int `json:"time_format"`
	} `json:"settings"`
	Success bool `json:"success"`
}

func islamic_finder_api(user_ip string) (IslamicFinder, error) {
	var base_api string = "https://www.islamicfinder.us/index.php/api/prayer_times"
	var ip_api string = base_api + "?user_ip=" + user_ip

	resp, err := http.Get(ip_api)

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("Islamic Finder API Not Available. Pass in City name using --city 'city'")
	}

	if err != nil {
		panic(err)
	}

	var islamic_finder IslamicFinder
	err = json.Unmarshal(body, &islamic_finder)

	return islamic_finder, err

}

func get_prayer_times_aladhan(latitude string, longitude string) string {
	var api_call string = fmt.Sprintf("https://api.aladhan.com/v1/timings/18-07-2025?latitude=%s&longitude=%s", latitude, longitude)

	resp, err := http.Get(api_call)

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("Prayer API Not Available.")
	}

	return string(body)
}

func convert_city_name_format(city string) string {
	var corrected_city string = strings.ReplaceAll(city, " ", "+")

	return corrected_city
}

func main() {
	city := flag.String("city", "montreal", "The city to get prayer times for.")
	flag.Parse()

	//var corr_city string = convert_city_name_format(*city)

	//coords := get_response_body_open_meteo(corr_city)

	user_ip := ip.LocalIpApi()

	islamic_finder_resp, err := islamic_finder_api(user_ip)

	if err != nil {
		panic(err)
	}

	prayer_times := islamic_finder_resp.Results
	city_name, country_name := islamic_finder_resp.Settings.Location.City, islamic_finder_resp.Settings.Location.Country

	fmt.Println(prayer_times)
	fmt.Println(city_name + ", " + country_name)

	fmt.Printf("Prayer times for %s\n", *city)

}
