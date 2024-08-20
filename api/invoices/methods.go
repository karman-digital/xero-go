package invoices

import invoicesmodels "github.com/karman-digital/xero-go/models/invoices"

func (s *InvoicesService) CreateInvoice(body invoicesmodels.Invoice) (invoicesmodels.Invoice, error) {
	return invoicesmodels.Invoice{}, nil
}

func (s *InvoicesService) GetInvoice(invoiceID string) (invoicesmodels.Invoices, error) {
	return invoicesmodels.Invoices{}, nil
}
