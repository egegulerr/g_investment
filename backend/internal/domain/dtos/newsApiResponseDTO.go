package dtos

type NewsApiResponseDTO struct {
	Items                    string `json:"items"`
	SentimentScoreDefinition string `json:"sentiment_score_definition"`
	RelevanceScoreDefinition string `json:"relevance_score_definition"`
	Feed                     []struct {
		Title                string   `json:"title"`
		URL                  string   `json:"url"`
		TimePublished        string   `json:"time_published"`
		Authors              []string `json:"authors"`
		Summary              string   `json:"summary"`
		BannerImage          string   `json:"banner_image"`
		Source               string   `json:"source"`
		CategoryWithinSource string   `json:"category_within_source"`
		SourceDomain         string   `json:"source_domain"`
		Topics               []struct {
			Topic          string `json:"topic"`
			RelevanceScore string `json:"relevance_score"`
		} `json:"topics"`
		OverallSentimentScore float64              `json:"overall_sentiment_score"`
		OverallSentimentLabel string               `json:"overall_sentiment_label"`
		TickerSentiment       []TickerSentimentDTO `json:"ticker_sentiment"`
	} `json:"feed"`
}

type TickerSentimentDTO struct {
	Ticker               string `json:"ticker"`
	RelevanceScore       string `json:"relevance_score"`
	TickerSentimentScore string `json:"ticker_sentiment_score"`
	TickerSentimentLabel string `json:"ticker_sentiment_label"`
}
