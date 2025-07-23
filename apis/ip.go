package apis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IPAPIResponse struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

func LocalIpApi() (IPAPIResponse, error) {
	resp, err := http.Get("http://ip-api.com/json/")

	if err != nil {
		return IPAPIResponse{}, fmt.Errorf("failed to make HTTP request: %w", err)
	}

	body, err := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return IPAPIResponse{}, fmt.Errorf("IP Detection API Unavailable (status: %d). Pass in City name using --city 'city'", resp.StatusCode)
	}

	if err != nil {
		return IPAPIResponse{}, fmt.Errorf("%w", err)
	}

	var ip IPAPIResponse

	err = json.Unmarshal(body, &ip)

	if err != nil {
		return IPAPIResponse{}, fmt.Errorf("failed to unmarshal JSON response: %w", err)
	}

	return ip, err

}
