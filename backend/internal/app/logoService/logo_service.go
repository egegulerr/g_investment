package logoService

import (
	"fmt"
	"g_investment/internal/ports/logoPorts"
	"os"
	"path/filepath"
	"strings"
)

type LogoService struct {
	provider logoPorts.LogoProvider
}

func NewLogoService(provider logoPorts.LogoProvider) *LogoService {
	return &LogoService{provider: provider}
}

func (s *LogoService) GetCompanyLogoFromFiles(ticker *string) ([]byte, error) {
	fileName := fmt.Sprintf("%s.png", *ticker)
	filePath := filepath.Join("logos", fileName)

	imageData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	return imageData, nil
}

func (s *LogoService) SaveCompanyLogoToFiles(ticker *string, logo []byte) error {
	fileName := fmt.Sprintf("%s.png", *ticker)
	filePath := filepath.Join("logos", fileName)
	file, err := os.Create(filePath)

	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}
	defer file.Close()

	_, err = file.Write(logo)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	return nil
}

func (s *LogoService) CheckIfCompanyLogoExistsInFiles(ticker *string) bool {
	fileName := fmt.Sprintf("%s.png", *ticker)
	filePath := filepath.Join("logos", fileName)
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func (s *LogoService) FetchNewCompanyLogoFromAPI(ticker *string) ([]byte, error) {
	img, err := s.provider.FetchCompanyLogoFromAPI(ticker)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func (s *LogoService) GetCompanyLogo(tickerName *string) ([]byte, error) {
	ticker := strings.ToUpper(*tickerName)

	isExisting := s.CheckIfCompanyLogoExistsInFiles(&ticker)
	if isExisting {
		img, err := s.GetCompanyLogoFromFiles(&ticker)
		if err != nil {
			return nil, err
		}
		return img, nil
	}
	img, err := s.FetchNewCompanyLogoFromAPI(&ticker)
	if err != nil {
		return nil, err
	}
	s.SaveCompanyLogoToFiles(&ticker, img)
	return img, nil
}
