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
		apiKey     string
		apiMethod  string
		httpMethod string
		body       any
		params     map[string]string
	}

	RequestOption func(*requestOptions)
)

// withApiKey applies the API Key for the AuthRequest.
func withApiKey(apiKey string) RequestOption {
	return func(ro *requestOptions) {
		ro.apiKey = apiKey
	}
}

// WithApiMethod sets the API method used (Logging)
func WithApiMethod(method string) RequestOption {
	return func(ro *requestOptions) {
		ro.apiMethod = method
	}
}

// WithBody sets the request body.
func WithBody(body any) RequestOption {
	return func(ro *requestOptions) {
		ro.body = body
	}
}

// WithHttpMethod sets the request method (Default to GET)
func WithHttpMethod(method string) RequestOption {
	return func(ro *requestOptions) {
		ro.httpMethod = method
	}
}

// WithParams sets the request params.
func WithParams(params map[string]string) RequestOption {
	return func(ro *requestOptions) {
		ro.params = params
	}
}

// AuthRequest makes a authenticated request with the provided client and returns the decode value to the expected generic type.
func AuthRequest[T any](ctx context.Context, client *Client, uri string, opts ...RequestOption) (T, error) {
	return Request[T](
		ctx,
		client,
		uri,
		append(opts, withApiKey(client.ApiKey))...,
	)
}

// Request is the basic request implementation with multiple options.
func Request[T any](ctx context.Context, client *Client, uri string, opts ...RequestOption) (T, error) {
	var ro requestOptions
	for _, o := range opts {
		o(&ro)
	}

	req, err := buildRequest(ctx, uri, &ro)
	if err != nil {
		var zero T
		return zero, err
	}

	return do[T](client, req, &ro)
}

// buildRequest mounts a new http request with all passed options.
func buildRequest(ctx context.Context, uri string, opts *requestOptions) (*http.Request, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	// Add any query param needed, some of the APIs use it for filtering.
	query := u.Query()
	for key, val := range opts.params {
		query.Add(key, val)
	}

	u.RawQuery = query.Encode()

	var bodyReader io.Reader = http.NoBody
	if opts.body != nil {
		b, err := json.Marshal(opts.body)
		if err != nil {
			return nil, err
		}
		bodyReader = strings.NewReader(string(b))
	}

	method := opts.httpMethod
	if method == "" {
		method = http.MethodGet
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), bodyReader)
	if err != nil {
		return nil, err
	}

	if opts.body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if opts.apiKey != "" {
		req.Header.Set(apiTokenHeader, opts.apiKey)
	}

	return req, nil
}

// do Executes the request itself and handles the status and unmarshal.
func do[T any](client *Client, req *http.Request, ro *requestOptions) (T, error) {
	var respData T

	logger := client.Logger.With(
		"apiMethod", ro.apiMethod,
		"httpMethod", ro.httpMethod,
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

	if resp.StatusCode < 200 || resp.StatusCode > 300 {
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
