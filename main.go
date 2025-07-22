package main

import (
	"fmt"
	"github.com/grandiser/salah/geo"
	"github.com/grandiser/salah/ip"
	"github.com/grandiser/salah/output"
	"github.com/grandiser/salah/times"
	"log"
	"os"
)
import "flag"

func autoBehavior(city string, country string, nextOnly bool) {
	userIp, err := ip.LocalIpApi()

	if err != nil {
		fmt.Println("IP API Not Available. Attempting to use defaults")
		aladhanTimes, err := times.AladhanLocationAPI(city, country)

		if err != nil {
			fmt.Println("AlAdhan API Not Available (default location). Try again later")
			log.Fatalf("Error getting prayer times: %v", err)
		}

		output.AladhanHandler(aladhanTimes, nextOnly)
		return
	}
	location, err := geo.LocationAPI(userIp)
	if err != nil {
		fmt.Println("IP Encoding API Unavailable. Try again by specifying --city and --country flags")
	}

	aladhanTimes, err := times.AladhanCoordsAPI(location.Lat, location.Lon)
	if err != nil {
		fmt.Println("AlAdhan API Not Available (coordinates). Try again later")
		os.Exit(1)
	}
	output.AladhanHandler(aladhanTimes, nextOnly)
}

func cityBehavior(city string, nextOnly bool) {
	geoEncoding, err := geo.OpenMeteoAPI(city)
	if err != nil {
		fmt.Println("Geo Encoding API Not Available. Use both --city and --country flags")
		os.Exit(1)
	}
	aladhanTimes, err := times.AladhanCoordsAPI(geoEncoding.Latitude, geoEncoding.Longitude)
	if err != nil {
		fmt.Println("AlAdhan API Not Available (coordinates). Try again later")
		os.Exit(1)
	}
	output.AladhanHandler(aladhanTimes, nextOnly)
}

func cityCountryBehavior(city string, country string, nextOnly bool) {
	aladhanTimes, err := times.AladhanLocationAPI(city, country)

	// Default behavior using City and Country values
	if err == nil {
		output.AladhanHandler(aladhanTimes, nextOnly)
		return

	} else {
		// Backup Call by GeoEncoding City and using latitude and longitude instead
		geoEncoding, err := geo.OpenMeteoAPI(city)
		if err != nil {
			fmt.Println("Geo Encoding API Not Available. Try again later")
			os.Exit(1)
		}
		aladhanTimes, err := times.AladhanCoordsAPI(geoEncoding.Latitude, geoEncoding.Longitude)
		if err != nil {
			fmt.Println("AlAdhan API Not Available (coordinates). Try again later")
			os.Exit(1)
		}
		output.AladhanHandler(aladhanTimes, nextOnly)
		return
	}
}

func main() {
	city := flag.String("city", "", "The city to get prayer times for")
	country := flag.String("country", "", "The country where the city is located")
	nextOnly := flag.Bool("nextonly", false, "Show current prayer only")
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
		autoBehavior(*city, *country, *nextOnly)
	}

	if cityProvided && !countryProvided {
		cityBehavior(*city, *nextOnly)
	}

	if !cityProvided && countryProvided {
		fmt.Println("The --country flag is meant to be used alongside a --city flag, not on its own")
		return
	}

	if cityProvided && countryProvided {
		cityCountryBehavior(*city, *country, *nextOnly)
	}

	//fmt.Println("       ﷽   ")
	//fmt.Println("   ╭────────۩────────╮")
	//fmt.Println("   │ Fajr    : 03:38 │")
	//fmt.Println("   │ Sunrise : 05:27 │")
	//fmt.Println("   │ Dhuhr   : 13:02 │")
	//fmt.Println("   │ Asr     : 17:08 │")
	//fmt.Println("   │ Maghrib : 20:36 │")
	//fmt.Println("   │ Isha    : 22:24 │")
	//fmt.Println("   ╰────────۞────────╯")
	//
	//fmt.Println("\n Next: Maghrib ▣▣▣▣▣▣▣▢▢▢ 34m\n")
}
