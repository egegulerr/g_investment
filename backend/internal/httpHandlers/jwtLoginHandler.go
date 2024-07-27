package httpHandlers

import (
	"encoding/json"
	"fmt"
	loginservice "g_investment/internal/app/loginService"
	"g_investment/internal/domain/dtos"
	"net/http"
	"time"
)

type JwtLoginHandler struct {
	service *loginservice.LoginService
}

func NewJwtLoginHandler(service *loginservice.LoginService) *JwtLoginHandler {
	return &JwtLoginHandler{service: service}
}

func (h *JwtLoginHandler) Register(w http.ResponseWriter, r *http.Request) {
	var payload dtos.LoginDTO
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}
	h.service.RegisterUser(&payload)
}

func (h *JwtLoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload dtos.LoginDTO
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}
	token, err := h.service.LoginUser(&payload)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
	})
	respond_with_json(w, token)
}

func (h *JwtLoginHandler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(-time.Hour)})
}

func (h *JwtLoginHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		http.Error(w, "No jwt cookie", http.StatusUnauthorized)
		return
	}

	claims, err := h.service.ParseToken(&cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	user, err := h.service.GetUser(*claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	respond_with_json(w, user)
}

func (h *JwtLoginHandler) IsTokenValid(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt")
	fmt.Println(r.Cookies())
	fmt.Print("error iccred")
	if err != nil {
		http.Error(w, "No jwt cookie", http.StatusUnauthorized)
		return
	}

	_, err = h.service.ParseToken(&cookie.Value)
	if err != nil {
		respond_with_json(w, false)
	}
	respond_with_json(w, true)
}
