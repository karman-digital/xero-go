package interfaces

import (
	contactsmodels "github.com/karman-digital/xero-go/models/contacts"
	invoicesmodels "github.com/karman-digital/xero-go/models/invoices"
)

type Invoices interface {
	CreateInvoice(body invoicesmodels.Invoices) (invoicesmodels.Invoice, error)
}

type Contacts interface {
	CreateContact(body contactsmodels.Contact) (contactsmodels.Contact, error)
	GetContact(id string) (contactsmodels.Contact, error)
}

type Validator interface {
	ValidateXeroToken() error
}
