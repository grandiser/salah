package main

import (
	"fmt"
	"github.com/grandiser/salah/geo"
	"github.com/grandiser/salah/ip"
	"github.com/grandiser/salah/times"
)
import "flag"

func main() {
	city := flag.String("city", "montreal", "The city to get prayer times for.")
	country := flag.String("country", "canada", "The country where the city is")
	flag.Parse()

	var cityProvided bool
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "city" {
			cityProvided = true
		}
	})

	var countryProvided bool
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "country" {
			countryProvided = true
		}
	})

	if !cityProvided && !countryProvided {

		userIp, err := ip.LocalIpApi()

		if err != nil {
			fmt.Println("IP API Not Available. Attempting to use defaults")
			aladhan_times, err := times.AladhanLocationAPI(*city, *country)

			if err != nil {
				fmt.Println("AlAdhan API Not Available (default location). Try again later")
				panic(err)
			}

			fmt.Printf("\n aladhan fajr time: " + aladhan_times.Data.Timings.Fajr)
		}

		location, err := geo.LocationAPI(userIp)

		if err != nil {
			fmt.Println("IP Encoding API Unavailable. Attempting to use IslamicFinder API with IP address")
			islamic_finder_times, err := times.IslamicFinderAPI(userIp)

			if err != nil {
				fmt.Println("Islamic Finder API Not Available (Public IP). Try again later")
			}

			fmt.Println("islamic finder fajr timr: " + islamic_finder_times.Results.Fajr)
		}

		aladhan_times := times.AladhanCoordsAPI(location.Lat, location.Lon)
		fmt.Printf("aladhan best fajr: " + aladhan_times.Data.Timings.Fajr)
	}

	if cityProvided && !countryProvided {
		geoEncoding, err := geo.OpenMeteoAPI(*city)

		if err != nil {
			fmt.Println("Geo Encoding API Not Available. Use both --city and --country flags")
		}

		aladhan_times := times.AladhanCoordsAPI(geoEncoding.Latitude, geoEncoding.Longitude)

		fmt.Printf("aladhan: " + aladhan_times.Data.Timings.Fajr)
	}

	if !cityProvided && countryProvided {
		fmt.Println("The --country flag is meant to be used alongside a --city flag, not on its own")
	}

	if cityProvided && countryProvided {
		aladhanTimes, err := times.AladhanLocationAPI(*city, *country)

		if err != nil {

			fmt.Println("AlAdhan API Not Available (city + country). Attempting to encode city coordinates")
			geoEncoding, err := geo.OpenMeteoAPI(*city)

			if err != nil {
				fmt.Println("Geo Encoding API Not Available. Try again later")
			}

			aladhanTimes := times.AladhanCoordsAPI(geoEncoding.Latitude, geoEncoding.Longitude)
			fmt.Printf("aladhan: " + aladhanTimes.Data.Timings.Fajr)
		}

		fmt.Println("aladhan fajr: " + aladhanTimes.Data.Timings.Fajr)
	}

	fmt.Printf("\nPrayer times for %s\n", *city)

}
