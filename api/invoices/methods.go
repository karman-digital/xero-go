package invoices

import (
	"encoding/json"
	"net/http"

	"github.com/karman-digital/xero-go/api/shared"
	xeroerrors "github.com/karman-digital/xero-go/app/errors"
	invoicesmodels "github.com/karman-digital/xero-go/models/invoices"
)

func (s *InvoicesService) CreateInvoice(body invoicesmodels.Invoices) (invoicesmodels.Invoices, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return invoicesmodels.Invoices{}, xeroerrors.New(xeroerrors.ErrInternal, err.Error())
	}
	resp, err := s.SendRequest(http.MethodPost, "/Invoices", reqBody)
	if err != nil {
		return invoicesmodels.Invoices{}, xeroerrors.New(xeroerrors.ErrApiError, err.Error())
	}
	defer resp.Body.Close()
	return shared.HandleInvoiceResponse(*resp)
}

func (s *InvoicesService) GetInvoice(invoiceID string) (invoicesmodels.Invoices, error) {
	resp, err := s.SendRequest(http.MethodGet, "/Invoices/"+invoiceID, nil)
	if err != nil {
		return invoicesmodels.Invoices{}, xeroerrors.New(xeroerrors.ErrApiError, err.Error())
	}
	defer resp.Body.Close()
	return shared.HandleInvoiceResponse(*resp)
}
