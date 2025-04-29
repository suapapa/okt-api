package okt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type OKT struct {
	baseURL string
	token   string
}

func NewOKT(baseURL string) (*OKT, error) {
	if baseURL == "" {
		baseURL = "http://localhost:8000"
	}
	token := os.Getenv("OKT_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("OKT_TOKEN environment variable is not set")
	}

	return &OKT{baseURL: baseURL, token: token}, nil
}

func (o *OKT) NormalizeText(text string) (string, error) {
	payload, _ := json.Marshal(textRequest{Text: text})
	req, err := http.NewRequest("POST", o.baseURL+"/normalize", bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+o.token)

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
	req, err := http.NewRequest("POST", o.baseURL+"/phrases", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+o.token)

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
