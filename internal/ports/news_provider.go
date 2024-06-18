package ports

import (
	"g_investment/internal/domain"
	"g_investment/internal/domain/dtos"
)

type NewsProvider interface {
	GetNews() (*dtos.NewsResponseDTO, error)
	FetchNews() (*dtos.NewsResponseDTO, error)
	SaveNews([]domain.News) error
}
