package loginports

import (
	"g_investment/internal/domain"
)

type LoginProvider interface {
	RegisterUser(user *domain.User)
	GetUserWithEmail(email *string) (*domain.User, error)
	GetUserWithID(id *int) (*domain.User, error)
	ParseUserToken(jwtToken *string) (*domain.User, error)
}
