package bigquery_invoice

import (
	bigquery_product "el_porvenir.com/cloudfunction/bigqueryPorvenir/repository/product/model"
)

type InvoiceSiigo struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	Date               string    `json:"date"`
	CustomerName       string    `json:"customer_name"`
	CustomerIdTypeCode string    `json:"customer_id_type_code"`
	CustomerIdTypeName string    `json:"customer_id_type_name"`
	Total              float64   `json:"total"`
	Items              []Item    `json:"items"`
	Payments           []Payment `json:"payments"`
	StampStatus        string    `json:"stamp_status"`
	StampCufe          string    `json:"stamp_cufe"`
	MailStatus         string    `json:"mail_status"`
	MailObservations   string    `json:"mail_observations"`
	Created            string    `json:"created"`
	LastUpdated        string    `json:"last_updated,omitempty"`
}

type Item struct {
	Product  bigquery_product.ProductSiigo `json:"product"`
	Quantity float64                       `json:"quantity"`
	Price    float64                       `json:"price"`
	Total    float64                       `json:"total"`
}

type Payment struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}
