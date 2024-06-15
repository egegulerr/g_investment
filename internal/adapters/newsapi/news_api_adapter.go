package newsapi

import (
	"encoding/json"
	"fmt"
	"g_investment/internal/domain"
	"g_investment/internal/ports"
	"io"
	"net/http"
	"os"
	"time"
)

type ApiConfig struct {
	BaseUrl string
	ApiKey  string
}

type NewsApiAdapter struct {
	api ApiConfig
}

type NewsResponseDTO struct {
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

func NewNewsApiAdapter(apiKey string) ports.NewsProvider {
	return &NewsApiAdapter{
		api: ApiConfig{
			BaseUrl: "https://www.alphavantage.co/query",
			ApiKey:  apiKey,
		},
	}
}

func (adapter *NewsApiAdapter) FetchNewsFromDB() ([]domain.News, error) {
	return locaJsonFile()
	/*
		 	url := fmt.Sprintf("%s?function=NEWS_SENTIMENT&limit=%s&apikey=%s", adapter.api.BaseUrl, "100", adapter.api.ApiKey)
			response, err := http.Get(url)
			if err != nil {
				return nil, fmt.Errorf("failed to fetch news: %w", err)
			}
			defer response.Body.Close()

			newsResponseDTO, err := convertResponseToDTO(response)
			if err != nil {
				return nil, fmt.Errorf("failed to convert response: %w", err)
			}

			newsItems := createNewsDomainObject(newsResponseDTO)
			return newsItems, nil
	*/
}

func locaJsonFile() ([]domain.News, error) {
	filepath := "./internal/adapters/newsapi/response.json"
	fileData, err := os.ReadFile(filepath)
	fmt.Println("fileData", fileData)

	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	newsResponseDTO, err := convertResponseToDTO2(fileData)
	if err != nil {
		return nil, fmt.Errorf("failed to convert response: %w", err)
	}
	newsItems := createNewsDomainObject(newsResponseDTO)
	return newsItems, nil

}

func authorsToString(authors []string) string {
	if len(authors) > 0 {
		return authors[0]
	} else {
		return "Unknown"
	}
}

func extractTickerList(tickerSentiments []TickerSentimentDTO) []domain.NewsStock {
	stockNewsList := make([]domain.NewsStock, 0)
	for _, tickerSentiment := range tickerSentiments {
		stockNewsList = append(stockNewsList, domain.NewsStock{
			Stock:               domain.Stock{Symbol: tickerSentiment.Ticker},
			RelevanceScore:      tickerSentiment.RelevanceScore,
			StockSentimentScore: tickerSentiment.TickerSentimentScore,
			StockSentimentLabel: tickerSentiment.TickerSentimentLabel,
		})
	}
	return stockNewsList
}

func convertResponseToDTO2(response []byte) (*NewsResponseDTO, error) {
	var responseJson NewsResponseDTO
	if err := json.Unmarshal(response, &responseJson); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &responseJson, nil
}

func convertResponseToDTO(response *http.Response) (*NewsResponseDTO, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)

	}
	var responseJson NewsResponseDTO
	if err := json.Unmarshal(body, &responseJson); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &responseJson, nil
}

func createNewsDomainObject(newsDTO *NewsResponseDTO) []domain.News {

	newsItems := make([]domain.News, 0)
	for _, item := range newsDTO.Feed {

		layout := "20060102T150405"

		parsedTime, err := time.Parse(layout, item.TimePublished)
		if err != nil {
			fmt.Printf("Error parsing time: %v\n", err)
			continue
		}

		news := domain.NewNews(
			item.URL,
			item.Title,
			authorsToString(item.Authors),
			item.OverallSentimentScore,
			item.Summary,
			item.BannerImage,
			parsedTime,
			extractTickerList(item.TickerSentiment),
		)
		newsItems = append(newsItems, *news)
	}
	return newsItems
}
