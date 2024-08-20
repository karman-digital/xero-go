package credentials

import (
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

func NewXeroOauthCredentials(clientId, clientSecret, redirectUri, accesstoken, refreshtoken, organisationId string, expiresAt time.Time) *Credentials {
	client := retryablehttp.NewClient()
	client.Logger = nil
	return &Credentials{
		client:         client,
		clientId:       clientId,
		clientSecret:   clientSecret,
		redirectUri:    redirectUri,
		accesstoken:    accesstoken,
		refreshtoken:   refreshtoken,
		organisationId: organisationId,
		expiresAt:      expiresAt,
	}
}
