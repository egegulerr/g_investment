package middleware

import (
	loginservice "g_investment/internal/app/loginService"
	"net/http"
)

type AuthMiddleWare struct {
	service *loginservice.LoginService
}

func NewAuthMiddleware(service *loginservice.LoginService) *AuthMiddleWare {
	return &AuthMiddleWare{service: service}
}

func (m *AuthMiddleWare) JwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")
		if err != nil {
			http.Error(w, "No jwt cookie", http.StatusUnauthorized)
			return
		}

		_, err = m.service.ParseToken(&cookie.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
