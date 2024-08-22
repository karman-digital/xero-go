package contacts

import "github.com/karman-digital/xero-go/api/credentials"

func NewContactsService(creds *credentials.Credentials) *ContactsService {
	return &ContactsService{creds}
}
