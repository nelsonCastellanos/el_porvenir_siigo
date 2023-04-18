package bigquery_customer

type CustomerSiigo struct {
	ID                     string                 `json:"id" gorm:"primaryKey;size:256"`
	Type                   string                 `json:"type"`
	PersonType             string                 `json:"person_type"`
	IDType                 IDType                 `json:"id_type" gorm:"embedded;embeddedPrefix:id_type_"`
	Identification         string                 `json:"identification"`
	BranchOffice           int64                  `json:"branch_office"`
	CheckDigit             string                 `json:"check_digit,omitempty"`
	Name                   string                 `json:"name"`
	Active                 bool                   `json:"active"`
	VatResponsible         bool                   `json:"vat_responsible"`
	FiscalResponsibilities []FiscalResponsibility `json:"fiscal_responsibilities,omitempty" gorm:"foreignKey:IDCustomer"`
	Address                Address                `json:"address"  gorm:"embedded;embeddedPrefix:adress_"`
	Phones                 []PhoneElement         `json:"phones,omitempty" gorm:"foreignKey:IDCustomer"`
	Contacts               []Contact              `json:"contacts" gorm:"foreignKey:IDCustomer"`
	Metadata               Metadata               `json:"metadata" gorm:"embedded;embeddedPrefix:metadata_"`
	CommercialName         string                 `json:"commercial_name,omitempty"`
}

type FiscalResponsibility struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Address struct {
	Address string `json:"address"`
}

type Contact struct {
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	Email     string       `json:"email"`
	Phone     ContactPhone `json:"phone,omitempty" gorm:"embedded;embeddedPrefix:phone_"`
}

type ContactPhone struct {
	Indicative string `json:"indicative,omitempty"`
	Number     string `json:"number,omitempty"`
}

type IDType struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Metadata struct {
	Created     string `json:"created"`
	LastUpdated string `json:"last_updated,omitempty"`
}

type PhoneElement struct {
	Indicative string `json:"indicative,omitempty"`
	Number     string `json:"number"`
	Extension  string `json:"extension,omitempty"`
}
