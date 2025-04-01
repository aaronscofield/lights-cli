package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	ErrConnection      = errors.New("connection error")
	ErrNotFound        = errors.New("not found")
	ErrInvalidResponse = errors.New("invalid server response")
	ErrInvalid         = errors.New("invalid data")
	ErrNotNumber       = errors.New("not a number")
)

type stateProperties map[string]interface{}

type stateData struct {
	Device     string            `json:"device"`
	Model      string            `json:"model"`
	Properties []stateProperties `json:"properties"`
}

type stateResponse struct {
	Data    stateData `json:"data"`
	Message string    `json:"message"`
	Code    int       `json:"code"`
}

type turnResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

type cmd struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type body struct {
	Device string `json:"device"`
	Model  string `json:"model"`
	Cmd    cmd    `json:"cmd"`
}

func newClient() *http.Client {
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	return c
}

func state(apiRoot string, apiKey string, deviceModel string, deviceId string) (stateResponse, error) {
	// Define the request with required headers
	url := fmt.Sprintf("%s/state?device=%s&model=%s", apiRoot, deviceId, deviceModel)

	fmt.Println(url)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Govee-API-Key", apiKey)

	// Execute the request
	r, err := newClient().Do(req)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

	// Construct the response object
	var resp stateResponse
	json.NewDecoder(r.Body).Decode(&resp)

	fmt.Printf("%+v", resp)
	return resp, nil
}

func turn(apiRoot string, apiKey string, deviceModel string, deviceId string, operation string) (turnResponse, error) {
	// Construct the payload
	body := body{
		Device: deviceId,
		Model:  deviceModel,
		Cmd:    cmd{"turn", operation}, // Only use if field order is stable
	}
	bod, _ := json.Marshal(body)

	// Define the request with required headers
	url := fmt.Sprintf("%s/control", apiRoot)
	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewReader(bod))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Govee-API-Key", apiKey)

	// Execute the request
	r, err := newClient().Do(req)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

	// Construct the response object
	var resp turnResponse
	json.NewDecoder(r.Body).Decode(&resp)

	fmt.Printf("%+v", resp)
	return resp, nil
}
