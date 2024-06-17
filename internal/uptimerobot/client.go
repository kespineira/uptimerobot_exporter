package uptimerobot

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/kespineira/uptimerobot_exporter/config"
)

type Client struct {
	apiKey string
}

func NewClient() *Client {
	return &Client{
		apiKey: config.ApiKey,
	}
}

func (c *Client) GetMonitors() ([]Monitor, error) {
	url := "https://api.uptimerobot.com/v2/getMonitors"
	data := []byte(`{"api_key":"` + c.apiKey + `", "format":"json", "response_times":"1", "response_times_average":"30"}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	var result monitorsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Error decoding response:", err)
		return nil, err
	}

	return result.Monitors, nil
}

func (c *Client) GetAccountDetails() (AccountDetails, error) {
	url := "https://api.uptimerobot.com/v2/getAccountDetails"
	data := []byte(`{"api_key": "` + c.apiKey + `", "format": "json"}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Println("Error creating request:", err)
		return AccountDetails{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making request:", err)
		return AccountDetails{}, err
	}
	defer resp.Body.Close()

	var result accountDetailsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Error decoding response:", err)
		return AccountDetails{}, err
	}

	return result.Account, nil
}
