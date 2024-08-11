package adapters

import (
	"encoding/json"
	"fmt"
	"g_investment/internal/domain/dtos"
	logoports "g_investment/internal/ports/logoPorts"
	"io"
	"net/http"
)

type LogoApiConfig struct {
	BaseUrl string
	ApiKey  string
}

type LogoApiAdapter struct {
	api *LogoApiConfig
}

func NewsLogoApiAdapter(apiKey *string) logoports.LogoProvider {
	return &LogoApiAdapter{
		api: &LogoApiConfig{
			ApiKey:  *apiKey,
			BaseUrl: "https://api.api-ninjas.com/v1/logo?ticker="},
	}
}

func (adapter *LogoApiAdapter) FetchCompanyLogoFromAPI(ticker *string) ([]byte, error) {

	apiUrl := fmt.Sprintf("%s%s", adapter.api.BaseUrl, *ticker)

	response, err := adapter.fetchLogoLink(&apiUrl)
	if err != nil {
		return nil, fmt.Errorf("logo api adapter: failed to fetch logo: %w", err)
	}

	imageUrl := response[0].Image
	imgData, err := adapter.fetchLogo(&imageUrl)
	if err != nil {
		return nil, fmt.Errorf("logo api adapter: failed to fetch logo: %w", err)
	}

	return imgData, nil
}

func (adapter *LogoApiAdapter) fetchLogoLink(url *string) ([]dtos.LogoDTO, error) {
	req, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		return nil, fmt.Errorf("logo api adapter: failed to create request: %w", err)
	}
	req.Header.Set("X-Api-Key", adapter.api.ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("logo api adapter: failed to send request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("logo api adapter: failed to fetch logo: %s", resp.Status)
	}

	var result []dtos.LogoDTO
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("logo api adapter: failed to decode response: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("logo api adapter: logo not found")
	}

	return result, nil
}

func (s *LogoApiAdapter) fetchLogo(logoUrl *string) ([]byte, error) {
	resp, err := http.Get(*logoUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch logo: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch logo: %s", resp.Status)
	}

	imageData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read image: %w", err)
	}

	return imageData, nil
}
