package credentials

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	oauthmodels "github.com/karman-digital/xero-go/models/oauth"
)

func GenerateTokenPair(code, clientId, clientSecret, redirectUri string) (oauthmodels.TokenResponse, error) {
	var respBody oauthmodels.TokenResponse
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", redirectUri)
	req, err := http.NewRequest(http.MethodPost, "https://identity.xero.com/connect/token", strings.NewReader(data.Encode()))
	if err != nil {
		return respBody, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", clientId, clientSecret))))
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return respBody, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return respBody, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return respBody, err
	}
	return respBody, nil
}

func (c *Credentials) SetTokens(tokens oauthmodels.TokenResponse) {
	c.accesstoken = tokens.AccessToken
	c.refreshtoken = tokens.RefreshToken
	c.expiresAt = time.Now().Add(time.Duration(tokens.ExpiresIn) * time.Second)
}

func (c *Credentials) RefreshToken() error {
	var respBody oauthmodels.TokenResponse
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", c.refreshtoken)
	req, err := http.NewRequest(http.MethodPost, "https://api.hubapi.com/oauth/v1/token", strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.clientId, c.clientSecret))))
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return err
	}
	c.SetTokens(respBody)
	return nil
}

func (c *Credentials) AccessTokenValid() bool {
	return c.expiresAt.Before(time.Now())
}
