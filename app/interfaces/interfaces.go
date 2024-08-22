package interfaces

import invoicesmodels "github.com/karman-digital/xero-go/models/invoices"

type Invoices interface {
	CreateInvoice(body invoicesmodels.Invoice) (invoicesmodels.Invoice, error)
}

type Contacts interface {
	CreateContact(body invoicesmodels.Contact) (invoicesmodels.Contact, error)
	GetContact(id string) (invoicesmodels.Contact, error)
}

type Validator interface {
	ValidateXeroToken() error
}
