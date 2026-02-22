package internal

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const apiTokenHeader = "X-Riot-Token"

type (
	requestOptions struct {
		method string
	}

	RequestOption func(*requestOptions)
)

func WithMethod(method string) RequestOption {
	return func(o *requestOptions) { o.method = method }
}

// AuthRequest makes a authenticated request with the provided client and returns the decode value to the expected generic type.
func AuthRequest[T any](ctx context.Context, client *Client, apiKey, uri, method string, params map[string]string, opts ...RequestOption) (T, error) {
	var zero T
	u, err := url.Parse(uri)
	if err != nil {
		return zero, err
	}

	// Add any query param needed, some of the APIs use it for filtering.
	query := u.Query()
	for key, val := range params {
		query.Add(key, val)
	}

	u.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, method, u.String(), http.NoBody)
	if err != nil {
		return zero, err
	}

	req.Header.Set(apiTokenHeader, apiKey)

	return do[T](client, req, opts...)
}

// Request makes a authenticated request with the provided client and returns the decode value to the expected generic type.
func Request[T any](ctx context.Context, client *Client, uri, method string, opts ...RequestOption) (T, error) {
	req, err := http.NewRequestWithContext(ctx, method, uri, http.NoBody)
	if err != nil {
		var zero T
		return zero, err
	}

	return do[T](client, req, opts...)
}

// do Executes the request itself and handles the status and unmarshal.
func do[T any](client *Client, req *http.Request, opts ...RequestOption) (T, error) {
	var respData T

	ro := &requestOptions{}
	for _, o := range opts {
		o(ro)
	}

	logger := client.Logger.With(
		"method", ro.method,
		"uri", req.URL.String(),
		"route", client.routePrefix,
	)

	resp, err := client.Http.Do(req)
	if err != nil {
		logger.Error("request failed", "error", err)
		return respData, err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("failed to read response body", "error", err)
		return respData, err
	}

	if resp.StatusCode != http.StatusOK {
		logger.Warn("non-OK HTTP status", "status", resp.StatusCode)
		return respData, &RiotError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			Body:       strings.TrimSpace(string(body)),
		}
	}

	if err := json.Unmarshal(body, &respData); err != nil {
		logger.Error("failed to unmarshal response body", "error", err)
		return respData, err
	}

	return respData, nil
}
