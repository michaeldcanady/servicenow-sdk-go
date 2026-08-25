package internalhttp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

type loggingLevel string

const (
	logLevelDebug loggingLevel = "DEBUG"
	logLevelInfo  loggingLevel = "INFO"
	logLevelWarn  loggingLevel = "WARN"
	logLevelError loggingLevel = "ERROR"

	// retryAttemptHeader mirrors nethttplibrary's marker, incremented by its
	// RetryHandler before every resend.
	retryAttemptHeader = "Retry-Attempt"
)

type attemptCountKey struct{}

// LoggingMiddleware logs every HTTP attempt (initial request, retries and
// redirect hops) through an internal.Logger. Only the method, a sanitized URL
// (no query string, fragment or userinfo), the status code, the duration and
// the attempt number are emitted — headers and bodies are never read.
type LoggingMiddleware struct {
	logger internal.Logger
}

// NewLoggingMiddleware creates a new LoggingMiddleware emitting through the provided logger.
func NewLoggingMiddleware(logger internal.Logger) *LoggingMiddleware {
	return &LoggingMiddleware{logger: logger}
}

// Intercept implements nethttplibrary.Middleware and must sit inside (after)
// the retry and redirect handlers so each physical attempt is observed.
// Retries are numbered via kiota's Retry-Attempt header; redirect hops are
// logged individually but restart at one, since handlers clone from a
// pre-middleware request reference.
func (m *LoggingMiddleware) Intercept(pipeline nethttplibrary.Pipeline, middlewareIndex int, req *http.Request) (*http.Response, error) {
	if m == nil || conversion.IsNil(m.logger) || req == nil {
		return pipeline.Next(req, middlewareIndex)
	}

	counter, ok := req.Context().Value(attemptCountKey{}).(*int)
	if !ok || counter == nil {
		counter = new(int)
		req = req.WithContext(context.WithValue(req.Context(), attemptCountKey{}, counter))
	}
	attempt := nextAttempt(counter)

	// kiota's RetryHandler holds its own pre-middleware request reference, so
	// context values do not survive retry loops; it marks resends with this
	// header instead.
	if retried, err := strconv.Atoi(req.Header.Get(retryAttemptHeader)); err == nil && retried+1 > attempt {
		attempt = retried + 1
	}

	target := sanitizeURL(req.URL)
	start := time.Now()
	m.log(logLevelDebug, "%s %s (attempt %d)", req.Method, target, attempt)

	resp, err := pipeline.Next(req, middlewareIndex)
	duration := time.Since(start)

	switch {
	case err != nil:
		m.log(logLevelError, "%s %s attempt %d failed after %s: %v", req.Method, target, attempt, duration, errorCause(err))
	case resp.StatusCode >= http.StatusBadRequest:
		m.log(logLevelWarn, "%s %s -> %d (%s, attempt %d)", req.Method, target, resp.StatusCode, duration, attempt)
	default:
		m.log(logLevelInfo, "%s %s -> %d (%s, attempt %d)", req.Method, target, resp.StatusCode, duration, attempt)
	}

	return resp, err
}

func (m *LoggingMiddleware) log(level loggingLevel, format string, args ...interface{}) {
	m.logger.Log(fmt.Sprintf(string(level)+" "+format, args...))
}

func nextAttempt(counter *int) int {
	*counter++
	return *counter
}

// sanitizeURL strips query strings, fragments and userinfo, since any of them
// can carry credentials (e.g. tokens in sysparm_ parameters).
func sanitizeURL(u *url.URL) string {
	if u == nil {
		return ""
	}

	sanitized := *u
	sanitized.User = nil
	sanitized.RawQuery = ""
	sanitized.RawFragment = ""
	sanitized.Fragment = ""

	return sanitized.String()
}

// errorCause unwraps *url.Error so its embedded URL — which may carry secrets
// in the query string — is never echoed into logs.
func errorCause(err error) error {
	var urlErr *url.Error
	if errors.As(err, &urlErr) && urlErr.Err != nil {
		return urlErr.Err
	}

	return err
}
