package loginservice

import (
	"errors"
	"fmt"
	"g_investment/internal/domain"
	"g_investment/internal/domain/dtos"
	loginports "g_investment/internal/ports/loginPorts"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

type LoginService struct {
	provider loginports.LoginProvider
	jwtKey   *string
}

func NewLoginService(provider loginports.LoginProvider, jwtKey *string) *LoginService {
	return &LoginService{provider: provider, jwtKey: jwtKey}
}

func (s *LoginService) RegisterUser(payload *dtos.LoginDTO) {
	password, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), 14)
	user := domain.NewUser(payload.Email, payload.FirstName, payload.LastName, password, []domain.UserFavoriteNews{}, []domain.UserStock{})
	s.provider.RegisterUser(user)
}

func (s *LoginService) LoginUser(payload *dtos.LoginDTO) (string, error) {
	user, err := s.provider.GetUserWithEmail(&payload.Email)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(payload.Password)); err != nil {
		return "", errors.New("invalid password")
	}

	return s.createJwtToken(user)
}

func (s *LoginService) createJwtToken(user *domain.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	})
	token, err := claims.SignedString([]byte(*s.jwtKey))
	if err != nil {
		fmt.Println(*s.jwtKey)
		return "", errors.New("could not login. Error creating jwt token")
	}

	return token, nil
}

func (s *LoginService) ParseUserTokenAndGetUser(jwtToken *string) (*domain.User, error) {
	token, err := jwt.ParseWithClaims(*jwtToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(*s.jwtKey), nil
	})

	if err != nil {
		return nil, errors.New("unauthenticated")
	}

	claims := token.Claims
	issuer, err := claims.GetIssuer()
	if err != nil {
		return nil, errors.New("cant find issuer")
	}
	issuerID, err := strconv.Atoi(issuer)
	if err != nil {
		return nil, errors.New("cant convert issuer to int")
	}
	user, err := s.provider.GetUserWithID(&issuerID)
	if err != nil {
		return nil, fmt.Errorf("no user found with id: %d", issuerID)
	}

	return user, nil
}
