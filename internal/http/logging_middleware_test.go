package internalhttp

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	nethttplibrary "github.com/microsoft/kiota-http-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type captureLogger struct {
	mu       sync.Mutex
	messages []string
}

func (l *captureLogger) Log(message string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.messages = append(l.messages, fmt.Sprintf(message, args...))
}

func (l *captureLogger) output() string {
	l.mu.Lock()
	defer l.mu.Unlock()

	return strings.Join(l.messages, "\n")
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

type stubPipeline struct {
	transport http.RoundTripper
}

func (p *stubPipeline) Next(req *http.Request, _ int) (*http.Response, error) {
	return p.transport.RoundTrip(req)
}

func newResponse(status int) *http.Response {
	return &http.Response{StatusCode: status, Header: http.Header{}, Body: http.NoBody}
}

func testRequest(t *testing.T, target string, body string) *http.Request {
	t.Helper()

	var reader io.Reader
	if body != "" {
		reader = strings.NewReader(body)
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, target, reader)
	require.NoError(t, err)

	return req
}

func TestLoggingMiddleware_Intercept(t *testing.T) {
	tests := []struct {
		name     string
		logger   func() internal.Logger
		request  func(*testing.T) *http.Request
		pipeline nethttplibrary.Pipeline
		verify   func(*testing.T, internal.Logger, *http.Response, error)
	}{
		{
			name: "successful response logged at DEBUG and INFO",
			request: func(t *testing.T) *http.Request {
				return testRequest(t, "https://instance.service-now.com/api/now/v1/table/incident", "")
			},
			pipeline: &stubPipeline{transport: roundTripperFunc(func(_ *http.Request) (*http.Response, error) {
				return newResponse(http.StatusOK), nil
			})},
			verify: func(t *testing.T, logger internal.Logger, resp *http.Response, err error) {
				require.NoError(t, err)
				require.NotNil(t, resp)

				messages := logger.(*captureLogger).messages
				require.Len(t, messages, 2)
				assert.True(t, strings.HasPrefix(messages[0], "DEBUG "))
				assert.Contains(t, messages[0], "GET https://instance.service-now.com/api/now/v1/table/incident")
				assert.Contains(t, messages[0], "(attempt 1)")
				assert.True(t, strings.HasPrefix(messages[1], "INFO "))
				assert.Contains(t, messages[1], "-> 200")
				assert.Contains(t, messages[1], "attempt 1")
			},
		},
		{
			name: "4xx and 5xx responses are logged at WARN",
			request: func(t *testing.T) *http.Request {
				return testRequest(t, "https://instance.service-now.com/api/now/v1/table/incident", "")
			},
			pipeline: &stubPipeline{transport: roundTripperFunc(func(_ *http.Request) (*http.Response, error) {
				return newResponse(http.StatusNotFound), nil
			})},
			verify: func(t *testing.T, logger internal.Logger, resp *http.Response, err error) {
				require.NoError(t, err)
				require.NotNil(t, resp)

				messages := logger.(*captureLogger).messages
				require.Len(t, messages, 2)
				assert.True(t, strings.HasPrefix(messages[1], "WARN "))
				assert.Contains(t, messages[1], "-> 404")
			},
		},
		{
			name: "transport errors are logged at ERROR",
			request: func(t *testing.T) *http.Request {
				return testRequest(t, "https://instance.service-now.com/api/now/v1/table/incident", "")
			},
			pipeline: &stubPipeline{transport: roundTripperFunc(func(_ *http.Request) (*http.Response, error) {
				return nil, &url.Error{
					Op:  "Get",
					URL: "https://instance.service-now.com/api/now/v1/table/incident?user_token=supersecrettoken",
					Err: errors.New("connection refused"),
				}
			})},
			verify: func(t *testing.T, logger internal.Logger, resp *http.Response, err error) {
				require.Error(t, err)
				assert.Nil(t, resp)

				messages := logger.(*captureLogger).messages
				require.Len(t, messages, 2)
				assert.True(t, strings.HasPrefix(messages[1], "ERROR "))
				assert.Contains(t, messages[1], "failed after ")
				assert.Contains(t, messages[1], "connection refused")
				assert.NotContains(t, messages[1], "supersecrettoken")
			},
		},
		{
			name: "query strings and userinfo are never logged",
			request: func(t *testing.T) *http.Request {
				return testRequest(
					t,
					"https://user:password123@instance.service-now.com/api/now/v1/table/incident?sysparm_query=secretquery&user_token=tok123456",
					"",
				)
			},
			pipeline: &stubPipeline{transport: roundTripperFunc(func(_ *http.Request) (*http.Response, error) {
				return newResponse(http.StatusOK), nil
			})},
			verify: func(t *testing.T, logger internal.Logger, _ *http.Response, _ error) {
				output := logger.(*captureLogger).output()
				assert.Contains(t, output, "/api/now/v1/table/incident")
				assert.NotContains(t, output, "secretquery")
				assert.NotContains(t, output, "tok123456")
				assert.NotContains(t, output, "password123")
				assert.NotContains(t, output, "?")
			},
		},
		{
			name: "authorization headers and bodies are never logged",
			request: func(t *testing.T) *http.Request {
				req := testRequest(t, "https://instance.service-now.com/api/now/v1/table/incident", `{"password":"hunter2"}`)
				req.Header.Set("Authorization", "Bearer supersecretbearertoken")
				req.Header.Set("Cookie", "sessionid=supersecretcookie")

				return req
			},
			pipeline: &stubPipeline{transport: roundTripperFunc(func(_ *http.Request) (*http.Response, error) {
				resp := newResponse(http.StatusOK)
				resp.Body = io.NopCloser(strings.NewReader(`{"secret":"responsebodysecret"}`))

				return resp, nil
			})},
			verify: func(t *testing.T, logger internal.Logger, _ *http.Response, _ error) {
				output := logger.(*captureLogger).output()
				assert.NotContains(t, output, "supersecretbearertoken")
				assert.NotContains(t, output, "supersecretcookie")
				assert.NotContains(t, output, "hunter2")
				assert.NotContains(t, output, "responsebodysecret")
				assert.NotContains(t, output, "Bearer")
			},
		},
		{
			name: "nil request passes through untouched",
			request: func(*testing.T) *http.Request {
				return nil
			},
			pipeline: &stubPipeline{transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
				assert.Nil(t, req)

				return newResponse(http.StatusOK), nil
			})},
			verify: func(t *testing.T, logger internal.Logger, resp *http.Response, err error) {
				require.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Empty(t, logger.(*captureLogger).messages)
			},
		},
		{
			name: "nil logger passes through without logging",
			logger: func() internal.Logger {
				var typedNil *captureLogger

				return typedNil
			},
			request: func(t *testing.T) *http.Request {
				return testRequest(t, "https://instance.service-now.com/api/now/v1/table/incident", "")
			},
			pipeline: &stubPipeline{transport: roundTripperFunc(func(_ *http.Request) (*http.Response, error) {
				return newResponse(http.StatusOK), nil
			})},
			verify: func(t *testing.T, _ internal.Logger, resp *http.Response, err error) {
				require.NoError(t, err)
				assert.NotNil(t, resp)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var logger internal.Logger = &captureLogger{}
			if test.logger != nil {
				logger = test.logger()
			}

			middleware := NewLoggingMiddleware(logger)
			resp, err := middleware.Intercept(test.pipeline, 0, test.request(t))
			test.verify(t, logger, resp, err)
			if resp != nil {
				defer resp.Body.Close()
			}
		})
	}
}

func TestLoggingMiddleware_NilReceiverPassesThrough(t *testing.T) {
	var middleware *LoggingMiddleware

	resp, err := middleware.Intercept(&stubPipeline{
		transport: roundTripperFunc(func(_ *http.Request) (*http.Response, error) {
			return newResponse(http.StatusOK), nil
		}),
	}, 0, testRequest(t, "https://instance.service-now.com/api/now/v1/table/incident", ""))

	require.NoError(t, err)
	require.NotNil(t, resp)
	defer resp.Body.Close()
	assert.NotNil(t, resp)
}

func TestNewLoggingMiddleware(t *testing.T) {
	tests := []struct {
		name   string
		logger internal.Logger
		verify func(*testing.T, *LoggingMiddleware)
	}{
		{
			name:   "with logger",
			logger: &internal.NoOpLogger{},
			verify: func(t *testing.T, middleware *LoggingMiddleware) {
				require.NotNil(t, middleware)
				assert.IsType(t, &internal.NoOpLogger{}, middleware.logger)
			},
		},
		{
			name:   "nil logger is tolerated at construction",
			logger: nil,
			verify: func(t *testing.T, middleware *LoggingMiddleware) {
				require.NotNil(t, middleware)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.verify(t, NewLoggingMiddleware(test.logger))
		})
	}
}

func TestLoggingMiddleware_AttemptCounting(t *testing.T) {
	t.Run("retries share the counter and increment per attempt", func(t *testing.T) {
		logger := &captureLogger{}
		middleware := NewLoggingMiddleware(logger)
		calls := 0

		handler := &handlerStub{
			followUps: 1,
			inner: innerPipeline{
				middleware: middleware,
				transport: roundTripperFunc(func(_ *http.Request) (*http.Response, error) {
					calls++
					if calls == 1 {
						return newResponse(http.StatusServiceUnavailable), nil
					}

					return newResponse(http.StatusOK), nil
				}),
			},
		}

		resp, err := handler.Next(testRequest(t, "https://instance.service-now.com/api/now/v1/table/incident", ""), 0)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		defer resp.Body.Close()
		assert.Equal(t, 2, calls)

		messages := logger.messages
		require.Len(t, messages, 4)
		assert.Contains(t, messages[0], "(attempt 1)")
		assert.True(t, strings.HasPrefix(messages[1], "WARN "))
		assert.Contains(t, messages[1], "-> 503")
		assert.True(t, strings.HasPrefix(messages[2], "DEBUG "))
		assert.Contains(t, messages[2], "attempt 2")
		assert.True(t, strings.HasPrefix(messages[3], "INFO "))
		assert.Contains(t, messages[3], "-> 200")
	})

	t.Run("redirect hops are logged individually", func(t *testing.T) {
		logger := &captureLogger{}
		middleware := NewLoggingMiddleware(logger)

		handler := &handlerStub{
			followUps: 1,
			redirect:  true,
			inner: innerPipeline{
				middleware: middleware,
				transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
					if req.URL.Path == "/api/now/v1/table/incident" {
						return newResponse(http.StatusMovedPermanently), nil
					}

					return newResponse(http.StatusOK), nil
				}),
			},
		}

		resp, err := handler.Next(testRequest(t, "https://instance.service-now.com/api/now/v1/table/incident", ""), 0)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		defer resp.Body.Close()

		messages := logger.messages
		require.Len(t, messages, 4)
		assert.Contains(t, messages[0], "(attempt 1)")
		assert.Contains(t, messages[0], "/api/now/v1/table/incident")
		assert.Contains(t, messages[1], "-> 301")
		assert.True(t, strings.HasPrefix(messages[2], "DEBUG "))
		assert.Contains(t, messages[2], "/api/now/v1/table/redirected")
		assert.True(t, strings.HasPrefix(messages[3], "INFO "))
		assert.Contains(t, messages[3], "-> 200")
	})
}

// innerPipeline descends from a simulated outer handler through the logging
// middleware down to the transport, honoring kiota's middlewareIndex contract.
type innerPipeline struct {
	middleware *LoggingMiddleware
	transport  http.RoundTripper
}

func (p innerPipeline) Next(req *http.Request, middlewareIndex int) (*http.Response, error) {
	if middlewareIndex > 0 {
		return p.transport.RoundTrip(req)
	}

	return p.middleware.Intercept(p, middlewareIndex+1, req)
}

// handlerStub mimics kiota's retry/redirect handlers: they sit above the
// middleware and re-invoke the inner pipeline once the previous attempt has
// fully completed. Like kiota's RetryHandler, retries mark the shared request
// with a Retry-Attempt header.
type handlerStub struct {
	inner     innerPipeline
	redirect  bool
	followUps int
	done      int
}

func (p *handlerStub) Next(req *http.Request, _ int) (*http.Response, error) {
	resp, err := p.inner.Next(req, 0)
	if err != nil || p.done >= p.followUps {
		return resp, err
	}
	p.done++

	nextReq := req
	if p.redirect {
		nextReq = req.Clone(req.Context())
		nextReq.URL = &url.URL{Scheme: "https", Host: req.URL.Host, Path: "/api/now/v1/table/redirected"}
	} else {
		nextReq.Header.Set(retryAttemptHeader, strconv.Itoa(p.done))
	}

	return p.inner.Next(nextReq, 0)
}

func TestSanitizeURL(t *testing.T) {
	tests := []struct {
		name  string
		input *url.URL
		want  string
	}{
		{
			name:  "nil URL yields empty string",
			input: nil,
			want:  "",
		},
		{
			name:  "plain URL unchanged",
			input: mustParseURL(t, "https://instance.service-now.com/api/now/v1/table"),
			want:  "https://instance.service-now.com/api/now/v1/table",
		},
		{
			name:  "query string stripped",
			input: mustParseURL(t, "https://instance.service-now.com/api/now/v1/table?sysparm_query=secret"),
			want:  "https://instance.service-now.com/api/now/v1/table",
		},
		{
			name:  "userinfo stripped",
			input: mustParseURL(t, "https://admin:hunter2@instance.service-now.com/api/now/v1/table"),
			want:  "https://instance.service-now.com/api/now/v1/table",
		},
		{
			name:  "fragment stripped",
			input: mustParseURL(t, "https://instance.service-now.com/api/now/v1/table#fragment"),
			want:  "https://instance.service-now.com/api/now/v1/table",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, sanitizeURL(test.input))
		})
	}
}

func mustParseURL(t *testing.T, raw string) *url.URL {
	t.Helper()

	parsed, err := url.Parse(raw)
	require.NoError(t, err)

	return parsed
}
