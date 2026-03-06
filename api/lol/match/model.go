package match

type (
	// Timeline and match types are separated in their own files due to their size.

	Replays struct {
		Total         int      `json:"total"`
		MatchFileURLs []string `json:"matchFileURLs"`
	}
)
