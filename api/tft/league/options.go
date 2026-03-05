package league

import (
	"strconv"

	"github.com/Gustavo-Feijo/leago/internal"
)

type baseOption internal.RequestOption

type UpperLeagueOption baseOption
type LeagueOption baseOption

// WithQueue applies queue filter to the request.
func WithQueueHighElo(queue Queue) UpperLeagueOption {
	return UpperLeagueOption(internal.WithParam("queue", string(queue)))
}

// WithQueue applies queue filter to the request.
func WithQueue(queue Queue) LeagueOption {
	return LeagueOption(internal.WithParam("queue", string(queue)))
}

// WithPage applies page filter to the request.
func WithPage(page int) LeagueOption {
	return LeagueOption(internal.WithParam("page", strconv.Itoa(page)))
}

// upperLeagueOptionsToRequestOptions converts the array of options into internal request options.
func upperLeagueOptionsToRequestOptions(opts []UpperLeagueOption) []internal.RequestOption {
	out := make([]internal.RequestOption, len(opts))
	for i, o := range opts {
		out[i] = internal.RequestOption(o)
	}
	return out
}

// leagueOptionsToRequestOptions converts the array of options into internal request options.
func leagueOptionsToRequestOptions(opts []LeagueOption) []internal.RequestOption {
	out := make([]internal.RequestOption, len(opts))
	for i, o := range opts {
		out[i] = internal.RequestOption(o)
	}
	return out
}
