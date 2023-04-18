package siigo_product

type ProductSiigo struct {
	ID                  string           `json:"id"`
	Code                string           `json:"code"`
	Name                string           `json:"name"`
	AccountGroup        AccountGroup     `json:"account_group"`
	Type                string           `json:"type"`
	StockControl        bool             `json:"stock_control"`
	Active              bool             `json:"active"`
	TaxClassification   string           `json:"tax_classification"`
	TaxIncluded         bool             `json:"tax_included"`
	TaxConsumptionValue *int64           `json:"tax_consumption_value,omitempty"`
	Unit                Unit             `json:"unit"`
	UnitLabel           *string          `json:"unit_label,omitempty"`
	Reference           string           `json:"reference"`
	Description         string           `json:"description"`
	AdditionalFields    AdditionalFields `json:"additional_fields"`
	AvailableQuantity   float64          `json:"available_quantity"`
	Warehouses          []Warehouse      `json:"warehouses"`
	Metadata            Metadata         `json:"metadata"`
	Taxes               []Tax            `json:"taxes,omitempty"`
}

type AccountGroup struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type AdditionalFields struct {
	Barcode string `json:"barcode"`
}

type Metadata struct {
	Created     string  `json:"created"`
	LastUpdated *string `json:"last_updated,omitempty"`
}

type Tax struct {
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
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
}
