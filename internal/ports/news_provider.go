package ports

import "g_investment/internal/domain"

type NewsProvider interface {
	FetchNewsFromDB() ([]domain.News, error)
}
