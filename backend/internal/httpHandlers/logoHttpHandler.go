package httpHandlers

import (
	"g_investment/internal/app/logoService"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type LogoHandler struct {
	service *logoService.LogoService
}

func NewLogoHandler(service *logoService.LogoService) *LogoHandler {
	return &LogoHandler{service: service}
}

func (h *LogoHandler) GetCompanyLogo(w http.ResponseWriter, r *http.Request) {
	ticker := chi.URLParam(r, "ticker")
	img, err := h.service.GetCompanyLogo(&ticker)

	if err != nil {
		http.Error(w, "Failed to fetch logo from logoHandler", http.StatusInternalServerError)
	}

	respond_with_json(w, img)
}
