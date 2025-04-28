package main

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
)

type TextRequest struct {
    Text string `json:"text"`
}

type NormalizeResponse struct {
    Normalized string `json:"normalized"`
}

type PhrasesResponse struct {
    Phrases []string `json:"phrases"`
}

func NormalizeText(text string) (string, error) {
    payload, _ := json.Marshal(TextRequest{Text: text})
    resp, err := http.Post("http://localhost:8000/normalize", "application/json", bytes.NewBuffer(payload))
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

    var res NormalizeResponse
    json.Unmarshal(body, &res)
    return res.Normalized, nil
}

func ExtractPhrases(text string) ([]string, error) {
    payload, _ := json.Marshal(TextRequest{Text: text})
    resp, err := http.Post("http://localhost:8000/phrases", "application/json", bytes.NewBuffer(payload))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

    var res PhrasesResponse
    json.Unmarshal(body, &res)
    return res.Phrases, nil
}
