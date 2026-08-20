package contacts

import (
	"encoding/json"
	"testing"

	contactsmodels "github.com/karman-digital/xero-go/models/contacts"
)

func TestGetContactsPath(t *testing.T) {
	path := getContactsPath(contactsmodels.GetOptions{
		Page:            2,
		PageSize:        250,
		IncludeArchived: true,
		Order:           "ContactID ASC",
		SummaryOnly:     true,
	})
	want := "/Contacts?includeArchived=true&order=ContactID+ASC&page=2&pageSize=250&summaryOnly=true"
	if path != want {
		t.Fatalf("getContactsPath() = %q, want %q", path, want)
	}
}

func TestContactDecodesMergeDestination(t *testing.T) {
	var contact contactsmodels.Contact
	if err := json.Unmarshal([]byte(`{"ContactID":"old","MergedToContactID":"new"}`), &contact); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if contact.MergedToContactID != "new" {
		t.Fatalf("MergedToContactID = %q, want new", contact.MergedToContactID)
	}
}
