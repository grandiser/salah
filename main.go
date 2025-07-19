package main

import (
	"fmt"
	"github.com/grandiser/salah/geo"
	"github.com/grandiser/salah/ip"
	"github.com/grandiser/salah/times"
	"log"
	"os"
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
			aladhanTimes, err := times.AladhanLocationAPI(*city, *country)

			if err != nil {
				fmt.Println("AlAdhan API Not Available (default location). Try again later")
				log.Fatalf("Error getting prayer times: %v", err)
			}

			fmt.Printf("\n aladhan fajr time: " + aladhanTimes.Data.Timings.Fajr)
			return
		}

		location, err := geo.LocationAPI(userIp)
		if err != nil {
			fmt.Println("IP Encoding API Unavailable. Attempting to use IslamicFinder API with IP address")
			islamicFinderTimes, err := times.IslamicFinderAPI(userIp)

			if err != nil {
				fmt.Println("Islamic Finder API Not Available (Public IP). Try again later")
				os.Exit(1)
			}
			fmt.Println("islamic finder fajr timr: " + islamicFinderTimes.Results.Fajr)
			return
		}

		aladhanTimes, err := times.AladhanCoordsAPI(location.Lat, location.Lon)
		if err != nil {
			fmt.Println("AlAdhan API Not Available (coordinates). Try again later")
			os.Exit(1)
		}

		fmt.Printf("aladhan best fajr: " + aladhanTimes.Data.Timings.Fajr)
	}

	if cityProvided && !countryProvided {
		geoEncoding, err := geo.OpenMeteoAPI(*city)

		if err != nil {
			fmt.Println("Geo Encoding API Not Available. Use both --city and --country flags")
			os.Exit(1)
		}

		aladhanTimes, err := times.AladhanCoordsAPI(geoEncoding.Latitude, geoEncoding.Longitude)
		if err != nil {
			fmt.Println("AlAdhan API Not Available (coordinates). Try again later")
			os.Exit(1)
		}

		fmt.Printf("aladhan: " + aladhanTimes.Data.Timings.Fajr)
	}

	if !cityProvided && countryProvided {
		fmt.Println("The --country flag is meant to be used alongside a --city flag, not on its own")
		return
	}

	if cityProvided && countryProvided {
		aladhanTimes, err := times.AladhanLocationAPI(*city, *country)

		if err != nil {
			fmt.Println("AlAdhan API Not Available (city + country). Attempting to encode city coordinates")
			geoEncoding, err := geo.OpenMeteoAPI(*city)

			if err != nil {
				fmt.Println("Geo Encoding API Not Available. Try again later")
				os.Exit(1)
			}

			aladhanTimes, err := times.AladhanCoordsAPI(geoEncoding.Latitude, geoEncoding.Longitude)

			if err != nil {
				fmt.Println("AlAdhan API Not Available (coordinates). Try again later")
				os.Exit(1)
			}

			fmt.Printf("aladhan: " + aladhanTimes.Data.Timings.Fajr)
			return
		}
		fmt.Println("aladhan fajr: " + aladhanTimes.Data.Timings.Fajr)
	}

	fmt.Printf("\nPrayer times for %s\n", *city)
}
