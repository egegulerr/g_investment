package newsapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"g_investment/internal/domain"
	"g_investment/internal/domain/dtos"
	"g_investment/internal/ports"
	"io"
	"log"
	"net/http"
	"os"

	"gorm.io/gorm"
)

type ApiConfig struct {
	BaseUrl string
	ApiKey  string
}

type NewsApiAdapter struct {
	api        ApiConfig
	repository *gorm.DB
}

func NewNewsApiAdapter(apiKey string, db *gorm.DB) ports.NewsProvider {
	return &NewsApiAdapter{
		api: ApiConfig{
			BaseUrl: "https://www.alphavantage.co/query",
			ApiKey:  apiKey,
		},
		repository: db,
	}
}

func (adapter *NewsApiAdapter) FetchNews() (*dtos.NewsResponseDTO, error) {
	return locaJsonFile()
}

func (adapter *NewsApiAdapter) SaveNews(newsSlice []domain.News) error {
	for j, news := range newsSlice[:3] {
		if adapter.CheckIfNewsExists(&news) {
			log.Printf("News already exists: %s", news.Url)
			continue
		}

		for i := range news.NewsStocks {
			ns := &news.NewsStocks[i]

			existingStock := adapter.GetStockIfAlreadyExists(ns)
			if existingStock != nil {
				ns.Stock.Model.ID = existingStock.Model.ID
				ns.StockID = existingStock.Model.ID
			} else {
				adapter.CreateStock(ns)
			}
		}

		if err := adapter.repository.Save(&news).Error; err != nil {
			log.Printf("Index: %d  Error: %s", j, err)
			continue
		}
	}
	return nil
}

func (adapter *NewsApiAdapter) CreateStock(newsStock *domain.NewsStock) {
	if err := adapter.repository.Create(newsStock.Stock).Error; err != nil {
		log.Printf("Error creating stock: %s", err)
	}
}

func (adapter *NewsApiAdapter) GetStockIfAlreadyExists(newsStock *domain.NewsStock) *domain.Stock {
	var existingStock domain.Stock
	result := adapter.repository.Where("symbol = ?", newsStock.Stock.Symbol).First(&existingStock)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &existingStock
	}
	return nil
}

func (adapter *NewsApiAdapter) CheckIfNewsExists(news *domain.News) bool {
	var existingNews domain.News
	result := adapter.repository.Where("url = ?", news.Url).First(&existingNews)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func (adapter *NewsApiAdapter) GetNews() (*dtos.NewsResponseDTO, error) {
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

func locaJsonFile() (*dtos.NewsResponseDTO, error) {
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
	return newsResponseDTO, nil

}

func convertResponseToDTO2(response []byte) (*dtos.NewsResponseDTO, error) {
	var responseJson dtos.NewsResponseDTO
	if err := json.Unmarshal(response, &responseJson); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &responseJson, nil
}

func convertResponseToDTO(response *http.Response) (*dtos.NewsResponseDTO, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)

	}
	var responseJson dtos.NewsResponseDTO
	if err := json.Unmarshal(body, &responseJson); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &responseJson, nil
}
