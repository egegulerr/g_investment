package app

import (
	"g_investment/internal/domain"
	"g_investment/internal/ports"
)

type NewsService struct {
	provider ports.NewsProvider
}

func NewNewsService(provider ports.NewsProvider) *NewsService {
	return &NewsService{provider: provider}
}

func (s *NewsService) GetNews() ([]domain.News, error) {
	return s.provider.FetchNewsFromDB()
	/*
	   TODO
	   dto := mapToDomainToDto(news)
	*/
}
