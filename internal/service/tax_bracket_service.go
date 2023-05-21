package service

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/siparisa/interview-test-server/internal/config"
	"github.com/siparisa/interview-test-server/internal/entity"
	"net/http"
)

// GetTaxBracket retrieves the tax response for the given year from the tax calculator API.
func GetTaxBracket(taxYear string) (*entity.TaxBrackets, error) {

	port := config.GetPort("PORT_TAX_YEAR", config.TaxYearDefaultPort)

	taxCalculatorURL := fmt.Sprintf("http://localhost:%s/tax-calculator/tax-year/%s", port, taxYear)

	// Make request to tax calculator API
	resp, err := http.Get(taxCalculatorURL)
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

	return &taxBrackets, nil
}
