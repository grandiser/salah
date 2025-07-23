package main

import (
	"fmt"
	"github.com/grandiser/salah/apis"
	"github.com/grandiser/salah/prayers"
	"log"
	"os"
)

func autoBehavior(config prayers.Config) {
	IPLocResponse, err := apis.LocalIpApi()

	if err != nil {
		fmt.Println("IP API Not Available. Attempting to use defaults")
		configLocationBehavior(config)
		return
	}

	aladhanTimes, err := apis.AladhanCoordsAPI(IPLocResponse.Lat, IPLocResponse.Lon)
	if err != nil {
		fmt.Println("IP Encoding API Unavailable. Trying with default location values in config")
		configLocationBehavior(config)
		return
	}

	prayers.AladhanHandler(aladhanTimes, config)
}

func configLocationBehavior(config prayers.Config) {
	aladhanTimes, err := apis.AladhanLocationAPI(config.City, config.Country)

	if err != nil {
		fmt.Println("AlAdhan API Not Available (config location). Try again later")
		log.Fatalf("Error getting prayer times: %v", err)
	}

	prayers.AladhanHandler(aladhanTimes, config)
}

func cityBehavior(config prayers.Config) {
	geoEncoding, err := apis.OpenMeteoAPI(config.City)
	if err != nil {
		fmt.Println("Geo Encoding API Not Available. Use both --city and --country flags")
		os.Exit(1)
	}
	aladhanTimes, err := apis.AladhanCoordsAPI(geoEncoding.Latitude, geoEncoding.Longitude)
	if err != nil {
		fmt.Println("AlAdhan API Not Available (coordinates). Try again later")
		os.Exit(1)
	}
	prayers.AladhanHandler(aladhanTimes, config)
}

func cityCountryBehavior(config prayers.Config) {
	aladhanTimes, err := apis.AladhanLocationAPI(config.City, config.Country)

	// Default behavior using City and Country values
	if err == nil {
		prayers.AladhanHandler(aladhanTimes, config)
		return

	} else {
		// Backup Call by GeoEncoding City and using latitude and longitude instead
		geoEncoding, err := apis.OpenMeteoAPI(config.City)
		if err != nil {
			fmt.Println("Geo Encoding API Not Available. Try again later")
			os.Exit(1)
		}
		aladhanTimes, err := apis.AladhanCoordsAPI(geoEncoding.Latitude, geoEncoding.Longitude)
		if err != nil {
			fmt.Println("AlAdhan API Not Available (coordinates). Try again later")
			os.Exit(1)
		}
		prayers.AladhanHandler(aladhanTimes, config)
		return
	}
}

func main() {
	flags := prayers.ParseFlags()
	config := prayers.ReadConfig()
	prayers.ApplyFlags(&config, flags)

	var cityProvided bool
	var countryProvided bool

	if flags.City != "" {
		cityProvided = true
	}

	if flags.Country != "" {
		countryProvided = true
	}

	if cityProvided && countryProvided {
		cityCountryBehavior(config)
		return
	}

	if cityProvided {
		cityBehavior(config)
		return
	}

	if config.LocateByIp {
		autoBehavior(config)
		return
	} else {
		configLocationBehavior(config)
	}
}
