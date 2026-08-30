package interfaces

import (
	accountsmodels "github.com/karman-digital/xero-go/models/accounts"
	contactsmodels "github.com/karman-digital/xero-go/models/contacts"
	invoicesmodels "github.com/karman-digital/xero-go/models/invoices"
)

type Accounts interface {
	GetAccounts(accountsmodels.GetOptions) (accountsmodels.Accounts, error)
}

type Invoices interface {
	CreateInvoice(body invoicesmodels.Invoices) (invoicesmodels.Invoices, error)
	GetInvoice(id string) (invoicesmodels.Invoices, error)
	GetInvoices(options invoicesmodels.GetOptions) (invoicesmodels.Invoices, error)
}

type Contacts interface {
	CreateContact(body contactsmodels.Contacts) (contactsmodels.Contacts, error)
	GetContact(id string) (contactsmodels.Contacts, error)
	GetContacts(options contactsmodels.GetOptions) (contactsmodels.Contacts, error)
}

type Validator interface {
	ValidateXeroToken() error
}
