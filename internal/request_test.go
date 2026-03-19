package internal

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"testing"

	mockDoer "github.com/Gustavo-Feijo/leago/internal/mock"
	rlmock "github.com/Gustavo-Feijo/leago/ratelimit/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestBuildRequest(t *testing.T) {
	tests := []struct {
		name string

		url     string
		reqOpts []RequestOption

		wantErr bool

		check func(*testing.T, *http.Request)
	}{
		{
			name: "invalid URI",

			url: "http://[::1",

			wantErr: true,
		},
		{
			name: "invalid request",

			url: "http://testexample.com",
			reqOpts: []RequestOption{
				WithHTTPMethod("invalid::{}"), // http.NewRequestWithContext validates the method tokens.
			},

			wantErr: true,
		},
		{
			name: "invalid body marshal",

			url: "http://test.com",
			reqOpts: []RequestOption{
				WithBody(make(chan int)),
			},

			wantErr: true,
		},
		{
			name: "success",

			url: "http://test.com",
			reqOpts: []RequestOption{
				withAPIKey("key"),
				WithAPIMethod("method"),
				WithParams(map[string]string{
					"z": "3",
				}),
				WithParam("z", "2"),
			},

			check: func(t *testing.T, req *http.Request) {
				assert.Equal(t, "2", req.URL.Query().Get("z"))
				assert.Equal(t, "key", req.Header.Get(apiTokenHeader))
			},
		},
		{
			name: "success with params and body",

			url: "http://test.com",
			reqOpts: []RequestOption{
				withAPIKey("key"),
				WithBody(PostRequest{Name: "test"}),
				WithAPIMethod("method"),
				WithHTTPMethod(http.MethodPost),
				WithParam("a", "1"),
			},

			check: func(t *testing.T, req *http.Request) {
				assert.Equal(t, "1", req.URL.Query().Get("a"))
				assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
				assert.Equal(t, "key", req.Header.Get(apiTokenHeader))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ro requestOptions
			for _, o := range tt.reqOpts {
				o(&ro)
			}
			req, err := buildRequest(context.Background(), tt.url, &ro)

			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

			if tt.check != nil {
				tt.check(t, req)
			}
		})
	}
}

func TestAuthRequest(t *testing.T) {
	tests := []struct {
		name string

		url       string
		setupMock func(*rlmock.MockLimiter)

		resp *http.Response

		wantErr     bool
		assertMocks bool
	}{
		{
			name: "success",

			url: "http://test.com",
			setupMock: func(ml *rlmock.MockLimiter) {
				ml.On("Acquire", mock.Anything, "app:test", "method:test:method").Return(nil)
			},

			resp: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"name":"ok"}`)),
			},
		},
		{
			name: "buildRequestErr",

			url:       "http://[::1",
			setupMock: func(ml *rlmock.MockLimiter) {}, // Error happens before actual call.

			wantErr: true,
		},
		{
			name: "429 full flow",

			url: "http://test.com",
			setupMock: func(ml *rlmock.MockLimiter) {
				ml.On("Acquire", mock.Anything, "app:test", "method:test:method").Return(nil)
				ml.On("NotifyTooManyRequests", mock.Anything, mock.Anything, "app:test", "method:test:method").Return()
			},

			resp: &http.Response{
				StatusCode: 429,
				Body:       io.NopCloser(strings.NewReader("err")),
			},

			wantErr:     true,
			assertMocks: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			limiter := &rlmock.MockLimiter{}
			tt.setupMock(limiter)

			doer := &mockDoer.Doer{
				Response: tt.resp,
			}

			client := &Client{
				HTTP:        doer,
				limiter:     limiter,
				Logger:      slog.Default(),
				routePrefix: "test",
				apiKey:      "apiKey",
			}

			_, err := AuthRequest[Response](
				context.Background(),
				client,
				tt.url,
				WithAPIMethod("method"),
			)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			if tt.assertMocks {
				limiter.AssertExpectations(t)
			}
		})
	}
}

func TestExecutor(t *testing.T) {
	tests := []struct {
		name        string
		setupMock   func(*rlmock.MockLimiter)
		resp        *http.Response
		wantErr     bool
		assertMocks bool
	}{
		{
			name: "limiter blocks",
			setupMock: func(ml *rlmock.MockLimiter) {
				ml.On("Acquire", mock.Anything, "app:test", "method:test:method").
					Return(errors.New("blocked"))
			},
			wantErr: true,
		},
		{
			name: "429 triggers notify",
			setupMock: func(ml *rlmock.MockLimiter) {
				ml.On("Acquire", mock.Anything, "app:test", "method:test:method").Return(nil)
				ml.On("NotifyTooManyRequests", mock.Anything, mock.Anything, "app:test", "method:test:method").Return()
			},
			resp: &http.Response{
				StatusCode: 429,
				Body:       io.NopCloser(strings.NewReader("err")),
			},
			wantErr:     true,
			assertMocks: true,
		},
		{
			name: "sync called",
			setupMock: func(ml *rlmock.MockLimiter) {
				ml.On("Acquire", mock.Anything, "app:test", "method:test:method").Return(nil)
				ml.On("Sync", mock.Anything, mock.Anything, "app:test", "method:test:method").Return()
			},
			resp: &http.Response{
				StatusCode: 200,
				Header:     http.Header{},
				Body:       io.NopCloser(strings.NewReader(`{"name":"ok"}`)),
			},
			assertMocks: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			limiter := &rlmock.MockLimiter{}

			if tt.setupMock != nil {
				tt.setupMock(limiter)
			}

			doer := &mockDoer.Doer{
				Response: tt.resp,
			}

			client := &Client{
				HTTP:        doer,
				limiter:     limiter,
				Logger:      slog.Default(),
				routePrefix: "test",
				apiKey:      "apiKey",
			}
			req, _ := http.NewRequestWithContext(context.Background(), "GET", "http://test.com", nil)

			_, err := executor[Response](context.Background(), client, req, &requestOptions{
				apiMethod: "method",
			})

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			if tt.assertMocks {
				limiter.AssertExpectations(t)
			}
		})
	}
}

func TestDo(t *testing.T) {
	tests := []struct {
		name string

		resp *http.Response

		err error

		expectedName string

		wantErr     bool
		wantRiotErr bool
	}{
		{
			name: "http error",

			err: errors.New("fail"),

			wantErr: true,
		},
		{
			name: "http error with resp",

			resp: &http.Response{
				StatusCode: 400,
				Body:       io.NopCloser(strings.NewReader(`{"err":"err"}`)),
				Header:     http.Header{},
			},

			err: errors.New("fail"),

			wantErr: true,
		},
		{
			name: "read error",

			resp: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(errorReader{}),
			},

			wantErr: true,
		},
		{
			name: "non 2xx",

			resp: &http.Response{
				StatusCode: 400,
				Status:     "400 Bad Request",
				Body:       io.NopCloser(strings.NewReader("err")),
			},

			wantErr:     true,
			wantRiotErr: true,
		},
		{
			name: "invalid json",

			resp: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader("invalid")),
			},

			wantErr: true,
		},
		{
			name: "success",

			resp: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"name":"ok"}`)),
			},

			expectedName: "ok",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doer := &mockDoer.Doer{
				Response: tt.resp,
				Err:      tt.err,
			}

			client := &Client{
				HTTP:        doer,
				limiter:     nil,
				Logger:      slog.Default(),
				routePrefix: "test",
				apiKey:      "apiKey",
			}

			req, _ := http.NewRequestWithContext(context.Background(), "GET", "http://test.com", nil)

			res, _, err := do[Response](client, req, &requestOptions{})

			if tt.wantErr {
				require.Error(t, err)

				if tt.wantRiotErr {
					var rErr *RiotError
					assert.ErrorAs(t, err, &rErr)
				}

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expectedName, res.Name)
		})
	}
}
