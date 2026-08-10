package models

type GenerateRequest struct {
	Text string `json:"text"`
}

type ScanResponse struct {
	Result string `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
