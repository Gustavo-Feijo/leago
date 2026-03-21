package content

import (
	"github.com/Gustavo-Feijo/leago/internal"
)

type ContentOption internal.RequestOption

// WithLocale applies locale filter to the request.
func WithLocale(locale string) ContentOption {
	return ContentOption(internal.WithParam("locale", locale))
}

// contentOptionsToRequestOptions converts the array of options into internal request options.
func contentOptionsToRequestOptions(opts []ContentOption) []internal.RequestOption {
	out := make([]internal.RequestOption, len(opts))
	for i, o := range opts {
		out[i] = internal.RequestOption(o)
	}
	return out
}
