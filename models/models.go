package models

type LoadTestRequest struct {
	Requests []struct {
		Method          string `json:"method"`
		Endpoint        string `json:"endpoint"`
		Data            string `json:"data"`
		Connections     int    `json:"connections"`
		Rate            int    `json:"rate"`
		ConcurrencyLimit int   `json:"concurrencyLimit"`
	} `json:"requests"`
	Host     string `json:"host"`
	Duration int    `json:"duration"`
}

type LoadTestResponse struct {
	TotalRequests   int     `json:"total_requests"`
	SuccessRate     float64 `json:"success_rate"`
	FailureRate     float64 `json:"failure_rate"`
	AvgResponseTime float64 `json:"avg_response_time"`
}
