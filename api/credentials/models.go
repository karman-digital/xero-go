package credentials

import (
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

type Credentials struct {
	client         *retryablehttp.Client
	accesstoken    string
	refreshtoken   string
	clientId       string
	clientSecret   string
	redirectUri    string
	organisationId string
	expiresAt      time.Time
}
