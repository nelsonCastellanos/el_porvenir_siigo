package bigquery_customer

type AccountGroupSiigo struct {
	ID     int64  `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}
