package matches

import (
	"strconv"

	"github.com/Gustavo-Feijo/leago/internal"
)

type GetMatchesByPUUIDOption internal.RequestOption

// WithStartTime applies the start time filter (Epoch timestamp).
func WithStartTime(startTime int64) GetMatchesByPUUIDOption {
	return GetMatchesByPUUIDOption(internal.WithParam("startTime", strconv.FormatInt(startTime, 10)))
}

// WithEndTime applies the end time filter (Epoch timestamp).
func WithEndTime(endTime int64) GetMatchesByPUUIDOption {
	return GetMatchesByPUUIDOption(internal.WithParam("endTime", strconv.FormatInt(endTime, 10)))
}

// WithQueue applies the queue filter.
func WithQueue(queue int) GetMatchesByPUUIDOption {
	return GetMatchesByPUUIDOption(internal.WithParam("queue", strconv.Itoa(queue)))
}

// WithType applies the type filter.
func WithType(matchType string) GetMatchesByPUUIDOption {
	return GetMatchesByPUUIDOption(internal.WithParam("type", matchType))
}

// WithStart applies the start index filter. Defaults to 0.
func WithStart(start int) GetMatchesByPUUIDOption {
	return GetMatchesByPUUIDOption(internal.WithParam("start", strconv.Itoa(start)))
}

// WithCount applies the count filter. Defaults to 20. Valid values: 0 to 100.
func WithCount(count int) GetMatchesByPUUIDOption {
	return GetMatchesByPUUIDOption(internal.WithParam("count", strconv.Itoa(count)))
}

// getMatchByPUUIDOptionsToRequestOptions converts the array of options into internal request options.
func getMatchByPUUIDOptionsToRequestOptions(opts []GetMatchesByPUUIDOption) []internal.RequestOption {
	out := make([]internal.RequestOption, len(opts))
	for i, o := range opts {
		out[i] = internal.RequestOption(o)
	}
	return out
}
