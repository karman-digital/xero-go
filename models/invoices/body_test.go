package invoicesmodels

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestLineItemSerializesXeroAccountIDSpelling(t *testing.T) {
	accountID := "11111111-1111-1111-1111-111111111111"
	encoded, err := json.Marshal(LineItem{AccountID: &accountID})
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}
	if !strings.Contains(string(encoded), `"AccountID":"`+accountID+`"`) {
		t.Fatalf("serialized LineItem = %s", encoded)
	}
	if strings.Contains(string(encoded), `"AccountId"`) {
		t.Fatalf("serialized LineItem used incorrect AccountId spelling: %s", encoded)
	}
}
