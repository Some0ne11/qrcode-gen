package handlers

import (
	"encoding/json"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"qrcode-gen/internal/models"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func HandleScan(w http.ResponseWriter, r *http.Request) {
	// Limit upload size to 10MB
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("image")
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid file upload")
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Failed to decode image. Ensure it's a valid PNG/JPEG.")
		return
	}

	// Prepare image for zxing
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to process image for scanning")
		return
	}

	// Decode QR Code
	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		// gozxing returns a specific error if not found, we can just say not found
		RespondWithError(w, http.StatusBadRequest, "No QR code found in image or failed to read")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.ScanResponse{Result: result.GetText()})
}
