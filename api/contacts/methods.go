package contacts

import invoicesmodels "github.com/karman-digital/xero-go/models/invoices"

func (s *ContactsService) CreateContact(body invoicesmodels.Contact) (invoicesmodels.Contact, error) {
	return invoicesmodels.Contact{}, nil
}

func (s *ContactsService) GetContact(id string) (invoicesmodels.Contact, error) {
	return invoicesmodels.Contact{}, nil
}
