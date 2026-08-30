package accounts

import (
	"net/http"

	"github.com/karman-digital/xero-go/api/credentials"
)

type requestSender interface {
	SendRequest(method, path string, body []byte) (*http.Response, error)
}

type AccountsService struct {
	sender requestSender
}

func NewAccountsService(creds *credentials.Credentials) *AccountsService {
	return newAccountsService(creds)
}

func newAccountsService(sender requestSender) *AccountsService {
	return &AccountsService{sender: sender}
}
