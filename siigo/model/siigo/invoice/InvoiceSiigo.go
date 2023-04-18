package siigo_invoice

type InvoiceSiigo struct {
	ID       string    `json:"id"`
	Document Document  `json:"document"`
	Prefix   string    `json:"prefix"`
	Number   int64     `json:"number"`
	Name     string    `json:"name"`
	Date     string    `json:"date"`
	Customer Customer  `json:"customer"`
	Seller   int64     `json:"seller"`
	Total    float64   `json:"total"`
	Balance  int64     `json:"balance"`
	Items    []Item    `json:"items"`
	Payments []Payment `json:"payments"`
	Stamp    Stamp     `json:"stamp"`
	Mail     Mail      `json:"mail"`
	Metadata Metadata  `json:"metadata"`
}

type Customer struct {
	ID             string `json:"id"`
	Identification string `json:"identification"`
	BranchOffice   int64  `json:"branch_office"`
}

type Document struct {
	ID int64 `json:"id"`
}

type Item struct {
	ID          string  `json:"id"`
	Code        string  `json:"code"`
	Quantity    float64 `json:"quantity"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Taxes       []Tax   `json:"taxes,omitempty"`
	Total       float64 `json:"total"`
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
	Created     string  `json:"created"`
	LastUpdated *string `json:"last_updated,omitempty"`
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
