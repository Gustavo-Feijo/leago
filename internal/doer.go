package internal

import "net/http"

// Doer is the minimal interface leago requires from an HTTP client.
// http.Client satisfies it.
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}
