package siigo_invoice

// https://api.siigo.com/v1/taxes
type TaxesSiigo struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	Percentage float64 `json:"percentage"`
	Active     bool    `json:"active"`
}
