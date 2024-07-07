package main

import (
	"bytes"
	"io"
	"net/http"
)

type SttApiClient struct {
	url string
}

func NewSttApiClient(url string) *SttApiClient {
	return &SttApiClient{url}
}

func (stt SttApiClient) Transcribe(data []byte) (string, error) {
	req, err := http.NewRequest("POST", stt.url, bytes.NewBuffer(data))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
