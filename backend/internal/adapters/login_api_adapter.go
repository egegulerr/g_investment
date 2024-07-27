package adapters

import (
	"errors"
	"fmt"
	"g_investment/internal/domain"
	"g_investment/internal/domain/dtos"
	loginports "g_investment/internal/ports/loginPorts"

	"gorm.io/gorm"
)

type LoginApiAdapter struct {
	repository *gorm.DB
}

func NewLoginApiAdapter(db *gorm.DB) loginports.LoginProvider {
	return &LoginApiAdapter{
		repository: db,
	}
}

func (adapter *LoginApiAdapter) RegisterUser(user *domain.User) {
	adapter.repository.Create(user)
}

func (adapter *LoginApiAdapter) LoginUser(user *dtos.LoginDTO) {
	fmt.Println("Login")
}

func (adapter *LoginApiAdapter) GetUserWithEmail(email *string) (*domain.User, error) {
	var user domain.User
	if err := adapter.repository.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {

			return nil, errors.New("no user found with email")
		}
		return nil, err
	}
	return &user, nil
}

func (adapter *LoginApiAdapter) GetUserWithID(id *int) (*domain.User, error) {
	var user domain.User
	if err := adapter.repository.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("no user found with id: %d", *id)
		}
		return nil, err
	}
	return &user, nil
}

func (adapter *LoginApiAdapter) ParseUserToken(jwtToken *string) (*domain.User, error) {
	return nil, nil
}
