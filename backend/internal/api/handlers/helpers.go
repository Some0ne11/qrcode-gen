package handlers

import (
	"encoding/json"
	"net/http"
	"qrcode-gen/internal/models"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(models.ErrorResponse{Error: message})
}
