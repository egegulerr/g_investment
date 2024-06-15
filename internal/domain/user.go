package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string
	FirstName    string
	LastName     string
	Password     string
	FavoriteNews []News  `gorm:"many2many:user_favorite_news;"`
	Stocks       []Stock `gorm:"many2many:user_stocks;"`
}

func NewUser(email string, firstName string, lastName string, password string, favoriteNews []News, stocks []Stock) *User {
	return &User{
		Email:        email,
		FirstName:    firstName,
		LastName:     lastName,
		Password:     password,
		FavoriteNews: favoriteNews,
		Stocks:       stocks,
	}
}
