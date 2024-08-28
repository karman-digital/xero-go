package webhookmodels

import (
	"encoding/json"
	"time"
)

const xeroTimeFormat = "2006-01-02T15:04:05.000"

type Event struct {
	ResourceUrl   string      `json:"resourceUrl"`
	ResourceId    string      `json:"resourceId"`
	EventDateUtc  xeroUtcDate `json:"eventDateUtc"`
	EventType     string      `json:"eventType"`
	EventCategory string      `json:"eventCategory"`
	TenantId      string      `json:"tenantId"`
	TenantType    string      `json:"tenantType"`
}

type WebhookPayload struct {
	Events             []Event `json:"events"`
	LastEventSequence  int     `json:"lastEventSequence"`
	FirstEventSequence int     `json:"firstEventSequence"`
	Entropy            string  `json:"entropy"`
}

type xeroUtcDate time.Time

func (t *xeroUtcDate) UnmarshalJSON(b []byte) error {
	str := string(b)
	if len(str) >= 2 && str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
	}
	parsedTime, err := time.Parse(xeroTimeFormat, str)
	if err != nil {
		return err
	}
	*t = xeroUtcDate(parsedTime)
	return nil
}

func (t xeroUtcDate) MarshalJSON() ([]byte, error) {
	formatted := time.Time(t).UTC().Format(xeroTimeFormat)
	return json.Marshal(formatted)
}
