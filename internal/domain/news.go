package domain

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	Url                      string         `gorm:"unique" json:"url"`
	Title                    string         `json:"title"`
	Authors                  pq.StringArray `gorm:"type:text[]" json:"authors"`
	SentimentalAnalysisScore float64        `json:"sentimentalAnalysisScore"`
	Date                     time.Time      `json:"date"`
	Summary                  string         `json:"summary"`
	Image                    string         `json:"image"`
	NewsStocks               []NewsStock    `gorm:"foreignKey:NewsID"`
}

type NewsStock struct {
	gorm.Model
	NewsID  uint `gorm:"index"`
	StockID uint `gorm:"index"`

	Stock Stock `gorm:"foreignKey:StockID"`

	RelevanceScore      float64 `json:"relevance_score"`
	StockSentimentScore float64 `json:"stock_sentiment_score"`
	StockSentimentLabel string  `json:"stock_sentiment_label"`
}

func NewNews(url string, title string, author []string, sentimentalAnalysisScore float64, summary string, image string, date time.Time) *News {
	return &News{
		Url:                      url,
		Title:                    title,
		Authors:                  author,
		SentimentalAnalysisScore: sentimentalAnalysisScore,
		Date:                     date,
		Summary:                  summary,
		Image:                    image,
	}
}
