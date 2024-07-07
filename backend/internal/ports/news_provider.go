package ports

import (
	"g_investment/internal/domain"
	"g_investment/internal/domain/dtos"
)

type NewsProvider interface {
	GetAllNewsFromDB() ([]domain.News, error)
	GetAllNewsGroupedByStock() ([]dtos.StockWithNewsDTO, error)
	FetchNewsFromAPI() (*dtos.NewsApiResponseDTO, error)
	SaveNewsToDB(*dtos.NewsApiResponseDTO) error
}
