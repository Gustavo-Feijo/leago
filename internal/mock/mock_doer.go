package mock

import (
	"io"
	"net/http"
	"strings"
)

type Doer struct {
	CapturedReq *http.Request
	Response    *http.Response
	Err         error
}

func (m *Doer) Do(req *http.Request) (*http.Response, error) {
	m.CapturedReq = req
	return m.Response, m.Err
}

func NewDefaultDoer(statusCode int, body string, err error) *Doer {
	return &Doer{
		Response: &http.Response{
			StatusCode: statusCode,
			Body:       io.NopCloser(strings.NewReader(body)),
		},
		Err: err,
	}
}
