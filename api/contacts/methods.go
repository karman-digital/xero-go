package contacts

import (
	"encoding/json"
	"net/http"

	"github.com/karman-digital/xero-go/api/shared"
	xeroerrors "github.com/karman-digital/xero-go/app/errors"
	contactsmodels "github.com/karman-digital/xero-go/models/contacts"
)

func (s *ContactsService) CreateContact(body contactsmodels.Contacts) (contactsmodels.Contacts, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return contactsmodels.Contacts{}, xeroerrors.New(xeroerrors.ErrInternal, err.Error())
	}
	resp, err := s.SendRequest(http.MethodPost, "/Contacts", reqBody)
	if err != nil {
		return contactsmodels.Contacts{}, xeroerrors.New(xeroerrors.ErrApiError, err.Error())
	}
	defer resp.Body.Close()
	return shared.HandleContactResponse(*resp)
}

func (s *ContactsService) GetContact(id string) (contactsmodels.Contacts, error) {
	resp, err := s.SendRequest(http.MethodGet, "/Contacts/"+id, nil)
	if err != nil {
		return contactsmodels.Contacts{}, xeroerrors.New(xeroerrors.ErrApiError, err.Error())
	}
	defer resp.Body.Close()
	return shared.HandleContactResponse(*resp)
}
