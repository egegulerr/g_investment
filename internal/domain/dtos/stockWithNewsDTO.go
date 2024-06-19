package dtos

import "time"

type StockWithNewsDTO struct {
	Symbol string          `json:"symbol"`
	News   []NewsSimpleDTO `json:"news"`
}

type NewsSimpleDTO struct {
	ID                       uint      `json:"id"`
	URL                      string    `json:"url"`
	Title                    string    `json:"title"`
	Authors                  []string  `json:"authors"`
	SentimentalAnalysisScore float64   `json:"sentimental_analysis_score"`
	Date                     time.Time `json:"date"`
	Summary                  string    `json:"summary"`
	Image                    string    `json:"image"`
	RelevanceScore           float64   `json:"relevance_score"`
	StockSentimentScore      float64   `json:"stock_sentiment_score"`
	StockSentimentLabel      string    `json:"stock_sentiment_label"`
}
