package invoices

import "github.com/karman-digital/xero-go/api/credentials"

func NewInvoicesService(creds *credentials.Credentials) *InvoicesService {
	return &InvoicesService{creds}
}
