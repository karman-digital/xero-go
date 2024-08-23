package shared

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	xeroerrors "github.com/karman-digital/xero-go/app/errors"
	contactsmodels "github.com/karman-digital/xero-go/models/contacts"
	invoicesmodels "github.com/karman-digital/xero-go/models/invoices"
)

func HandleContactResponse(resp http.Response) (objResp contactsmodels.Contacts, err error) {
	rawBody, err := handleBasicResponseCode(&resp)
	if err != nil {
		return contactsmodels.Contacts{}, err
	}
	err = json.Unmarshal(rawBody, &objResp)
	if err != nil {
		return contactsmodels.Contacts{}, err
	}
	return objResp, nil
}

func HandleInvoiceResponse(resp http.Response) (objResp invoicesmodels.Invoices, err error) {
	rawBody, err := handleBasicResponseCode(&resp)
	if err != nil {
		return invoicesmodels.Invoices{}, err
	}
	err = json.Unmarshal(rawBody, &objResp)
	if err != nil {
		return invoicesmodels.Invoices{}, err
	}
	return objResp, nil
}

func handleBasicResponseCode(resp *http.Response) (rawBody []byte, err error) {
	rawBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, xeroerrors.New(xeroerrors.ErrInternal, fmt.Sprintf("error reading response body: %s", err))
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 404 {
			return rawBody, xeroerrors.New(xeroerrors.ErrNotFound, "resource not found")
		}
		return rawBody, xeroerrors.New(xeroerrors.ErrApiError, fmt.Sprintf("error body: %s", string(rawBody)))
	}
	return rawBody, nil
}
