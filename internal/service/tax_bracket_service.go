package service

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/siparisa/interview-test-server/internal/entity"
	"net/http"
	"time"
)

// ITaxBracketService defines the interface for bracket-related calculations.
type ITaxBracketService interface {
	GetTaxBracket(taxYear string, maxRetries int, retryInterval time.Duration) (*entity.TaxBrackets, error)
}

type taxBracketService struct {
	taxCalculatorURL string
}

// NewTaxBracketService creates a new instance of the taxBracketService.
func NewTaxBracketService(taxCalculatorURL string) ITaxBracketService {
	return &taxBracketService{
		taxCalculatorURL: taxCalculatorURL,
	}
}

// GetTaxBracket retrieves the tax response for the given year from the tax calculator API.
func (s *taxBracketService) GetTaxBracket(taxYear string, maxRetries int, retryInterval time.Duration) (*entity.TaxBrackets, error) {
	url := fmt.Sprintf(s.taxCalculatorURL + taxYear)

	for retry := 0; retry <= maxRetries; retry++ {
		if retry > 0 {
			// Wait for the specified retry interval before the next attempt
			time.Sleep(retryInterval)
		}

		resp, err := http.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			continue
		}

		var taxBrackets entity.TaxBrackets
		err = json.NewDecoder(resp.Body).Decode(&taxBrackets)
		if err != nil {
			continue
		}

		for i := range taxBrackets.TaxBrackets {
			bandName := fmt.Sprintf("band%d", i+1)
			taxBrackets.TaxBrackets[i].Band = bandName
		}

		return &taxBrackets, nil
	}

	return nil, fmt.Errorf("failed to get tax bracket after %d retries", maxRetries)
}
