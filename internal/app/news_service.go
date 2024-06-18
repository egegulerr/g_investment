package app

import (
	"fmt"
	"g_investment/internal/domain"
	"g_investment/internal/domain/dtos"
	"g_investment/internal/ports"
	"time"
)

type NewsService struct {
	provider ports.NewsProvider
}

func NewNewsService(provider ports.NewsProvider) *NewsService {
	return &NewsService{provider: provider}
}

func (s *NewsService) GetNewsFromDB() ([]domain.News, error) {
	newsDTO, err := s.provider.GetNews()
	if err != nil {
		return nil, fmt.Errorf("news service: failed to fetch news from db: %w", err)
	}
	news := createNewsDomainObject(newsDTO)

	return news, nil
}

func (s *NewsService) FetchAndSaveNews() error {
	newsDTO, err := s.provider.FetchNews()

	if err != nil {
		return fmt.Errorf("news service: failed to fetch news: %w", err)
	}
	err = s.provider.SaveNews(newsDTO)
	if err != nil {
		return fmt.Errorf("news service: failed to save news to db: %w", err)

	}
	return nil
}

func createNewsDomainObject(newsDTO *dtos.NewsResponseDTO) []domain.News {

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
			item.Authors,
			item.OverallSentimentScore,
			item.Summary,
			item.BannerImage,
			parsedTime,
		)
		newsItems = append(newsItems, *news)
	}
	return newsItems
}

func extractTickerList(tickerSentiments []dtos.TickerSentimentDTO) []domain.NewsStock {
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
