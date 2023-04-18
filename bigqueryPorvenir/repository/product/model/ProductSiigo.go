package bigquery_product

import bigquery_customer "el_porvenir.com/cloudfunction/bigqueryPorvenir/repository/customer/model"

type ProductSiigo struct {
	ID                  string                              `json:"id" gorm:"primaryKey;size:256"`
	Code                string                              `json:"code"`
	Name                string                              `json:"name"`
	AccountGroup        bigquery_customer.AccountGroupSiigo `json:"account_group" gorm:"foreignKey:IDAccountGroup;references:ID"`
	Type                string                              `json:"type"`
	StockControl        bool                                `json:"stock_control"`
	Active              bool                                `json:"active"`
	TaxClassification   string                              `json:"tax_classification"`
	TaxIncluded         bool                                `json:"tax_included"`
	TaxConsumptionValue int64                               `json:"tax_consumption_value,omitempty"`
	Unit                Unit                                `json:"unit" gorm:"embedded;embeddedPrefix:unit_"`
	UnitLabel           string                              `json:"unit_label,omitempty"`
	Reference           string                              `json:"reference"`
	Description         string                              `json:"description"`
	AdditionalFields    AdditionalFields                    `json:"additional_fields"  gorm:"embedded;embeddedPrefix:additional_fields_"`
	AvailableQuantity   float64                             `json:"available_quantity"`
	Metadata            Metadata                            `json:"metadata" gorm:"embedded;embeddedPrefix:meta_data_"`
}

type AdditionalFields struct {
	Barcode string `json:"barcode"`
}

type Metadata struct {
	Created     string `json:"created"`
	LastUpdated string `json:"last_updated,omitempty"`
}

type Tax struct {
	IDProduct  string `gorm:"size:256" bigquery:"-"`
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Percentage int64  `json:"percentage"`
}

type Unit struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Warehouse struct {
	ID        int64   `json:"id"  gorm:"primaryKey"`
	IDProduct string  `gorm:"size:256"`
	Name      string  `json:"name"`
	Quantity  float64 `json:"quantity"`
}
