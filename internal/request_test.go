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

/*func TestAuthRequest(t *testing.T) {
	tests := []struct {
		name string

		url     string
		reqOpts []RequestOption

		setupLimiterMock func(*rlmock.MockLimiter)

		httpStatusCode int
		httpBody       io.ReadCloser
		httpHeaders    http.Header
		httpErr        error

		wantName       string
		wantTokenParam string
		wantErr        bool
		wantRiotErr    bool
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
			name: "invalid post request body",

			url: "http://testexample.com",
			reqOpts: []RequestOption{
				WithHTTPMethod(http.MethodPost),
				WithBody(make(chan int)), // json.Marshal will fail with channels.
			},

			wantErr: true,
		},
		{
			name: "request failed",

			url: "http://testexample.com",

			reqOpts: []RequestOption{
				WithAPIMethod("method"),
			},

			setupLimiterMock: func(ml *rlmock.MockLimiter) {
				ml.On("Acquire", mock.Anything, "app:test", "method:test:method").Return(nil)
			},

			httpErr: &http.MaxBytesError{},

			wantErr: true,
		},
		{
			name: "io reader failed",

			url: "http://testexample.com",

			reqOpts: []RequestOption{
				WithAPIMethod("method"),
			},

			setupLimiterMock: func(ml *rlmock.MockLimiter) {
				ml.On("Acquire", mock.Anything, "app:test", "method:test:method").Return(nil)
			},

			httpBody: io.NopCloser(errorReader{}),

			wantErr: true,
		},
		{
			name: "non ok http status code",

			url: "http://testexample.com",

			reqOpts: []RequestOption{
				WithAPIMethod("method"),
			},

			setupLimiterMock: func(ml *rlmock.MockLimiter) {
				ml.On("Acquire", mock.Anything, "app:test", "method:test:method").Return(nil)
			},

			httpStatusCode: 400,
			httpBody:       io.NopCloser(strings.NewReader("Error")),

			wantErr:     true,
			wantRiotErr: true,
		},
		{
			name: "notify 429 status code",

			url: "http://testexample.com",

			reqOpts: []RequestOption{
				WithAPIMethod("method"),
			},

			setupLimiterMock: func(ml *rlmock.MockLimiter) {
				ml.On("Acquire", mock.Anything, "app:test", "method:test:method").Return(nil)
				ml.On("Sync", mock.Anything, mock.Anything, "app:test", "method:test:method").Return()
				ml.On("NotifyTooManyRequests", mock.Anything, mock.Anything, "app:test", "method:test:method").Return()
			},

			httpStatusCode: 429,
			httpBody:       io.NopCloser(strings.NewReader("Error")),

			wantErr:     true,
			wantRiotErr: true,
		},
		{
			name: "json unmarshal error",

			url: "http://testexample.com",

			reqOpts: []RequestOption{
				WithAPIMethod("method"),
			},

			setupLimiterMock: func(ml *rlmock.MockLimiter) {
				ml.On("Acquire", mock.Anything, "app:test", "method:test:method").Return(nil)
			},

			httpStatusCode: 200,
			httpBody:       io.NopCloser(strings.NewReader("invalid json")),

			wantErr:     true,
			wantRiotErr: false,
		},
		{
			name: "limiter error",

			url: "http://testexample.com",

			reqOpts: []RequestOption{
				WithAPIMethod("method"),
			},

			setupLimiterMock: func(ml *rlmock.MockLimiter) {
				ml.On("Acquire", mock.Anything, "app:test", "method:test:method").Return(errors.New("err"))
			},

			wantErr:     true,
			wantRiotErr: false,
		},

		{
			name: "success",

			url: "http://testexample.com",

			reqOpts: []RequestOption{
				WithAPIMethod("method"),
			},

			setupLimiterMock: func(ml *rlmock.MockLimiter) {
				ml.On("Acquire", mock.Anything, "app:test", "method:test:method").Return(nil)
			},

			httpStatusCode: 200,
			httpBody:       io.NopCloser(strings.NewReader(`{"name":"valid name"}`)),

			wantName: "valid name",
		},
		{
			name: "limiter sync",

			url: "http://testexample.com",

			reqOpts: []RequestOption{
				WithAPIMethod("method"),
			},

			setupLimiterMock: func(ml *rlmock.MockLimiter) {
				ml.On("Acquire", mock.Anything, "app:test", "method:test:method").Return(nil)
				ml.On("Sync", mock.Anything, mock.Anything, "app:test", "method:test:method").Return()
			},

			httpStatusCode: 200,
			httpBody:       io.NopCloser(strings.NewReader(`{"name":"valid name"}`)),
			httpHeaders:    http.Header{},

			wantName: "valid name",
		},
		{
			name: "success post request",

			url: "http://testexample.com",
			reqOpts: []RequestOption{
				WithHTTPMethod(http.MethodPost),
				WithBody(PostRequest{
					Name: "posttest",
				}),
				WithAPIMethod("Test.Create"),
			},

			setupLimiterMock: func(ml *rlmock.MockLimiter) {
				ml.On("Acquire", mock.Anything, "app:test", "method:test:Test.Create").Return(nil)
			},

			httpStatusCode: 204,
			httpBody:       io.NopCloser(strings.NewReader(`{"name":"posttest"}`)),

			wantName: "posttest",
		},
		{
			name: "success with params",

			url: "http://testexample.com",
			reqOpts: []RequestOption{
				// Override.
				WithParam(apiTokenHeader, "validKey1"),
				WithParams(map[string]string{apiTokenHeader: "invalidKey"}),
				WithParam(apiTokenHeader, "validKey"),
				WithAPIMethod("method"),
			},

			setupLimiterMock: func(ml *rlmock.MockLimiter) {
				ml.On("Acquire", mock.Anything, "app:test", "method:test:method").Return(nil)
			},

			httpStatusCode: 200,
			httpBody:       io.NopCloser(strings.NewReader(`{"name":"valid name"}`)),

			wantName: "valid name",
			// Token can also be as param.
			wantTokenParam: "validKey",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDoer := &mockDoer.Doer{
				Response: &http.Response{
					StatusCode: tt.httpStatusCode,
					Body:       tt.httpBody,
					Header:     tt.httpHeaders,
				},
				Err: tt.httpErr,
			}

			limiter := &rlmock.MockLimiter{}

			if tt.setupLimiterMock != nil {
				tt.setupLimiterMock(limiter)
			}

			client := &Client{
				HTTP:        mockDoer,
				limiter:     limiter,
				Logger:      slog.Default(),
				routePrefix: "test",
				apiKey:      "apiKey",
			}

			got, err := AuthRequest[Response](context.Background(), client, tt.url, tt.reqOpts...)

			if tt.wantErr {
				require.Error(t, err)

				if tt.wantRiotErr {
					var rErr *RiotError
					assert.ErrorAs(t, err, &rErr)
					assert.Equal(t, tt.httpStatusCode, rErr.StatusCode)
				}

				return
			}

			require.NoError(t, err)

			if tt.wantTokenParam != "" {
				assert.Equal(t, tt.wantTokenParam, mockDoer.CapturedReq.URL.Query().Get(apiTokenHeader))
			}

			assert.Equal(t, tt.wantName, got.Name)
		})
	}
}*/

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
