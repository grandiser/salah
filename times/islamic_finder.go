package times

import (
	"encoding/json"
	"io"
	"net/http"
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

func IslamicFinderAPI(userIp string) (IslamicFinder, error) {
	var baseApi string = "https://www.islamicfinder.us/index.php/api/prayer_times"
	var ipApi string = baseApi + "?user_ip=" + userIp

	resp, err := http.Get(ipApi)

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		panic("Islamic Finder API Not Available. Pass in City name using --city 'city_name'")
	}

	var islamicFinder IslamicFinderAPIResponse
	err = json.Unmarshal(body, &islamicFinder)

	return IslamicFinder{
		Results:  islamicFinder.Results,
		Location: islamicFinder.Settings.Location,
	}, err

}
