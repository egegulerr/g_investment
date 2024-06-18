package newsapi

import (
	"encoding/json"
	"fmt"
	"g_investment/internal/domain"
	"g_investment/internal/domain/dtos"
	"g_investment/internal/ports"
	"io"
	"net/http"
	"os"
	"time"

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

func (adapter *NewsApiAdapter) SaveNews(newsDTO *dtos.NewsResponseDTO) error {

	for _, feed := range newsDTO.Feed {
		parsedTime, err := parseTime(&feed.TimePublished)
		if err != nil {
			return err
		}
		news := domain.NewNews(
			feed.URL, feed.Title, feed.Authors, feed.OverallSentimentScore, feed.Summary, feed.BannerImage, *parsedTime)

		for _, tickerSentiment := range feed.TickerSentiment {
			var stock domain.Stock
			stock.Symbol = tickerSentiment.Ticker
			err = adapter.repository.Where("symbol = ?", stock.Symbol).FirstOrCreate(&stock).Error
			if err != nil {
				return err
			}

			news.NewsStocks = append(news.NewsStocks, domain.NewsStock{
				StockID:             stock.ID,
				RelevanceScore:      tickerSentiment.RelevanceScore,
				StockSentimentScore: tickerSentiment.TickerSentimentScore,
				StockSentimentLabel: tickerSentiment.TickerSentimentLabel,
			})
		}
		err = adapter.repository.Create(&news).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func parseTime(timeString *string) (*time.Time, error) {
	layout := "20060102T150405"
	parsedTime, err := time.Parse(layout, *timeString)
	if err != nil {
		fmt.Printf("Error parsing time: %v\n", err)
		return nil, err
	}
	return &parsedTime, nil
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
