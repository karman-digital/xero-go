package xeroapp

import (
	"github.com/karman-digital/xero-go/api/credentials"
	"github.com/karman-digital/xero-go/api/invoices"
)

func InitXero() *Xero {
	return &Xero{}
}

func (x *Xero) InitClient(creds *credentials.Credentials) {
	x.Credentials = creds
}

func NewApiClient(creds *credentials.Credentials) ApiClient {
	return ApiClient{
		Invoices: invoices.NewInvoicesService(creds),
	}
}
