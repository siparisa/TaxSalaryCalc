package service

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/siparisa/interview-test-server/internal/entity"
	"net/http"
)

type ITaxBracketService interface {
	GetTaxBracket(taxYear string) (*entity.TaxBrackets, error)
}

type taxBracketService struct {
	taxCalculatorURL string
}

func NewTaxBracketService(taxCalculatorURL string) ITaxBracketService {
	return &taxBracketService{
		taxCalculatorURL: taxCalculatorURL,
	}
}

// GetTaxBracket retrieves the tax response for the given year from the tax calculator API.
func (s *taxBracketService) GetTaxBracket(taxYear string) (*entity.TaxBrackets, error) {

	url := fmt.Sprintf(s.taxCalculatorURL + taxYear)

	// Make request to tax calculator API
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request to tax calculator API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get tax bracket, received non-OK status code: %d", resp.StatusCode)
	}

	var taxBrackets entity.TaxBrackets
	err = json.NewDecoder(resp.Body).Decode(&taxBrackets)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %v", err)
	}

	// Add the Band values dynamically
	for i := range taxBrackets.TaxBrackets {
		bandName := fmt.Sprintf("band%d", i+1)
		taxBrackets.TaxBrackets[i].Band = bandName
	}

	return &taxBrackets, nil
}
