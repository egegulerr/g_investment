package app

import (
	"fmt"
	"g_investment/internal/domain"
	"g_investment/internal/domain/dtos"
	"g_investment/internal/ports"
)

type NewsService struct {
	provider ports.NewsProvider
}

func NewNewsService(provider ports.NewsProvider) *NewsService {
	return &NewsService{provider: provider}
}

func (s *NewsService) GetAllNews() ([]domain.News, error) {
	news, err := s.provider.GetAllNewsFromDB()
	if err != nil {
		return nil, fmt.Errorf("news service: failed to fetch news from db: %w", err)
	}

	return news, nil
}

func (s *NewsService) GetAllNewsGroupedByStock() ([]dtos.StockWithNewsDTO, error) {
	stocks, err := s.provider.GetAllNewsGroupedByStock()
	if err != nil {
		return nil, fmt.Errorf("news service: failed to fetch news from db: %w", err)
	}
	return stocks, nil
}

func (s *NewsService) FetchAndSaveNews() error {
	newsDTO, err := s.provider.FetchNewsFromAPI()

	if err != nil {
		return fmt.Errorf("news service: failed to fetch news: %w", err)
	}
	err = s.provider.SaveNewsToDB(newsDTO)
	if err != nil {
		return fmt.Errorf("news service: failed to save news to db: %w", err)

	}
	return nil
}
