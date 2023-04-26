package bigquery_product

type ProductSiigo struct {
	Code             string `json:"product_code"`
	Name             string `json:"product_name"`
	AccountGroupID   int64  `json:"product_account_group_id"`
	AccountGroupName string `json:"product_account_group_name"`
}
