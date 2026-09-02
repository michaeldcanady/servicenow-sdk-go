// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package oauth2

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// --------------------
// NewServer / GetAddr
// --------------------

func TestNewServer(t *testing.T) {
	tests := []struct {
		name  string
		state string
		port  int
	}{
		{
			name:  "random port",
			state: "state1",
			port:  0,
		},
		{
			name:  "another random port with empty state",
			state: "",
			port:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := NewServer(tt.state, tt.port)
			require.NoError(t, err)
			require.NotNil(t, s)
			require.NotEmpty(t, s.Addr)

			parsed, err := url.Parse(s.Addr)
			require.NoError(t, err)
			require.Equal(t, "http", parsed.Scheme)
			require.NotEmpty(t, parsed.Port())

			require.Equal(t, s.Addr, s.GetAddr())

			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			require.NoError(t, s.Shutdown(ctx))
		})
	}
}

func TestServer_GetAddr_BeforeAndAfterStart(t *testing.T) {
	// Before NewServer is called there's no server instance, but GetAddr on a
	// zero-value Server (never started) should simply return the empty Addr field.
	zero := &Server{}
	require.Empty(t, zero.GetAddr())

	s, err := NewServer("state", 0)
	require.NoError(t, err)
	require.NotEmpty(t, s.GetAddr())

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	require.NoError(t, s.Shutdown(ctx))
}

// --------------------
// Shutdown
// --------------------

func TestServer_Shutdown(t *testing.T) {
	tests := []struct {
		name   string
		server func(t *testing.T) *Server
	}{
		{
			name: "nil underlying server is a no-op",
			server: func(_ *testing.T) *Server {
				return &Server{}
			},
		},
		{
			name: "running server shuts down cleanly",
			server: func(t *testing.T) *Server {
				s, err := NewServer("state", 0)
				require.NoError(t, err)
				return s
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.server(t)
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			err := s.Shutdown(ctx)
			require.NoError(t, err)
		})
	}
}

// --------------------
// Result
// --------------------

func TestServer_Result(t *testing.T) {
	t.Run("returns delivered result", func(t *testing.T) {
		s := &Server{result: make(chan AuthorizationResult, 1)}
		s.result <- AuthorizationResult{Code: "abc", State: "st"}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		res := s.Result(ctx)
		require.NoError(t, res.Err)
		require.Equal(t, "abc", res.Code)
		require.Equal(t, "st", res.State)
	})

	t.Run("context cancellation returns context error", func(t *testing.T) {
		s := &Server{result: make(chan AuthorizationResult)}

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		res := s.Result(ctx)
		require.Error(t, res.Err)
		require.ErrorIs(t, res.Err, context.Canceled)
	})

	t.Run("context timeout returns deadline exceeded", func(t *testing.T) {
		s := &Server{result: make(chan AuthorizationResult)}

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
		defer cancel()

		res := s.Result(ctx)
		require.Error(t, res.Err)
		require.ErrorIs(t, res.Err, context.DeadlineExceeded)
	})
}

// --------------------
// handleCallback
// --------------------

func TestServer_handleCallback(t *testing.T) {
	tests := []struct {
		name          string
		state         string
		query         url.Values
		wantErrString bool
		wantStatus    int
		wantResult    AuthorizationResult
	}{
		{
			name:  "successful callback delivers code and state",
			state: "expected-state",
			query: url.Values{
				CodeKey:  []string{"auth-code"},
				StateKey: []string{"expected-state"},
			},
			wantStatus: http.StatusOK,
			wantResult: AuthorizationResult{Code: "auth-code", State: "expected-state"},
		},
		{
			name:  "provider error with description",
			state: "",
			query: url.Values{
				ErrorKey:            []string{"access_denied"},
				ErrorDescriptionKey: []string{"user cancelled"},
			},
			wantErrString: true,
			wantStatus:    http.StatusBadRequest,
		},
		{
			name:  "provider error without description",
			state: "",
			query: url.Values{
				ErrorKey: []string{"access_denied"},
			},
			wantErrString: true,
			wantStatus:    http.StatusBadRequest,
		},
		{
			name:          "missing code",
			state:         "",
			query:         url.Values{},
			wantErrString: true,
			wantStatus:    http.StatusBadRequest,
		},
		{
			name:  "state mismatch - missing state param",
			state: "expected-state",
			query: url.Values{
				CodeKey: []string{"auth-code"},
			},
			wantErrString: true,
			wantStatus:    http.StatusBadRequest,
		},
		{
			name:  "state mismatch - wrong state param",
			state: "expected-state",
			query: url.Values{
				CodeKey:  []string{"auth-code"},
				StateKey: []string{"wrong-state"},
			},
			wantErrString: true,
			wantStatus:    http.StatusBadRequest,
		},
		{
			name:  "success with no expected state configured",
			state: "",
			query: url.Values{
				CodeKey: []string{"auth-code"},
			},
			wantStatus: http.StatusOK,
			wantResult: AuthorizationResult{Code: "auth-code"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				result: make(chan AuthorizationResult, 1),
				state:  tt.state,
			}

			req := httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/callback?"+tt.query.Encode(), nil)
			rec := httptest.NewRecorder()

			s.handleCallback(rec, req)

			resp := rec.Result()
			defer resp.Body.Close()
			require.Equal(t, tt.wantStatus, resp.StatusCode)
			require.Equal(t, "text/html; charset=utf-8", resp.Header.Get(ContentTypeKey))

			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			require.NotEmpty(t, body)

			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			res := s.Result(ctx)

			if tt.wantErrString {
				require.Error(t, res.Err)
				require.Contains(t, string(body), "Authentication Failed")
				return
			}

			require.NoError(t, res.Err)
			require.Equal(t, tt.wantResult.Code, res.Code)
			require.Equal(t, tt.wantResult.State, res.State)
			require.Contains(t, string(body), "Authentication Successful")
		})
	}
}

// --------------------
// resultCh
// --------------------

func TestServer_resultCh(t *testing.T) {
	t.Run("delivers when channel has room", func(t *testing.T) {
		s := &Server{result: make(chan AuthorizationResult, 1)}
		s.resultCh(AuthorizationResult{Code: "delivered"})

		select {
		case res := <-s.result:
			require.Equal(t, "delivered", res.Code)
		default:
			t.Fatal("expected a result to be delivered")
		}
	})

	t.Run("does not block when channel already full", func(t *testing.T) {
		s := &Server{result: make(chan AuthorizationResult, 1)}
		s.result <- AuthorizationResult{Code: "first"}

		done := make(chan struct{})
		go func() {
			s.resultCh(AuthorizationResult{Code: "second"})
			close(done)
		}()

		select {
		case <-done:
		case <-time.After(2 * time.Second):
			t.Fatal("resultCh blocked when channel was full")
		}

		res := <-s.result
		require.Equal(t, "first", res.Code)
	})
}

// --------------------
// writeResponse
// --------------------

func TestServer_writeResponse(t *testing.T) {
	tests := []struct {
		name       string
		title      string
		message    string
		isError    bool
		wantStatus int
		wantColor  string
	}{
		{
			name:       "success response",
			title:      "Authentication Successful",
			message:    "all good",
			isError:    false,
			wantStatus: http.StatusOK,
			wantColor:  "#28a745",
		},
		{
			name:       "error response",
			title:      "Authentication Failed",
			message:    "bad state",
			isError:    true,
			wantStatus: http.StatusBadRequest,
			wantColor:  "#dc3545",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{}
			rec := httptest.NewRecorder()

			s.writeResponse(rec, tt.title, tt.message, tt.isError)

			resp := rec.Result()
			defer resp.Body.Close()
			require.Equal(t, tt.wantStatus, resp.StatusCode)
			require.Equal(t, "text/html; charset=utf-8", resp.Header.Get(ContentTypeKey))

			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			bodyStr := string(body)
			require.Contains(t, bodyStr, tt.title)
			require.Contains(t, bodyStr, tt.message)
			require.Contains(t, bodyStr, tt.wantColor)
		})
	}
}
