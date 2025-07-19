package times

import (
	"encoding/json"
	"fmt"
	"github.com/grandiser/salah/geo"
	"io"
	"net/http"
	"strconv"
	"time"
)

type AladhanAPIResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   struct {
		Timings struct {
			Fajr       string `json:"Fajr"`
			Sunrise    string `json:"Sunrise"`
			Dhuhr      string `json:"Dhuhr"`
			Asr        string `json:"Asr"`
			Sunset     string `json:"Sunset"`
			Maghrib    string `json:"Maghrib"`
			Isha       string `json:"Isha"`
			Imsak      string `json:"Imsak"`
			Midnight   string `json:"Midnight"`
			Firstthird string `json:"Firstthird"`
			Lastthird  string `json:"Lastthird"`
		} `json:"timings"`
		Date struct {
			Readable  string `json:"readable"`
			Timestamp string `json:"timestamp"`
			Hijri     struct {
				Date    string `json:"date"`
				Format  string `json:"format"`
				Day     string `json:"day"`
				Weekday struct {
					En string `json:"en"`
					Ar string `json:"ar"`
				} `json:"weekday"`
				Month struct {
					Number int    `json:"number"`
					En     string `json:"en"`
					Ar     string `json:"ar"`
					Days   int    `json:"days"`
				} `json:"month"`
				Year        string `json:"year"`
				Designation struct {
					Abbreviated string `json:"abbreviated"`
					Expanded    string `json:"expanded"`
				} `json:"designation"`
				Holidays         []interface{} `json:"holidays"`
				AdjustedHolidays []interface{} `json:"adjustedHolidays"`
				Method           string        `json:"method"`
			} `json:"hijri"`
			Gregorian struct {
				Date    string `json:"date"`
				Format  string `json:"format"`
				Day     string `json:"day"`
				Weekday struct {
					En string `json:"en"`
				} `json:"weekday"`
				Month struct {
					Number int    `json:"number"`
					En     string `json:"en"`
				} `json:"month"`
				Year        string `json:"year"`
				Designation struct {
					Abbreviated string `json:"abbreviated"`
					Expanded    string `json:"expanded"`
				} `json:"designation"`
				LunarSighting bool `json:"lunarSighting"`
			} `json:"gregorian"`
		} `json:"date"`
		Meta struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
			Timezone  string  `json:"timezone"`
			Method    struct {
				Id     int    `json:"id"`
				Name   string `json:"name"`
				Params struct {
					Fajr int `json:"Fajr"`
					Isha int `json:"Isha"`
				} `json:"params"`
				Location struct {
					Latitude  float64 `json:"latitude"`
					Longitude float64 `json:"longitude"`
				} `json:"location"`
			} `json:"method"`
			LatitudeAdjustmentMethod string `json:"latitudeAdjustmentMethod"`
			MidnightMode             string `json:"midnightMode"`
			School                   string `json:"school"`
			Offset                   struct {
				Imsak    int `json:"Imsak"`
				Fajr     int `json:"Fajr"`
				Sunrise  int `json:"Sunrise"`
				Dhuhr    int `json:"Dhuhr"`
				Asr      int `json:"Asr"`
				Maghrib  int `json:"Maghrib"`
				Sunset   int `json:"Sunset"`
				Isha     int `json:"Isha"`
				Midnight int `json:"Midnight"`
			} `json:"offset"`
		} `json:"meta"`
	} `json:"data"`
}

func AladhanCoordsAPI(latitude float64, longitude float64) AladhanAPIResponse {
	year, month, day := time.Now().Date()
	lat := strconv.FormatFloat(latitude, 'f', -1, 32)
	lon := strconv.FormatFloat(longitude, 'f', -1, 32)

	date := fmt.Sprintf("%d-%d-%d", year, month, day)

	var api_call string = fmt.Sprintf("https://api.aladhan.com/v1/timings/%s?latitude=%s&longitude=%s", date, lat, lon)

	resp, err := http.Get(api_call)

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		panic("Prayer API Not Available.")
	}

	defer resp.Body.Close()

	var aladhan AladhanAPIResponse
	err = json.Unmarshal(body, &aladhan)
	if err != nil {
		panic(err)
	}

	return aladhan
}

func AladhanLocationAPI(city_name string, country_name string) (AladhanAPIResponse, error) {
	year, month, day := time.Now().Date()
	fixedCityName := geo.FixCityName(city_name)
	date := fmt.Sprintf("%d-%d-%d", year, month, day)

	var api_call string = fmt.Sprintf("https://api.aladhan.com/v1/timingsByCity/%s?city=%s&country=%s", date, fixedCityName, country_name)

	resp, err := http.Get(api_call)

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		panic("Prayer API Not Available.")
	}

	defer resp.Body.Close()

	var aladhan AladhanAPIResponse
	err = json.Unmarshal(body, &aladhan)
	if err != nil {
		panic(err)
	}

	return aladhan, err
}
