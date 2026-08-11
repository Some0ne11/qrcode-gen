package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
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

	// Validate payload format on the backend based on type
	switch req.Type {
	case "url":
		if !strings.HasPrefix(req.Text, "http://") && !strings.HasPrefix(req.Text, "https://") {
			RespondWithError(w, http.StatusBadRequest, "Invalid URL format: missing http/https")
			return
		}
	case "phone":
		if !strings.HasPrefix(req.Text, "tel:") || len(req.Text) <= 4 {
			RespondWithError(w, http.StatusBadRequest, "Invalid phone format")
			return
		}
	case "wifi":
		if !strings.HasPrefix(req.Text, "WIFI:") || !strings.Contains(req.Text, "S:") {
			RespondWithError(w, http.StatusBadRequest, "Invalid WiFi format: missing SSID")
			return
		}
	case "contact":
		if !strings.HasPrefix(req.Text, "BEGIN:VCARD") || !strings.Contains(req.Text, "FN:") {
			RespondWithError(w, http.StatusBadRequest, "Invalid contact format: missing Full Name")
			return
		}
	case "text":
		// text can be anything as long as it's not empty
	default:
		// If type is empty or unknown, we just fall back to accepting any non-empty text
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
