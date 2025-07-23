package apis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IPAPIResponse struct {
	Origin string `json:"origin"`
}

func LocalIpApi() (string, error) {
	resp, err := http.Get("https://httpbin.org/ip")

	if err != nil {
		return "", fmt.Errorf("failed to make HTTP request: %w", err)
	}

	body, err := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("IP Detection API Unavailable (status: %d). Pass in City name using --city 'city'", resp.StatusCode)
	}

	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	var ip IPAPIResponse

	err = json.Unmarshal(body, &ip)

	if err != nil {
		return "", fmt.Errorf("failed to unmarshal JSON response: %w", err)
	}

	return ip.Origin, err

}
