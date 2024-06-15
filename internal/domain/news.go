package domain

import (
	"time"

	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	Url                      string
	Title                    string
	Author                   string
	SentimentalAnalysisScore float64
	Date                     time.Time
	Summary                  string
	Image                    string
	NewsStocks               []NewsStock `gorm:"foreignKey:NewsID"`
}

type NewsStock struct {
	gorm.Model
	NewsID  uint `gorm:"index"`
	StockID uint `gorm:"index"`

	Stock Stock `gorm:"foreignKey:StockID"`
	News  News  `gorm:"foreignKey:NewsID"`

	RelevanceScore      string
	StockSentimentScore string
	StockSentimentLabel string
}

func NewNews(url string, title string, author string, sentimentalAnalysisScore float64, summary string, image string, date time.Time, newsTicker []NewsStock) *News {
	return &News{
		Url:                      url,
		Title:                    title,
		Author:                   author,
		SentimentalAnalysisScore: sentimentalAnalysisScore,
		Date:                     date,
		Summary:                  summary,
		Image:                    image,
		NewsStocks:               newsTicker,
	}
}
