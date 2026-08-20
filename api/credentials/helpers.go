package credentials

import (
	"fmt"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

func (c *Credentials) SendRequest(method, path string, body []byte) (*http.Response, error) {
	return c.SendRequestWithHeaders(method, path, body, nil)
}

func (c *Credentials) SendRequestWithHeaders(method, path string, body []byte, headers http.Header) (*http.Response, error) {
	req, err := retryablehttp.NewRequest(method, "https://api.xero.com/api.xro/2.0"+path, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accesstoken))
	req.Header.Set("Xero-tenant-id", c.organisationId)
	for name, values := range headers {
		for _, value := range values {
			req.Header.Add(name, value)
		}
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %s", err)
	}
	return resp, nil
}
