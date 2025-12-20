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
	baseURL      string
	apiKeyID     string
	apiSecretKey string
	region       string
	projectID    string
	httpClient   *http.Client
}

type Config struct {
	APIBaseURL   string
	APIKeyID     string
	APISecretKey string
}

func New(cfg Config) (*Client, error) {
	if cfg.APIBaseURL == "" || cfg.APIKeyID == "" || cfg.APISecretKey == "" {
		return nil, fmt.Errorf("all config fields are required")
	}

	return &Client{
		baseURL:      strings.TrimRight(cfg.APIBaseURL, "/") + "/panel-main",
		apiKeyID:     cfg.APIKeyID,
		apiSecretKey: cfg.APISecretKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}, nil
}

// HTTP helpers
type apiResponse[T any] struct {
	Success bool       `json:"success"`
	Data    T          `json:"data"`
	Errors  []apiError `json:"errors"`
}

type apiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *Client) get(ctx context.Context, path string, result any) error {
	return c.do(ctx, http.MethodGet, path, nil, result)
}

func (c *Client) post(ctx context.Context, path string, body, result any) error {
	return c.do(ctx, http.MethodPost, path, body, result)
}

func (c *Client) delete(ctx context.Context, path string) error {
	return c.do(ctx, http.MethodDelete, path, nil, nil)
}

func (c *Client) do(ctx context.Context, method, path string, body, result any) error {
	var reqBody io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal request: %w", err)
		}
		reqBody = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, reqBody)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key-Id", c.apiKeyID)
	req.Header.Set("X-Api-Secret-Key", c.apiSecretKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	var apiResp apiResponse[json.RawMessage]

	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return fmt.Errorf("parse response: %w", err)
	}

	if !apiResp.Success {
		return fmt.Errorf("api error: %s", formatErrors(apiResp.Errors))
	}

	if result == nil {
		return nil
	}

	if err := json.Unmarshal(apiResp.Data, result); err != nil {
		return fmt.Errorf("parse data: %w", err)
	}

	return nil
}

func formatErrors(errs []apiError) string {
	if len(errs) == 0 {
		return "unknown error"
	}
	msgs := make([]string, len(errs))
	for i, e := range errs {
		msgs[i] = fmt.Sprintf("[%d] %s", e.Code, e.Message)
	}
	return strings.Join(msgs, "; ")
}

// Image
type Image struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	IsCustom bool   `json:"isCustom"`
}

func (c *Client) GetImageBySlug(ctx context.Context, slug string) (*Image, error) {
	var img Image
	err := c.get(ctx, "/api/v2/image?slug="+url.QueryEscape(slug), &img)
	return &img, err
}

func (c *Client) GetImageByName(ctx context.Context, name string) (*Image, error) {
	var img Image
	err := c.get(ctx, "/api/v2/image?name="+url.QueryEscape(name), &img)
	return &img, err
}
