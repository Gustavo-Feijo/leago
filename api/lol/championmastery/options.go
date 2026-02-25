package championmastery

import (
	"leago/internal"
	"strconv"
)

type GetByPUUIDTopOption internal.RequestOption

// WithCount applies the count param to the internal request.
func WithCount(count int) GetByPUUIDTopOption {
	return GetByPUUIDTopOption(internal.WithParam("count", strconv.Itoa(count)))
}

// puuidTopOptionsToRequestOptions converts the array of options into internal request options.
func puuidTopOptionsToRequestOptions(opts []GetByPUUIDTopOption) []internal.RequestOption {
	out := make([]internal.RequestOption, len(opts))
	for i, o := range opts {
		out[i] = internal.RequestOption(o)
	}
	return out
}
