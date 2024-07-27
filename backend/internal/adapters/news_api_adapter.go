package adapters

import (
	"encoding/json"
	"fmt"
	"g_investment/internal/domain"
	"g_investment/internal/domain/dtos"
	"g_investment/internal/ports"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
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

func (adapter *NewsApiAdapter) FetchNewsFromAPI() (*dtos.NewsApiResponseDTO, error) {
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

func (adapter *NewsApiAdapter) SaveNewsToDB(newsDTO *dtos.NewsApiResponseDTO) error {

	for _, feed := range newsDTO.Feed {
		parsedTime, err := parseTime(&feed.TimePublished)
		if err != nil {
			return err
		}
		news := domain.NewNews(
			feed.URL, feed.Title, feed.Authors, feed.OverallSentimentScore, feed.Summary, feed.BannerImage, *parsedTime)

		err = adapter.repository.Where("url = ?", news.Url).FirstOrCreate(&news).Error
		if err != nil {
			log.Printf("Found duplicate news. Skipping it: %v\n", err)
			continue
		}

		for _, tickerSentiment := range feed.TickerSentiment {
			var stock domain.Stock
			stock.Symbol = tickerSentiment.Ticker
			err = adapter.repository.Where("symbol = ?", stock.Symbol).FirstOrCreate(&stock).Error
			if err != nil {
				return err
			}

			news.NewsStocks = append(news.NewsStocks, domain.NewsStock{
				StockID:             stock.ID,
				RelevanceScore:      parseScore(&tickerSentiment.RelevanceScore),
				StockSentimentScore: parseScore(&tickerSentiment.TickerSentimentScore),
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

func parseScore(scoreString *string) float64 {
	score, err := strconv.ParseFloat(*scoreString, 64)
	if err != nil {
		fmt.Printf("Error parsing score: %v\n", err)
		return 0.0
	}
	return score
}

func (adapter *NewsApiAdapter) GetAllNewsFromDB() ([]domain.News, error) {
	var newsList []domain.News
	err := adapter.repository.Preload("NewsStocks.Stock").Find(&newsList).Error
	if err != nil {
		return nil, fmt.Errorf("failed to fetch news from db: %w", err)
	}
	return newsList, nil
}

func (adapter *NewsApiAdapter) GetAllNewsGroupedByStock() ([]dtos.StockWithNewsDTO, error) {
	var newsStocks []dtos.NewsStockWithDetails

	err := adapter.repository.
		Table("news_stocks").
		Select(`news_stocks.*, 
			news.id as news_id, news.url, news.title, news.authors, news.sentimental_analysis_score, 
			news.date, news.summary, news.image, 
			stocks.symbol as stock_symbol`).
		Joins("left join news on news_stocks.news_id = news.id").
		Joins("left join stocks on news_stocks.stock_id = stocks.id").
		Find(&newsStocks).Error
	if err != nil {
		return nil, err
	}

	stockMap := make(map[string][]dtos.NewsSimpleDTO)
	for _, ns := range newsStocks {
		stockMap[ns.StockSymbol] = append(stockMap[ns.StockSymbol], dtos.NewsSimpleDTO{
			ID:                       ns.NewsID,
			URL:                      ns.URL,
			Title:                    ns.Title,
			Authors:                  ns.Authors,
			SentimentalAnalysisScore: ns.SentimentalAnalysisScore,
			Date:                     ns.Date,
			Summary:                  ns.Summary,
			Image:                    ns.Image,
			RelevanceScore:           ns.RelevanceScore,
			StockSentimentScore:      ns.StockSentimentScore,
			StockSentimentLabel:      ns.StockSentimentLabel,
		})
	}

	var result []dtos.StockWithNewsDTO
	for symbol, newsList := range stockMap {
		result = append(result, dtos.StockWithNewsDTO{
			Symbol: symbol,
			News:   newsList,
		})
	}

	return result, nil
}

func locaJsonFile() (*dtos.NewsApiResponseDTO, error) {
	filepath := "./internal/adapters/newsapi/response.json"
	fileData, err := os.ReadFile(filepath)
	fmt.Println("fileData", fileData)

	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	newsResponseDTO, err := convertLocalJsonToDTO(fileData)
	if err != nil {
		return nil, fmt.Errorf("failed to convert response: %w", err)
	}
	return newsResponseDTO, nil

}

func convertLocalJsonToDTO(response []byte) (*dtos.NewsApiResponseDTO, error) {
	var responseJson dtos.NewsApiResponseDTO
	if err := json.Unmarshal(response, &responseJson); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &responseJson, nil
}

func convertResponseToDTO(response *http.Response) (*dtos.NewsApiResponseDTO, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)

	}
	var responseJson dtos.NewsApiResponseDTO
	if err := json.Unmarshal(body, &responseJson); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &responseJson, nil
}
