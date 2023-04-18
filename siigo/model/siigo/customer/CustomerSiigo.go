package siigo_customer

type CustomerSiigo struct {
	ID                     string         `json:"id"`
	Type                   string         `json:"type"`
	PersonType             string         `json:"person_type"`
	IDType                 IDType         `json:"id_type"`
	Identification         string         `json:"identification"`
	BranchOffice           int64          `json:"branch_office"`
	CheckDigit             *string        `json:"check_digit,omitempty"`
	Name                   *[]string      `json:"name"`
	Active                 bool           `json:"active"`
	VatResponsible         bool           `json:"vat_responsible"`
	FiscalResponsibilities []IDType       `json:"fiscal_responsibilities,omitempty"`
	Address                Address        `json:"address"`
	Phones                 []PhoneElement `json:"phones,omitempty"`
	Contacts               []Contact      `json:"contacts"`
	Metadata               Metadata       `json:"metadata"`
	CommercialName         *string        `json:"commercial_name,omitempty"`
}

type Address struct {
	Address string `json:"address"`
}

type Contact struct {
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Email     string        `json:"email"`
	Phone     *ContactPhone `json:"phone,omitempty"`
}

type ContactPhone struct {
	Indicative *string `json:"indicative,omitempty"`
	Number     *string `json:"number,omitempty"`
}

type IDType struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Metadata struct {
	Created     string  `json:"created"`
	LastUpdated *string `json:"last_updated,omitempty"`
}

type PhoneElement struct {
	Indicative *string `json:"indicative,omitempty"`
	Number     string  `json:"number"`
	Extension  *string `json:"extension,omitempty"`
}
