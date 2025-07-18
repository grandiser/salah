package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)
import "flag"

func get_user_local_ip() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

func get_response_body_open_meteo(city string) string {
	var base_api string = "https://geocoding-api.open-meteo.com/v1/search?name="
	var city_api string = base_api + city
	resp, err := http.Get(city_api)

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		panic("Open Meteo API Not Available.")
	}

	defer resp.Body.Close()

	return string(body)
}

func get_lat_long(body string) string {
	return "1"
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

	var corr_city string = convert_city_name_format(*city)

	coords := get_response_body_open_meteo(corr_city)

	fmt.Println(coords)
	fmt.Printf("Prayer times for %s\n", *city)
}
