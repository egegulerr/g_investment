package httpHandler

import (
	"encoding/json"
	"g_investment/internal/app"
	"g_investment/internal/domain"
	"net/http"
)

type NewsHandler struct {
	service *app.NewsService
}

func NewNewsHandler(service *app.NewsService) *NewsHandler {
	return &NewsHandler{service: service}
}

func (h *NewsHandler) GetCompanyAndMarketNews(w http.ResponseWriter, r *http.Request) {
	news, err := h.service.GetNews()
	if err != nil {
		http.Error(w, "Failed to fetch news from newsHandler", http.StatusInternalServerError)
		return
	}
	respond_with_json(w, news)
}

func (h *NewsHandler) SaveUserFavoriteNews(w http.ResponseWriter, r *http.Request) {
	var news domain.News

	if err := json.NewDecoder(r.Body).Decode(&news); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *NewsHandler) UpdateNews(w http.ResponseWriter, r *http.Request) {
	//id := chi.URLParam(r, "id")
	var news domain.News

	if err := json.NewDecoder(r.Body).Decode(&news); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func respond_with_json(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
