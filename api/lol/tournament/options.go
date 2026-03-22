package tournament

import (
	"strconv"

	"github.com/Gustavo-Feijo/leago/internal"
)

type CreateCodeOption internal.RequestOption

// WithCount applies the count param to the internal request.
func WithCount(count int) CreateCodeOption {
	return CreateCodeOption(internal.WithParam("count", strconv.Itoa(count)))
}

// createCodeOptionsToRequestOptions converts the array of options into internal request options.
func createCodeOptionsToRequestOptions(opts []CreateCodeOption) []internal.RequestOption {
	out := make([]internal.RequestOption, len(opts))
	for i, o := range opts {
		out[i] = internal.RequestOption(o)
	}
	return out
}
