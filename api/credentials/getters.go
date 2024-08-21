package credentials

import "time"

func (c Credentials) GetAccessToken() string {
	return c.accesstoken
}

func (c Credentials) GetRefreshToken() string {
	return c.refreshtoken
}

func (c Credentials) GetExpiresAt() time.Time {
	return c.expiresAt
}
