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

// AuthRequest makes a authenticated request with the provided client and returns the decode value to the expected generic type.
func AuthRequest[T any](ctx context.Context, client Doer, apiKey, uri, method string, params map[string]string) (T, error) {
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

	return do[T](client, req)
}

// Request makes a authenticated request with the provided client and returns the decode value to the expected generic type.
func Request[T any](ctx context.Context, client Doer, uri, method string) (T, error) {
	req, err := http.NewRequestWithContext(ctx, method, uri, http.NoBody)
	if err != nil {
		var zero T
		return zero, err
	}

	return do[T](client, req)
}

// do Executes the request itself and handles the status and unmarshal.
func do[T any](client Doer, req *http.Request) (T, error) {
	var respData T

	resp, err := client.Do(req)
	if err != nil {
		return respData, err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return respData, err
	}

	if resp.StatusCode != http.StatusOK {
		return respData, &RiotError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			Body:       strings.TrimSpace(string(body)),
		}
	}

	if err := json.Unmarshal(body, &respData); err != nil {
		return respData, err
	}

	return respData, nil
}
