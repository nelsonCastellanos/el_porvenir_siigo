package bigquery_invoice

// https://api.siigo.com/v1/payment-types?document_type=FV
type PaymentTypesSiigo struct {
	ID      int64  `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Active  bool   `json:"active"`
	DueDate bool   `json:"due_date"`
}
