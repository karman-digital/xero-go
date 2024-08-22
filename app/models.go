package xeroapp

import (
	"github.com/karman-digital/xero-go/api/credentials"
	"github.com/karman-digital/xero-go/app/interfaces"
)

type ApiClient struct {
	interfaces.Invoices
	interfaces.Contacts
}

type Xero struct {
	*credentials.Credentials
	ApiClient
}
