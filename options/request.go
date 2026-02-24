package options

import "leago/internal"

type (
	// PublicOption is a safe subset of RequestOption that is exposed to the public client.
	// Used to avoid things like: Champion.GetRotation(ctx, internal.WithMethod("POST")).
	PublicOption struct {
		apply internal.RequestOption
	}
)

// WithApiMethod is the public wrapper to change the API method name used in the a request.
// Identifier used to logging and, when implemented, rate limiting.
func WithApiMethod(method string) PublicOption {
	return PublicOption{
		apply: internal.WithApiMethod(method),
	}
}

// toRequestOptions converts a slice of public options to RequestOptions.
func toRequestOptions(opts []PublicOption) []internal.RequestOption {
	out := make([]internal.RequestOption, len(opts))
	for i, o := range opts {
		out[i] = o.apply
	}
	return out
}

// MergeOptions applies the default options to the request and add the public ones.
// Public ones are applied latter, so they will override the default.
func MergeOptions(defaultOpts []internal.RequestOption, publicOpts []PublicOption) []internal.RequestOption {
	return append(defaultOpts, toRequestOptions(publicOpts)...)
}
