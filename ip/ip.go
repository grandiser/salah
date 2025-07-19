package ip

import (
	"encoding/json"
	"io"
	"net/http"
)

type IPAPIResponse struct {
	Origin string `json:"origin"`
}

func LocalIpApi() (string, error) {
	resp, err := http.Get("https://httpbin.org/ip")

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("IP Detection API Unavailable. Pass in City name using --city 'city'")
	}

	if err != nil {
		panic(err)
	}

	var ip IPAPIResponse

	err = json.Unmarshal(body, &ip)

	if err != nil {
		panic(err)
	}

	return ip.Origin, err

}
