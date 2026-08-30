package accountsmodels

type Accounts struct {
	ID           string    `json:"Id"`
	Status       string    `json:"Status"`
	ProviderName string    `json:"ProviderName"`
	DateTimeUTC  string    `json:"DateTimeUTC"`
	Accounts     []Account `json:"Accounts"`
}

type Account struct {
	AccountID      string `json:"AccountID"`
	Code           string `json:"Code"`
	Name           string `json:"Name"`
	Type           string `json:"Type"`
	Status         string `json:"Status"`
	UpdatedDateUTC string `json:"UpdatedDateUTC"`
}
