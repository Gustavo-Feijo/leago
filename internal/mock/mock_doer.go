package mock

import "net/http"

type MockDoer struct {
	CapturedReq *http.Request
	Response    *http.Response
	Err         error
}

func (m *MockDoer) Do(req *http.Request) (*http.Response, error) {
	m.CapturedReq = req
	return m.Response, m.Err
}
