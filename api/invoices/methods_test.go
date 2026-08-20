package invoices

import (
	"net/http"
	"testing"
	"time"

	invoicesmodels "github.com/karman-digital/xero-go/models/invoices"
)

func TestGetInvoicesPath(t *testing.T) {
	path := getInvoicesPath(invoicesmodels.GetOptions{
		Page:        2,
		PageSize:    250,
		Where:       `Type=="ACCREC"`,
		Order:       "InvoiceID ASC",
		SummaryOnly: true,
	})
	want := "/Invoices?order=InvoiceID+ASC&page=2&pageSize=250&summaryOnly=true&where=Type%3D%3D%22ACCREC%22"
	if path != want {
		t.Fatalf("getInvoicesPath() = %q, want %q", path, want)
	}
}

func TestGetInvoicesHeaders(t *testing.T) {
	modifiedAfter := time.Date(2026, time.August, 20, 6, 30, 0, 0, time.FixedZone("BST", 60*60))
	headers := getInvoicesHeaders(invoicesmodels.GetOptions{ModifiedAfter: modifiedAfter})
	if got, want := headers.Get("If-Modified-Since"), "2026-08-20T05:30:00"; got != want {
		t.Fatalf("If-Modified-Since = %q, want %q", got, want)
	}
	if got := getInvoicesHeaders(invoicesmodels.GetOptions{}); len(got) != 0 {
		t.Fatalf("empty headers = %#v", got)
	}
	if http.CanonicalHeaderKey("if-modified-since") != "If-Modified-Since" {
		t.Fatal("unexpected canonical header name")
	}
}

func TestGetInvoicesPathWithoutOptions(t *testing.T) {
	if got := getInvoicesPath(invoicesmodels.GetOptions{}); got != "/Invoices" {
		t.Fatalf("getInvoicesPath() = %q, want /Invoices", got)
	}
}
