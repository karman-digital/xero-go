package invoicesmodels

import "time"

type GetOptions struct {
	ModifiedAfter  time.Time
	Statuses       []string
	IDs            []string
	InvoiceNumbers []string
	ContactIDs     []string
	Where          string
	Order          string
	Page           int
	SummaryOnly    bool
	SearchTerm     string
	CreatedByMyApp bool
}
