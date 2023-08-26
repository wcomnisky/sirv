package sirv

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	BaseURL = "https://api.sirv.com/v2"
)

var (
	EnterprisePlan = PlanLimit{14000, 2000, 400, 400, 4000}
	BusinessPlan   = PlanLimit{7000, 1000, 200, 200, 2000}
	FreePlan       = PlanLimit{500, 50, 20, 20, 300}
)

func NewClient(httpClient *http.Client, limit PlanLimit) *Client {
	return &Client{
		BaseURL:    BaseURL,
		HTTPClient: httpClient,
		Limit:      limit,
	}
}

func (c *Client) GetToken(ctx context.Context, payload AuthPayload) (*TokenResponse, error) {
	var tokenResp TokenResponse
	err := c.makeRequest(ctx, http.MethodPost, c.BaseURL+"/token", payload, &tokenResp)
	if err != nil {
		return nil, err
	}

	c.Token = tokenResp.Token
	return &tokenResp, nil
}

func (c *Client) makeHTTPRequest(ctx context.Context, method, url string, body io.Reader, isJson bool) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.Token)

	if isJson {
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return resp, nil
}

func (c *Client) makeRequest(ctx context.Context, method, url string, requestBody interface{}, result interface{}) error {
	var body io.Reader
	if requestBody != nil {
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return err
		}
		body = bytes.NewReader(bodyBytes)
	}

	resp, err := c.makeHTTPRequest(ctx, method, url, body, true)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(result)
}

func (c *Client) makeSimpleRequest(ctx context.Context, method, url string, body io.Reader) (*http.Response, error) {
	return c.makeHTTPRequest(ctx, method, url, body, false)
}
