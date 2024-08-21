package interfaces

import invoicesmodels "github.com/karman-digital/xero-go/models/invoices"

type Invoices interface {
	CreateInvoice(body invoicesmodels.Invoice) (invoicesmodels.Invoice, error)
}

type Validator interface {
	ValidateXeroToken() error
}
