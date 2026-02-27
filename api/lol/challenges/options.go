package challenges

import (
	"leago/internal"
	"strconv"
)

type GetLeaderboardOption internal.RequestOption

// WithLimit adds a limit of entries returned.
func WithLimit(count int) GetLeaderboardOption {
	return GetLeaderboardOption(internal.WithParam("limit", strconv.Itoa(count)))
}

// getLeaderboardOptionsToRequestOptions converts the array of options into internal request options.
func getLeaderboardOptionsToRequestOptions(opts []GetLeaderboardOption) []internal.RequestOption {
	out := make([]internal.RequestOption, len(opts))
	for i, o := range opts {
		out[i] = internal.RequestOption(o)
	}
	return out
}
