package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	ApiBaseUrl   string
	ApiKeyId     string
	ApiSecretKey string
	Region       string
	Project      string
	HTTPClient   *http.Client
}

type ClientConfig struct {
	ApiBaseUrl   string
	ApiKeyId     string
	ApiSecretKey string
	Region       string
	Project      string
}

func NewClient(cfg *ClientConfig) (*Client, error) {
	if cfg.ApiBaseUrl == "" {
		return nil, fmt.Errorf("api_base_url is required")
	}
	if cfg.ApiKeyId == "" {
		return nil, fmt.Errorf("api_key_id is required")
	}
	if cfg.ApiSecretKey == "" {
		return nil, fmt.Errorf("api_secret_key is required")
	}
	if cfg.Region == "" {
		return nil, fmt.Errorf("region is required")
	}
	if cfg.Project == "" {
		return nil, fmt.Errorf("project is required")
	}

	return &Client{
		ApiBaseUrl:   cfg.ApiBaseUrl,
		ApiKeyId:     cfg.ApiKeyId,
		ApiSecretKey: cfg.ApiSecretKey,
		Region:       cfg.Region,
		Project:      cfg.Project,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}, nil
}

func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.ApiBaseUrl+path, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set authentication headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key-Id", c.ApiKeyId)
	req.Header.Set("X-Api-Secret-Key", c.ApiSecretKey)
	req.Header.Set("X-Region", c.Region)
	req.Header.Set("X-Project", c.Project)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	return respBody, nil
}
