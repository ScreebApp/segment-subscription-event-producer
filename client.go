package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var client *http.Client

func init() {
	client = httpClient()
}

func httpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: *concurrency * 2,
		},
		Timeout: 10 * time.Second,
	}
}

func sendWebhookRequest(body interface{}) error {
	strJson, err := json.Marshal(body)
	if err != nil {
		return nil
	}

	req, err := http.NewRequest(http.MethodPost, *endpointURL, bytes.NewBuffer(strJson))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("status code: %s", resp.Status)
	}

	return nil
}
