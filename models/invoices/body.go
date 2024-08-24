package invoicesmodels

import sharedmodels "github.com/karman-digital/xero-go/models/shared"

type Invoice struct {
	Type            string                      `json:"Type"`
	Contact         Contact                     `json:"Contact"`
	LineItems       []LineItem                  `json:"LineItems"`
	Date            *string                     `json:"Date,omitempty"`
	DueDate         *string                     `json:"DueDate,omitempty"`
	DateString      *string                     `json:"DateString,omitempty"`
	DueDateString   *string                     `json:"DueDateString,omitempty"`
	Status          *string                     `json:"Status,omitempty"`
	LineAmountTypes *string                     `json:"LineAmountTypes,omitempty"`
	SubTotal        *sharedmodels.FloatOrString `json:"SubTotal,omitempty"`
	TotalTax        *sharedmodels.FloatOrString `json:"TotalTax,omitempty"`
	Total           *sharedmodels.FloatOrString `json:"Total,omitempty"`
	UpdatedDateUTC  *string                     `json:"UpdatedDateUTC,omitempty"`
	CurrencyCode    *string                     `json:"CurrencyCode,omitempty"`
	InvoiceID       *string                     `json:"InvoiceID,omitempty"`
	InvoiceNumber   *string                     `json:"InvoiceNumber,omitempty"`
	Payments        []Payment                   `json:"Payments,omitempty"`
	AmountDue       *sharedmodels.FloatOrString `json:"AmountDue,omitempty"`
	AmountPaid      *sharedmodels.FloatOrString `json:"AmountPaid,omitempty"`
	AmountCredited  *sharedmodels.FloatOrString `json:"AmountCredited,omitempty"`
	Reference       *string                     `json:"Reference,omitempty"`
	BrandingThemeID *string                     `json:"BrandingThemeID,omitempty"`
	Url             *string                     `json:"Url,omitempty"`
}

type Contact struct {
	ContactID      string    `json:"ContactID"`
	ContactStatus  *string   `json:"ContactStatus,omitempty"`
	Name           *string   `json:"Name,omitempty"`
	Addresses      []Address `json:"Addresses,omitempty"`
	Phones         []Phone   `json:"Phones,omitempty"`
	UpdatedDateUTC *string   `json:"UpdatedDateUTC,omitempty"`
	IsSupplier     *string   `json:"IsSupplier,omitempty"`
	IsCustomer     *string   `json:"IsCustomer,omitempty"`
}

type Address struct {
	AddressType  string  `json:"AddressType"`
	AddressLine1 *string `json:"AddressLine1,omitempty"`
	AddressLine2 *string `json:"AddressLine2,omitempty"`
	City         *string `json:"City,omitempty"`
	PostalCode   *string `json:"PostalCode,omitempty"`
}

type Phone struct {
	PhoneType string `json:"PhoneType"`
}

type LineItem struct {
	ItemCode    string                      `json:"ItemCode"`
	Description string                      `json:"Description"`
	Quantity    *sharedmodels.FloatOrString `json:"Quantity"`
	UnitAmount  *sharedmodels.FloatOrString `json:"UnitAmount"`
	TaxType     string                      `json:"TaxType"`
	TaxAmount   *sharedmodels.FloatOrString `json:"TaxAmount"`
	LineAmount  *sharedmodels.FloatOrString `json:"LineAmount"`
	AccountCode string                      `json:"AccountCode"`
	AccountID   *string                     `json:"AccountId,omitempty"`
	Item        *Item                       `json:"Item,omitempty"`
	Tracking    []Tracking                  `json:"Tracking,omitempty"`
	LineItemID  *string                     `json:"LineItemID,omitempty"`
}

type Item struct {
	ItemID string `json:"ItemID"`
	Name   string `json:"Name"`
	Code   string `json:"Code"`
}

type Tracking struct {
	TrackingCategoryID string `json:"TrackingCategoryID"`
	Name               string `json:"Name"`
	Option             string `json:"Option"`
}

type Payment struct {
	Date      string                      `json:"Date"`
	Amount    *sharedmodels.FloatOrString `json:"Amount"`
	PaymentID string                      `json:"PaymentID"`
}

type Invoices struct {
	Invoices []Invoice `json:"Invoices"`
}
