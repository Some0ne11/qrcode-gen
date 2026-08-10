package handlers

import (
	"encoding/json"
	"net/http"
	"qrcode-gen/internal/models"

	goqrcode "github.com/skip2/go-qrcode"
)

func HandleGenerate(w http.ResponseWriter, r *http.Request) {
	var req models.GenerateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if req.Text == "" {
		RespondWithError(w, http.StatusBadRequest, "Text cannot be empty")
		return
	}

	// Generate QR code
	png, err := goqrcode.Encode(req.Text, goqrcode.Medium, 256)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to generate QR code")
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(png)
}
