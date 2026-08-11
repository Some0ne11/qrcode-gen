package models

type GenerateRequest struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type ScanResponse struct {
	Result string `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
