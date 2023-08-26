package sirv

import (
	"context"
	"net/http"
)

func (c *Client) GetAccountInfo(ctx context.Context) (*AccountInfo, error) {
	var accountInfo AccountInfo
	err := c.makeRequest(ctx, http.MethodGet, c.BaseURL+"/account", nil, &accountInfo)
	if err != nil {
		return nil, err
	}

	return &accountInfo, nil
}

func (c *Client) GetAPILimits(ctx context.Context) (*APILimits, error) {
	var limits APILimits
	err := c.makeRequest(ctx, http.MethodGet, c.BaseURL+"/account/limits", nil, &limits)
	if err != nil {
		return nil, err
	}

	return &limits, nil
}

func (c *Client) GetStorageInfo(ctx context.Context) (*StorageInfo, error) {
	var storageInfo StorageInfo
	err := c.makeRequest(ctx, http.MethodGet, c.BaseURL+"/account/storage", nil, &storageInfo)
	if err != nil {
		return nil, err
	}
	return &storageInfo, nil
}

func (c *Client) GetUsers(ctx context.Context) ([]User, error) {
	var users []User
	err := c.makeRequest(ctx, http.MethodGet, c.BaseURL+"/account/users", nil, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
