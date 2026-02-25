package league

import (
	"leago/internal"
	"strconv"
)

type GetLeagueOption internal.RequestOption

// WithPage applies the count param to the internal request.
func WithPage(count int) GetLeagueOption {
	return GetLeagueOption(internal.WithParam("page", strconv.Itoa(count)))
}

// getLeagueOptionsToRequestOptions converts the array of options into internal request options.
func getLeagueOptionsToRequestOptions(opts []GetLeagueOption) []internal.RequestOption {
	out := make([]internal.RequestOption, len(opts))
	for i, o := range opts {
		out[i] = internal.RequestOption(o)
	}
	return out
}
