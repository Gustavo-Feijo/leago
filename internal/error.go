package internal

import "fmt"

type RiotError struct {
	StatusCode int
	Status     string
	Body       string
}

func (e *RiotError) Error() string {
	if e.Body != "" {
		return fmt.Sprintf("riot api error: %s (body: %s)", e.Status, e.Body)
	}
	return fmt.Sprintf("riot api error: %s", e.Status)
}
