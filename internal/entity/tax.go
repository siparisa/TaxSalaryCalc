package entity

// TaxBracket represents a tax bracket with minimum and maximum values and a tax rate.
type TaxBracket struct {
	Max  float64 `json:"max"`
	Min  float64 `json:"min"`
	Rate float64 `json:"rate"`
}

// TaxBrackets represents the response containing an array of TaxBrackets.
type TaxBrackets struct {
	TaxBrackets []TaxBracket `json:"tax_brackets"`
}
