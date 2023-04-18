package bigquery_invoice

import (
	bigquery_customer "el_porvenir.com/cloudfunction/bigqueryPorvenir/repository/customer/model"
	bigquery_product "el_porvenir.com/cloudfunction/bigqueryPorvenir/repository/product/model"
)

type InvoiceSiigo struct {
	ID       string                          `json:"id"`
	Document DocumentTypesSiigo              `json:"document"`
	Prefix   string                          `json:"prefix"`
	Number   int64                           `json:"number"`
	Name     string                          `json:"name"`
	Date     string                          `json:"date"`
	Customer bigquery_customer.CustomerSiigo `json:"customer"`
	Seller   int64                           `json:"seller"`
	Total    float64                         `json:"total"`
	Balance  int64                           `json:"balance"`
	Items    []Item                          `json:"items"`
	Payments []Payment                       `json:"payments"`
	Stamp    Stamp                           `json:"stamp"`
	Mail     Mail                            `json:"mail"`
	Metadata Metadata                        `json:"metadata"`
}

type Document struct {
	ID int64 `json:"id"`
}

type Item struct {
	Product     bigquery_product.ProductSiigo `json:"product"`
	Quantity    float64                       `json:"quantity"`
	Price       float64                       `json:"price"`
	Description string                        `json:"description"`
	Taxes       []Tax                         `json:"taxes,omitempty"`
	Total       float64                       `json:"total"`
}

type Tax struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	Percentage float64 `json:"percentage"`
	Value      float64 `json:"value"`
}

type Mail struct {
	Status       string `json:"status"`
	Observations string `json:"observations"`
}

type Metadata struct {
	Created     string `json:"created"`
	LastUpdated string `json:"last_updated,omitempty"`
}

type Payment struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type Stamp struct {
	Status string `json:"status"`
	Cufe   string `json:"cufe"`
}
