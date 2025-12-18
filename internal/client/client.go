package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	ApiBaseUrl   string
	ApiKeyId     string
	ApiSecretKey string
	Region       string
	ProjectId    string
	HTTPClient   *http.Client
}

type ClientConfig struct {
	ApiBaseUrl   string
	ApiKeyId     string
	ApiSecretKey string
	Region       string
	ProjectId    string
}

// API Response structures
type ApiResponse[T any] struct {
	Success bool       `json:"success"`
	Data    T          `json:"data"`
	Errors  []ApiError `json:"errors"`
}

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Image struct {
	ID       int64 `json:"id"`
	IsCustom bool  `json:"isCustom"`
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
	if cfg.ProjectId == "" {
		return nil, fmt.Errorf("project is required")
	}

	return &Client{
		ApiBaseUrl:   cfg.ApiBaseUrl,
		ApiKeyId:     cfg.ApiKeyId,
		ApiSecretKey: cfg.ApiSecretKey,
		Region:       cfg.Region,
		ProjectId:    cfg.ProjectId,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}, nil
}

func (c *Client) GetImageBySlug(ctx context.Context, slug string) (*Image, error) {
	return c.getImage(ctx, fmt.Sprintf("slug=%s", url.QueryEscape(slug)))
}

func (c *Client) GetImageByName(ctx context.Context, name string) (*Image, error) {
	return c.getImage(ctx, fmt.Sprintf("name=%s", url.QueryEscape(name)))
}

func (c *Client) getImage(ctx context.Context, query string) (*Image, error) {
	path := fmt.Sprintf("/api/v2/image?%s", query)

	body, err := c.Get(ctx, path)
	if err != nil {
		return nil, err
	}

	var response ApiResponse[*Image]
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse image response: %w", err)
	}

	if !response.Success {
		return nil, fmt.Errorf("%s", formatApiErrors(response.Errors))
	}

	return response.Data, nil
}

// Helpers
func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.ApiBaseUrl+"/panel-main"+path, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set authentication headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key-Id", c.ApiKeyId)
	req.Header.Set("X-Api-Secret-Key", c.ApiSecretKey)
	req.Header.Set("X-Region", c.Region)
	req.Header.Set("X-Project-Id", c.ProjectId)

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

func (c *Client) Get(ctx context.Context, path string) ([]byte, error) {
	return c.doRequest(ctx, http.MethodGet, path, nil)
}

func formatApiErrors(errors []ApiError) string {
	if len(errors) == 0 {
		return "unknown error"
	}

	var messages []string
	for _, e := range errors {
		messages = append(messages, fmt.Sprintf("[%d] %s", e.Code, e.Message))
	}
	return strings.Join(messages, "; ")
}
