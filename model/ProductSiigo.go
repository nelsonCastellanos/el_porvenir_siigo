package model

import "time"

type AccountGroup struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type AdditionalFields struct {
	Barcode string `json:"barcode"`
}
type Metadata struct {
	Created     *time.Time `json:"created"`
	LastUpdated *time.Time `json:"last_updated"`
}

type Unit struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
type Warehouses struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type ProductSiigo struct {
	AccountGroup        AccountGroup      `json:"account_group"`
	Active              bool              `json:"active"`
	AdditionalFields    *AdditionalFields `json:"additional_fields,omitempty"`
	AvailableQuantity   int               `json:"available_quantity"`
	Code                string            `json:"code"`
	Description         string            `json:"description,omitempty"`
	ID                  string            `json:"id"`
	Metadata            *Metadata         `json:"metadata,omitempty"`
	Name                string            `json:"name"`
	Reference           string            `json:"reference,omitempty"`
	StockControl        bool              `json:"stock_control"`
	TaxClassification   string            `json:"tax_classification"`
	TaxIncluded         bool              `json:"tax_included"`
	Taxes               []Taxes           `json:"taxes,omitempty"`
	Type                string            `json:"type"`
	Unit                Unit              `json:"unit,omitempty"`
	Warehouses          []Warehouses      `json:"warehouses"`
	TaxConsumptionValue int               `json:"tax_consumption_value,omitempty"`
}
