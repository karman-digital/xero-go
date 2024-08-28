package interfaces

import (
	contactsmodels "github.com/karman-digital/xero-go/models/contacts"
	invoicesmodels "github.com/karman-digital/xero-go/models/invoices"
)

type Invoices interface {
	CreateInvoice(body invoicesmodels.Invoices) (invoicesmodels.Invoices, error)
	GetInvoice(id string) (invoicesmodels.Invoices, error)
}

type Contacts interface {
	CreateContact(body contactsmodels.Contacts) (contactsmodels.Contacts, error)
	GetContact(id string) (contactsmodels.Contacts, error)
}

type Validator interface {
	ValidateXeroToken() error
}
