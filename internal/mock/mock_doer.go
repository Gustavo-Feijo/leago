package mock

import "net/http"

type Doer struct {
	CapturedReq *http.Request
	Response    *http.Response
	Err         error
}

func (m *Doer) Do(req *http.Request) (*http.Response, error) {
	m.CapturedReq = req
	return m.Response, m.Err
}
