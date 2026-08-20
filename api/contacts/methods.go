package contacts

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/karman-digital/xero-go/api/shared"
	xeroerrors "github.com/karman-digital/xero-go/app/errors"
	contactsmodels "github.com/karman-digital/xero-go/models/contacts"
)

func (s *ContactsService) GetContacts(options contactsmodels.GetOptions) (contactsmodels.Contacts, error) {
	resp, err := s.SendRequest(http.MethodGet, getContactsPath(options), nil)
	if err != nil {
		return contactsmodels.Contacts{}, xeroerrors.New(xeroerrors.ErrApiError, err.Error())
	}
	defer resp.Body.Close()
	return shared.HandleContactResponse(*resp)
}

func getContactsPath(options contactsmodels.GetOptions) string {
	query := url.Values{}
	if options.IncludeArchived {
		query.Set("includeArchived", "true")
	}
	if options.Order != "" {
		query.Set("order", options.Order)
	}
	if options.Page > 0 {
		query.Set("page", strconv.Itoa(options.Page))
	}
	if options.PageSize > 0 {
		query.Set("pageSize", strconv.Itoa(options.PageSize))
	}
	if options.SummaryOnly {
		query.Set("summaryOnly", "true")
	}
	if len(query) == 0 {
		return "/Contacts"
	}
	return "/Contacts?" + query.Encode()
}

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
