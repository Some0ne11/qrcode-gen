package api

import (
	"qrcode-gen/internal/api/handlers"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) {
	r.Post("/api/generate", handlers.HandleGenerate)
	r.Post("/api/scan", handlers.HandleScan)
}
