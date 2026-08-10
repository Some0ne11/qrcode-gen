package main

import (
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	goqrcode "github.com/skip2/go-qrcode"
)

type GenerateRequest struct {
	Text string `json:"text"`
}

type ScanResponse struct {
	Result string `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"}, // Adjust this in production
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	r.Post("/api/generate", handleGenerate)
	r.Post("/api/scan", handleScan)

	port := 8080
	fmt.Printf("Server starting on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

func handleGenerate(w http.ResponseWriter, r *http.Request) {
	var req GenerateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if req.Text == "" {
		respondWithError(w, http.StatusBadRequest, "Text cannot be empty")
		return
	}

	// Generate QR code
	png, err := goqrcode.Encode(req.Text, goqrcode.Medium, 256)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to generate QR code")
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(png)
}

func handleScan(w http.ResponseWriter, r *http.Request) {
	// Limit upload size to 10MB
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("image")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid file upload")
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to decode image. Ensure it's a valid PNG/JPEG.")
		return
	}

	// Prepare image for zxing
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to process image for scanning")
		return
	}

	// Decode QR Code
	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		// gozxing returns a specific error if not found, we can just say not found
		respondWithError(w, http.StatusBadRequest, "No QR code found in image or failed to read")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ScanResponse{Result: result.GetText()})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}
