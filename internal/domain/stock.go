package domain

import (
	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model
	Symbol string `gorm:"uniqueIndex" json:"symbol"`
}

func NewStock(symbol string, news []NewsStock) *Stock {
	return &Stock{
		Symbol: symbol,
	}
}
