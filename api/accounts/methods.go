package accounts

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	xeroerrors "github.com/karman-digital/xero-go/app/errors"
	accountsmodels "github.com/karman-digital/xero-go/models/accounts"
)

func (service *AccountsService) GetAccounts() (accountsmodels.Accounts, error) {
	response, err := service.sender.SendRequest(http.MethodGet, "/Accounts", nil)
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
