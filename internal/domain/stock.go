package domain

import (
	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model
	Symbol     string
	NewsStocks []NewsStock `gorm:"foreignKey:StockID"`
}

func NewStock(symbol string, news []NewsStock) *Stock {
	return &Stock{
		Symbol:     symbol,
		NewsStocks: news,
	}
}
