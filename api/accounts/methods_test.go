package accounts

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	accountsmodels "github.com/karman-digital/xero-go/models/accounts"
)

type fakeSender struct {
	method string
	path   string
	body   []byte
	resp   *http.Response
	err    error
}

func (sender *fakeSender) SendRequest(method, path string, body []byte) (*http.Response, error) {
	sender.method = method
	sender.path = path
	sender.body = body
	return sender.resp, sender.err
}

func TestGetAccountsDecodesCompleteResponse(t *testing.T) {
	sender := &fakeSender{resp: response(http.StatusOK, `{
		"Id":"request-1",
		"Status":"OK",
		"ProviderName":"Karman Digital",
		"DateTimeUTC":"2026-08-30T12:00:00Z",
		"Accounts":[
			{"AccountID":"11111111-1111-1111-1111-111111111111","Code":"4100","Name":"Partnership","Type":"REVENUE","Status":"ACTIVE","UpdatedDateUTC":"/Date(1788091200000+0000)/"},
			{"AccountID":"22222222-2222-2222-2222-222222222222","Name":"Old revenue","Type":"REVENUE","Status":"ARCHIVED"}
		]
	}`)}

	got, err := newAccountsService(sender).GetAccounts(accountsmodels.GetOptions{Where: `Type=="REVENUE"`})
	if err != nil {
		t.Fatalf("GetAccounts() error = %v", err)
	}
	if sender.method != http.MethodGet || sender.path != "/Accounts?where=Type%3D%3D%22REVENUE%22" || sender.body != nil {
		t.Fatalf("request = %s %s %#v", sender.method, sender.path, sender.body)
	}
	if got.ID != "request-1" || got.Status != "OK" || got.ProviderName != "Karman Digital" || got.DateTimeUTC != "2026-08-30T12:00:00Z" {
		t.Fatalf("response metadata = %#v", got)
	}
	if len(got.Accounts) != 2 {
		t.Fatalf("len(Accounts) = %d, want 2", len(got.Accounts))
	}
	if account := got.Accounts[0]; account.AccountID != "11111111-1111-1111-1111-111111111111" || account.Code != "4100" || account.Name != "Partnership" || account.Type != "REVENUE" || account.Status != "ACTIVE" || account.UpdatedDateUTC != "/Date(1788091200000+0000)/" {
		t.Fatalf("active account = %#v", account)
	}
	if account := got.Accounts[1]; account.Status != "ARCHIVED" || account.Code != "" || account.UpdatedDateUTC != "" {
		t.Fatalf("archived account with optional fields = %#v", account)
	}
}

func TestGetAccountsRejectsProviderFailure(t *testing.T) {
	sender := &fakeSender{resp: response(http.StatusBadGateway, `{"ErrorNumber":10,"Message":"provider unavailable"}`)}
	if _, err := newAccountsService(sender).GetAccounts(accountsmodels.GetOptions{}); err == nil || !strings.Contains(err.Error(), "xero api error") {
		t.Fatalf("GetAccounts() error = %v, want Xero API error", err)
	}
}

func TestGetAccountsRejectsMalformedResponse(t *testing.T) {
	sender := &fakeSender{resp: response(http.StatusOK, `{"Accounts":[`)}
	if _, err := newAccountsService(sender).GetAccounts(accountsmodels.GetOptions{}); err == nil {
		t.Fatal("GetAccounts() error = nil, want decode error")
	}
}

func TestGetAccountsReturnsRequestFailure(t *testing.T) {
	want := errors.New("network unavailable")
	sender := &fakeSender{err: want}
	if _, err := newAccountsService(sender).GetAccounts(accountsmodels.GetOptions{}); err == nil || !strings.Contains(err.Error(), want.Error()) {
		t.Fatalf("GetAccounts() error = %v, want request failure", err)
	}
}

func response(status int, body string) *http.Response {
	return &http.Response{StatusCode: status, Status: http.StatusText(status), Body: io.NopCloser(strings.NewReader(body))}
}
