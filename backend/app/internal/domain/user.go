package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string             `gorm:"unique" json:"email"`
	FirstName    string             `json:"first_name"`
	LastName     string             `json:"last_name"`
	Password     []byte             `json:"-"`
	FavoriteNews []UserFavoriteNews `gorm:"foreignKey:UserID;" json:"favorite_news"`
	Stocks       []UserStock        `gorm:"foreignKey:UserID;" json:"stocks"`
}

type UserFavoriteNews struct {
	gorm.Model
	UserID uint
	NewsID uint

	News News `gorm:"foreignKey:NewsID;"`
}

type UserStock struct {
	gorm.Model
	UserID  uint
	StockID uint

	Stock Stock `gorm:"foreignKey:StockID;"`
}

func NewUser(email string, firstName string, lastName string, password []byte, favoriteNews []UserFavoriteNews, stocks []UserStock) *User {
	return &User{
		Email:        email,
		FirstName:    firstName,
		LastName:     lastName,
		Password:     password,
		FavoriteNews: favoriteNews,
		Stocks:       stocks,
	}
}
