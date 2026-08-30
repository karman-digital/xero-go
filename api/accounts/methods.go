package accounts

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	xeroerrors "github.com/karman-digital/xero-go/app/errors"
	accountsmodels "github.com/karman-digital/xero-go/models/accounts"
)

func (service *AccountsService) GetAccounts(options accountsmodels.GetOptions) (accountsmodels.Accounts, error) {
	response, err := service.sender.SendRequest(http.MethodGet, getAccountsPath(options), nil)
	if err != nil {
		return accountsmodels.Accounts{}, xeroerrors.New(xeroerrors.ErrApiError, err.Error())
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return accountsmodels.Accounts{}, xeroerrors.New(xeroerrors.ErrInternal, fmt.Sprintf("read accounts response: %v", err))
	}
	if response.StatusCode != http.StatusOK {
		return accountsmodels.Accounts{}, xeroerrors.New(xeroerrors.ErrApiError, fmt.Sprintf("unexpected status %d", response.StatusCode))
	}
	var accounts accountsmodels.Accounts
	if err := json.Unmarshal(body, &accounts); err != nil {
		return accountsmodels.Accounts{}, xeroerrors.New(xeroerrors.ErrInternal, fmt.Sprintf("decode accounts response: %v", err))
	}
	return accounts, nil
}

func getAccountsPath(options accountsmodels.GetOptions) string {
	query := url.Values{}
	if options.Where != "" {
		query.Set("where", options.Where)
	}
	if len(query) == 0 {
		return "/Accounts"
	}
	return "/Accounts?" + query.Encode()
}
