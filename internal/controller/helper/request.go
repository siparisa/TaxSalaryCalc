package helper

// GetIncomeTaxParams is  query params for getting salary and year to calculate tax
type GetIncomeTaxParams struct {
	Salary string `form:"salary" binding:"required,numeric"`
	Year   string `form:"year" binding:"required,numeric,len=4"`
}
