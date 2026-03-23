package internal

import "fmt"

// RiotError is returned for any non-2xx response from the Riot API.
// Callers can unwrap it with errors.As to inspect the status code.
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
