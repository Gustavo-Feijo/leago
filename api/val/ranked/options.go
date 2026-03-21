package ranked

import (
	"strconv"

	"github.com/Gustavo-Feijo/leago/internal"
)

type RankedOption internal.RequestOption

// WithSize applies size filter to the request. Default 200, valid from 1 to 200.
func WithSize(size int) RankedOption {
	return RankedOption(internal.WithParam("size", strconv.Itoa(size)))
}

// WithStartIndex applies start index filter to the request. Default 0.
func WithStartIndex(index int) RankedOption {
	return RankedOption(internal.WithParam("startIndex", strconv.Itoa(index)))
}

// rankedOptionsToRequestOptions converts the array of options into internal request options.
func rankedOptionsToRequestOptions(opts []RankedOption) []internal.RequestOption {
	out := make([]internal.RequestOption, len(opts))
	for i, o := range opts {
		out[i] = internal.RequestOption(o)
	}
	return out
}
