package invoices

import (
	"testing"

	invoicesmodels "github.com/karman-digital/xero-go/models/invoices"
)

func TestGetInvoicesPath(t *testing.T) {
	path := getInvoicesPath(invoicesmodels.GetOptions{
		Page:     2,
		PageSize: 250,
		Where:    `Type=="ACCREC"`,
	})
	want := "/Invoices?page=2&pageSize=250&where=Type%3D%3D%22ACCREC%22"
	if path != want {
		t.Fatalf("getInvoicesPath() = %q, want %q", path, want)
	}
}

func TestGetInvoicesPathWithoutOptions(t *testing.T) {
	if got := getInvoicesPath(invoicesmodels.GetOptions{}); got != "/Invoices" {
		t.Fatalf("getInvoicesPath() = %q, want /Invoices", got)
	}
}
