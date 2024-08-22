package contacts

import contactsmodels "github.com/karman-digital/xero-go/models/contacts"

func (s *ContactsService) CreateContact(body contactsmodels.Contact) (contactsmodels.Contact, error) {
	return contactsmodels.Contact{}, nil
}

func (s *ContactsService) GetContact(id string) (contactsmodels.Contact, error) {
	return contactsmodels.Contact{}, nil
}
