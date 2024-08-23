package credentials

import (
	"fmt"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

func (c *Credentials) SendRequest(method, path string, body []byte) (*http.Response, error) {
	req, err := retryablehttp.NewRequest(method, "https://api.xero.com/api.xro/2.0"+path, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accesstoken))
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %s", err)
	}
	return resp, nil
}
