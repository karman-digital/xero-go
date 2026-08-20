package invoices

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

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

func (s *InvoicesService) GetInvoices(options invoicesmodels.GetOptions) (invoicesmodels.Invoices, error) {
	resp, err := s.SendRequestWithHeaders(http.MethodGet, getInvoicesPath(options), nil, getInvoicesHeaders(options))
	if err != nil {
		return invoicesmodels.Invoices{}, xeroerrors.New(xeroerrors.ErrApiError, err.Error())
	}
	defer resp.Body.Close()
	return shared.HandleInvoiceResponse(*resp)
}

func getInvoicesPath(options invoicesmodels.GetOptions) string {
	query := url.Values{}
	if options.Page > 0 {
		query.Set("page", strconv.Itoa(options.Page))
	}
	if options.PageSize > 0 {
		query.Set("pageSize", strconv.Itoa(options.PageSize))
	}
	if options.Where != "" {
		query.Set("where", options.Where)
	}
	if options.Order != "" {
		query.Set("order", options.Order)
	}
	if options.SummaryOnly {
		query.Set("summaryOnly", "true")
	}
	if len(query) == 0 {
		return "/Invoices"
	}
	return "/Invoices?" + query.Encode()
}

func getInvoicesHeaders(options invoicesmodels.GetOptions) http.Header {
	headers := http.Header{}
	if !options.ModifiedAfter.IsZero() {
		headers.Set("If-Modified-Since", options.ModifiedAfter.UTC().Format("2006-01-02T15:04:05"))
	}
	return headers
}
