package webhookmodels

import "time"

type Event struct {
	ResourceUrl   string    `json:"resourceUrl"`
	ResourceId    string    `json:"resourceId"`
	EventDateUtc  time.Time `json:"eventDateUtc"`
	EventType     string    `json:"eventType"`
	EventCategory string    `json:"eventCategory"`
	TenantId      string    `json:"tenantId"`
	TenantType    string    `json:"tenantType"`
}

type WebhookPayload struct {
	Events             []Event `json:"events"`
	LastEventSequence  int     `json:"lastEventSequence"`
	FirstEventSequence int     `json:"firstEventSequence"`
	Entropy            string  `json:"entropy"`
}
