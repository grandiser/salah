package times

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type IslamicFinderAPIResponse struct {
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
type IslamicFinder struct {
	Results struct {
		Fajr    string `json:"Fajr"`
		Duha    string `json:"Duha"`
		Dhuhr   string `json:"Dhuhr"`
		Asr     string `json:"Asr"`
		Maghrib string `json:"Maghrib"`
		Isha    string `json:"Isha"`
	} `json:"results"`
	Location struct {
		City    string `json:"city"`
		State   string `json:"state"`
		Country string `json:"country"`
	} `json:"location"`
}

func IslamicFinderAPI(user_ip string) (IslamicFinder, error) {
	var base_api string = "https://www.islamicfinder.us/index.php/api/prayer_times"
	var ip_api string = base_api + "?user_ip=" + user_ip

	resp, err := http.Get(ip_api)

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("Islamic Finder API Not Available. Pass in City name using --city 'city_name'")
	}

	if err != nil {
		panic(err)
	}

	var islamic_finder IslamicFinderAPIResponse
	err = json.Unmarshal(body, &islamic_finder)

	return IslamicFinder{
		Results:  islamic_finder.Results,
		Location: islamic_finder.Settings.Location,
	}, err

}

func OutputListIslamicFinder(islamicFinderTimes IslamicFinder) {
	fajr := &islamicFinderTimes.Results.Fajr
	sunrise := &islamicFinderTimes.Results.Duha
	dhuhr := &islamicFinderTimes.Results.Dhuhr
	asr := &islamicFinderTimes.Results.Asr
	maghrib := &islamicFinderTimes.Results.Maghrib
	isha := &islamicFinderTimes.Results.Isha

	today := time.Now()

	fmt.Println(today)
	fmt.Printf("Fajr: %s\nSunrise: %s\nDhuhr: %s\nAsr: %s\nMaghrib: %s\nIsha: %s\n", *fajr, *sunrise, *dhuhr, *asr, *maghrib, *isha)

}
