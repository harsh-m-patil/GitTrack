package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// A generic function to handle HTTP GET requests and JSON unmarshalling
func GetResponse[T any](url string) (*T, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("User not found")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to get a valid response, Status: %d\n", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %v\n", err)
	}

	var response T
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling JSON: %v\n", err)
	}

	return &response, nil
}
