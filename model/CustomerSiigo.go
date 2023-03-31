package model

type CustomerSiigo struct {
	Active                 bool                     `json:"active"`
	Address                Address                  `json:"address"`
	BranchOffice           int                      `json:"branch_office"`
	CheckDigit             string                   `json:"check_digit,omitempty"`
	Contacts               []Contacts               `json:"contacts"`
	FiscalResponsibilities []FiscalResponsibilities `json:"fiscal_responsibilities,omitempty"`
	ID                     string                   `json:"id"`
	IDType                 IDType                   `json:"id_type"`
	Identification         string                   `json:"identification"`
	Metadata               *Metadata                `json:"metadata,omitempty"`
	Name                   []string                 `json:"name"`
	PersonType             string                   `json:"person_type"`
	Phones                 []Phones                 `json:"phones,omitempty"`
	Type                   string                   `json:"type"`
	VatResponsible         bool                     `json:"vat_responsible"`
	CommercialName         string                   `json:"commercial_name,omitempty"`
}

type Address struct {
	Address string `json:"address"`
}

type Contacts struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     Phones `json:"phone"`
}

type FiscalResponsibilities struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type IDType struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Phones struct {
	Indicative string `json:"indicative"`
	Number     string `json:"number"`
}
