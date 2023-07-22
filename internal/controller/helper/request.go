package helper

// GetIncomeTaxParams is  query params for getting salary and year to calculate tax
type GetIncomeTaxParams struct {
	Salary string `form:"salary" binding:"required,numeric"`
	Year   string `form:"year" binding:"required,numeric,len=4"`
}

// IsValidTaxYear checks if the provided tax year is valid.
// It returns true if the tax year is valid, otherwise false.
func IsValidTaxYear(taxYear string) bool {
	validTaxYears := []string{"2019", "2020", "2021", "2022"}
	for _, year := range validTaxYears {
		if year == taxYear {
			return true
		}
	}
	return false
}
