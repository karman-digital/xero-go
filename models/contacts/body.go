package contacts

type Contact struct {
	ContactID                 string    `json:"ContactID"`
	ContactStatus             string    `json:"ContactStatus"`
	Name                      string    `json:"Name"`
	FirstName                 string    `json:"FirstName"`
	LastName                  string    `json:"LastName"`
	CompanyNumber             string    `json:"CompanyNumber"`
	EmailAddress              string    `json:"EmailAddress"`
	BankAccountDetails        string    `json:"BankAccountDetails"`
	TaxNumber                 string    `json:"TaxNumber"`
	AccountsReceivableTaxType string    `json:"AccountsReceivableTaxType"`
	AccountsPayableTaxType    string    `json:"AccountsPayableTaxType"`
	Addresses                 []Address `json:"Addresses"`
	Phones                    []Phone   `json:"Phones"`
	UpdatedDateUTC            string    `json:"UpdatedDateUTC"`
	IsSupplier                bool      `json:"IsSupplier"`
	IsCustomer                bool      `json:"IsCustomer"`
	DefaultCurrency           string    `json:"DefaultCurrency"`
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
