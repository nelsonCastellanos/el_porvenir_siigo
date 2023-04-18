package bigquery_invoice

// https://api.siigo.com/v1/document-types?type=FV
type DocumentTypesSiigo struct {
	ID                   int64  `json:"id" gorm:"primaryKey"`
	Code                 string `json:"code"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	Type                 string `json:"type"`
	Active               bool   `json:"active"`
	SellerByItem         bool   `json:"seller_by_item"`
	CostCenter           bool   `json:"cost_center"`
	CostCenterMandatory  bool   `json:"cost_center_mandatory"`
	AutomaticNumber      bool   `json:"automatic_number"`
	Consecutive          int64  `json:"consecutive"`
	DiscountType         string `json:"discount_type"`
	Decimals             bool   `json:"decimals"`
	AdvancePayment       bool   `json:"advance_payment"`
	Reteiva              bool   `json:"reteiva"`
	Reteica              bool   `json:"reteica"`
	SelfWithholding      bool   `json:"self_withholding"`
	SelfWithholdingLimit int64  `json:"self_withholding_limit"`
	ElectronicType       string `json:"electronic_type"`
}
