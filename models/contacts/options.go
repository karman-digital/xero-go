package contacts

import "time"

type GetOptions struct {
	ModifiedAfter   time.Time
	Statuses        []string
	IDs             []string
	Where           string
	Order           string
	Page            int
	SummaryOnly     bool
	IncludeArchived bool
	SearchTerm      string
}
