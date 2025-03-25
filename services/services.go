package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"testerbackend/models"
)

func ExecuteLoadTest(request models.LoadTestRequest) (*models.LoadTestResponse, error) {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("http://localhost:8080/api/run-test", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, errors.New("Load tester service unreachable")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Failed to read response from load tester")
	}

	var result models.LoadTestResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("Invalid response format from load tester")
	}

	return &result, nil
}
