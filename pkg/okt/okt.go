package okt

import (
	"bytes"
	"cmp"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type OKT struct {
	endpoint string
	apiKey   string
}

func New(endpoint, apiKey string) (*OKT, error) {
	if endpoint == "" {
		endpoint = "http://localhost:8000"
	}
	apiKey = cmp.Or(apiKey, os.Getenv("OKT_APIKEY"))
	if apiKey == "" {
		return nil, fmt.Errorf("OKT_APIKEY environment variable is not set")
	}

	return &OKT{endpoint: endpoint, apiKey: apiKey}, nil
}

func (o *OKT) NormalizeText(text string) (string, error) {
	payload, _ := json.Marshal(textRequest{Text: text})
	req, err := http.NewRequest("POST", o.endpoint+"/normalize", bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+o.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code: %d", resp.StatusCode)
	}

	var res normalizeResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return res.Normalized, nil
}

func (o *OKT) ExtractPhrases(text string) ([]string, error) {
	payload, _ := json.Marshal(textRequest{Text: text})
	req, err := http.NewRequest("POST", o.endpoint+"/phrases", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+o.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	var res phrasesResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return res.Phrases, nil
}

type textRequest struct {
	Text string `json:"text"`
}

type normalizeResponse struct {
	Normalized string `json:"normalized"`
}

type phrasesResponse struct {
	Phrases []string `json:"phrases"`
}
