package model

type InvoiceSiigo struct {
	ID       string     `json:"id"`
	Document Document   `json:"document"`
	Prefix   string     `json:"prefix"`
	Number   int        `json:"number"`
	Name     string     `json:"name"`
	Date     string     `json:"date"`
	Customer Customer   `json:"customer"`
	Seller   int        `json:"seller"`
	Total    float64    `json:"total"`
	Balance  float64    `json:"balance"`
	Items    []Items    `json:"items"`
	Payments []Payments `json:"payments"`
	Stamp    Stamp      `json:"stamp"`
	Mail     Mail       `json:"mail"`
	Metadata *Metadata  `json:"metadata"`
}
type Document struct {
	ID int `json:"id"`
}
type Customer struct {
	ID             string `json:"id"`
	Identification string `json:"identification"`
	BranchOffice   int    `json:"branch_office"`
}
type Taxes struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	Percentage float64 `json:"percentage"`
	Value      float64 `json:"value"`
}
type Items struct {
	ID          string  `json:"id"`
	Code        string  `json:"code"`
	Quantity    float64 `json:"quantity"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Taxes       []Taxes `json:"taxes"`
	Total       float64 `json:"total"`
}
type Payments struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}
type Stamp struct {
	Status string `json:"status"`
	Cufe   string `json:"cufe"`
}
type Mail struct {
	Status       string `json:"status"`
	Observations string `json:"observations"`
}
