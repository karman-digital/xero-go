package contacts

type Contact struct {
	ContactID                 string          `json:"ContactID,omitempty"`
	ContactStatus             string          `json:"ContactStatus,omitempty"`
	ContactNumber             string          `json:"ContactNumber,omitempty"`
	AccountNumber             string          `json:"AccountNumber,omitempty"`
	Name                      string          `json:"Name"`
	FirstName                 string          `json:"FirstName,omitempty"`
	LastName                  string          `json:"LastName,omitempty"`
	CompanyNumber             string          `json:"CompanyNumber,omitempty"`
	EmailAddress              string          `json:"EmailAddress,omitempty"`
	BankAccountDetails        string          `json:"BankAccountDetails,omitempty"`
	TaxNumber                 string          `json:"TaxNumber,omitempty"`
	AccountsReceivableTaxType string          `json:"AccountsReceivableTaxType,omitempty"`
	AccountsPayableTaxType    string          `json:"AccountsPayableTaxType,omitempty"`
	Addresses                 []Address       `json:"Addresses,omitempty"`
	Phones                    []Phone         `json:"Phones,omitempty"`
	ContactPersons            []ContactPerson `json:"ContactPersons,omitempty"`
	UpdatedDateUTC            string          `json:"UpdatedDateUTC,omitempty"`
	IsSupplier                bool            `json:"IsSupplier,omitempty"`
	IsCustomer                bool            `json:"IsCustomer,omitempty"`
	DefaultCurrency           string          `json:"DefaultCurrency,omitempty"`
}

type Address struct {
	AddressType  string `json:"AddressType"`
	AddressLine1 string `json:"AddressLine1,omitempty"`
	City         string `json:"City,omitempty"`
	PostalCode   string `json:"PostalCode,omitempty"`
	AttentionTo  string `json:"AttentionTo,omitempty"`
}

type Phone struct {
	PhoneType        string `json:"PhoneType"`
	PhoneNumber      string `json:"PhoneNumber,omitempty"`
	PhoneAreaCode    string `json:"PhoneAreaCode,omitempty"`
	PhoneCountryCode string `json:"PhoneCountryCode,omitempty"`
}

type Contacts struct {
	Contacts []Contact `json:"Contacts"`
}

type ContactPerson struct {
	FirstName       string `json:"FirstName"`
	LastName        string `json:"LastName"`
	EmailAddress    string `json:"EmailAddress"`
	IncludeInEmails bool   `json:"IncludeInEmails"`
}
