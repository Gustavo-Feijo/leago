package internal

import (
	"context"
	"fmt"
	"io"
	"leago/internal/mock"
	"log/slog"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type (
	errorReader struct{}

	PostRequest struct {
		Name string `json:"name"`
	}

	Response struct {
		Name string `json:"name"`
	}
)

func (e errorReader) Read(p []byte) (int, error) {
	return 0, fmt.Errorf("forced read error")
}

func newTestClient(doer *mock.Doer) *Client {
	return &Client{
		Http:        doer,
		Logger:      slog.Default(),
		routePrefix: "test",
		apiKey:      "apiKey",
	}
}

func TestAuthRequest(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		httpStatusCode int
		httpBody       io.ReadCloser
		httpErr        error
		reqOpts        []RequestOption
		wantName       string
		wantTokenParam string
		wantErr        bool
		wantRiotErr    bool
	}{
		{
			name:    "invalid URI",
			url:     "http://[::1",
			wantErr: true,
		},
		{
			name: "invalid request",
			url:  "http://testexample.com",
			reqOpts: []RequestOption{
				WithHttpMethod("invalid::{}"), // http.NewRequestWithContext validates the method tokens.
			},
			wantErr: true,
		},
		{
			name: "invalid post request body",
			url:  "http://testexample.com",
			reqOpts: []RequestOption{
				WithHttpMethod(http.MethodPost),
				WithBody(make(chan int)), // json.Marshal will fail with channels.
			},
			wantErr: true,
		},
		{
			name:    "request failed",
			url:     "http://testexample.com",
			httpErr: &http.MaxBytesError{},
			wantErr: true,
		},
		{
			name:     "io reader failed",
			url:      "http://testexample.com",
			httpBody: io.NopCloser(errorReader{}),
			wantErr:  true,
		},
		{
			name:           "non ok http status code",
			url:            "http://testexample.com",
			httpStatusCode: 400,
			httpBody:       io.NopCloser(strings.NewReader("Error")),
			wantErr:        true,
			wantRiotErr:    true,
		},
		{
			name:           "json unmarshal error",
			url:            "http://testexample.com",
			httpStatusCode: 200,
			httpBody:       io.NopCloser(strings.NewReader("invalid json")),
			wantErr:        true,
			wantRiotErr:    false,
		},
		{
			name:           "success",
			url:            "http://testexample.com",
			httpStatusCode: 200,
			httpBody:       io.NopCloser(strings.NewReader(`{"name":"valid name"}`)),
			wantName:       "valid name",
		},
		{
			name:           "success post request",
			url:            "http://testexample.com",
			httpStatusCode: 204,
			reqOpts: []RequestOption{
				WithHttpMethod(http.MethodPost),
				WithBody(PostRequest{
					Name: "posttest",
				}),
				WithApiMethod("Test.Create"),
			},
			httpBody: io.NopCloser(strings.NewReader(`{"name":"posttest"}`)),
			wantName: "posttest",
		},
		{
			name:           "success with params",
			url:            "http://testexample.com",
			httpStatusCode: 200,
			httpBody:       io.NopCloser(strings.NewReader(`{"name":"valid name"}`)),
			wantName:       "valid name",

			// Token can also be as param.
			wantTokenParam: "validKey",
			reqOpts: []RequestOption{
				WithParams(map[string]string{apiTokenHeader: "validKey"}),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDoer := &mock.Doer{
				Response: &http.Response{
					StatusCode: tt.httpStatusCode,
					Body:       tt.httpBody,
				},
				Err: tt.httpErr,
			}

			client := newTestClient(mockDoer)

			got, err := AuthRequest[Response](context.Background(), client, tt.url, tt.reqOpts...)

			if tt.wantErr {
				assert.NotNil(t, err)

				if tt.wantRiotErr {
					var rErr *RiotError
					assert.ErrorAs(t, err, &rErr)
					assert.Equal(t, tt.httpStatusCode, rErr.StatusCode)
				}

				return
			}

			require.Nil(t, err)

			if tt.wantTokenParam != "" {
				assert.Equal(t, tt.wantTokenParam, mockDoer.CapturedReq.URL.Query().Get(apiTokenHeader))
			}

			assert.Equal(t, tt.wantName, got.Name)
		})
	}
}
