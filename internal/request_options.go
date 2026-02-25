package internal

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

// WithApiMethod sets the API method used (Logging).
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

// WithHttpMethod sets the request method (Default to GET).
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

// WithParam sets one param on the current request params.
func WithParam(key, val string) RequestOption {
	return func(ro *requestOptions) {
		if ro.params == nil {
			ro.params = make(map[string]string)
		}
		ro.params[key] = val
	}
}
