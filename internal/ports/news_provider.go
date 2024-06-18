package ports

import (
	"g_investment/internal/domain/dtos"
)

type NewsProvider interface {
	GetNews() (*dtos.NewsResponseDTO, error)
	FetchNews() (*dtos.NewsResponseDTO, error)
	SaveNews(*dtos.NewsResponseDTO) error
}
