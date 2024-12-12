package authentication

import (
	"context"
	"errors"
	"fmt"
	"html"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var okPage = []byte(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <title>Authentication Complete</title>
</head>
<body>
    <p>Authentication complete. You can return to the application. Feel free to close this browser tab.</p>
</body>
</html>
`)

const failPage = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <title>Authentication Failed</title>
</head>
<body>
	<p>Authentication failed. You can return to the application. Feel free to close this browser tab.</p>
	<p>Error details: error %s error_description: %s</p>
</body>
</html>
`

// authenticationCodeResult is the result from the redirect.
type authenticationCodeResult struct {
	// Code is the code sent by the authority server.
	Code string
	// Err is set if there was an error.
	Err error
}

// authenticationCodeRedirectServer is an HTTP server.
type authenticationCodeRedirectServer struct {
	// Addr is the address the server is listening on.
	Addr     string
	resultCh chan authenticationCodeResult
	s        *http.Server
	reqState string
}

// NewAuthenticationCodeRedirectServer creates a local HTTP server and starts it.
func NewAuthenticationCodeRedirectServer(reqState string, port int) (*authenticationCodeRedirectServer, error) {
	var l net.Listener
	var err error
	var portStr string
	if port > 0 {
		// use port provided by caller
		l, err = net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
		portStr = strconv.FormatInt(int64(port), 10)
	} else {
		// find a free port
		for i := 0; i < 10; i++ {
			l, err = net.Listen("tcp", "localhost:0")
			if err != nil {
				continue
			}
			addr := l.Addr().String()
			portStr = addr[strings.LastIndex(addr, ":")+1:]
			break
		}
	}
	if err != nil {
		return nil, err
	}

	serv := &authenticationCodeRedirectServer{
		Addr:     fmt.Sprintf("http://localhost:%s", portStr),
		s:        &http.Server{Addr: "localhost:0", ReadHeaderTimeout: time.Second},
		reqState: reqState,
		resultCh: make(chan authenticationCodeResult, 1),
	}
	serv.s.Handler = http.HandlerFunc(serv.handler)

	if err := serv.start(l); err != nil {
		return nil, err
	}

	return serv, nil
}

func (s *authenticationCodeRedirectServer) start(l net.Listener) error { //nolint:unparam
	go func() {
		err := s.s.Serve(l)
		if err != nil {
			select {
			case s.resultCh <- authenticationCodeResult{Err: err}:
			default:
			}
		}
	}()

	return nil
}

// Result gets the result of the redirect operation. Once a single result is returned, the server
// is shutdown. ctx deadline will be honored.
func (s *authenticationCodeRedirectServer) Result(ctx context.Context) authenticationCodeResult {
	select {
	case <-ctx.Done():
		return authenticationCodeResult{Err: ctx.Err()}
	case r := <-s.resultCh:
		return r
	}
}

// Shutdown shuts down the server.
func (s *authenticationCodeRedirectServer) Shutdown() {
	// Note: You might get clever and think you can do this in handler() as a defer, you can't.
	_ = s.s.Shutdown(context.Background())
}

func (s *authenticationCodeRedirectServer) putResult(r authenticationCodeResult) {
	select {
	case s.resultCh <- r:
	default:
	}
}

func (s *authenticationCodeRedirectServer) handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	headerErr := q.Get("error")
	if headerErr != "" {
		desc := html.EscapeString(q.Get("error_description"))
		// Note: It is a little weird we handle some errors by not going to the failPage. If they all should,
		// change this to s.error() and make s.error() write the failPage instead of an error code.
		_, _ = w.Write([]byte(fmt.Sprintf(failPage, headerErr, desc)))
		s.putResult(authenticationCodeResult{Err: errors.New(desc)})
		return
	}

	respState := q.Get("state")
	switch respState {
	case s.reqState:
	case "":
		s.error(w, http.StatusInternalServerError, "server didn't send OAuth state")
		return
	default:
		s.error(w, http.StatusInternalServerError, "mismatched OAuth state, req(%s), resp(%s)", s.reqState, respState)
		return
	}

	code := q.Get("code")
	if code == "" {
		s.error(w, http.StatusInternalServerError, "authorization code missing in query string")
		return
	}

	_, _ = w.Write(okPage)
	s.putResult(authenticationCodeResult{Code: code})
}

func (s *authenticationCodeRedirectServer) error(w http.ResponseWriter, code int, str string, i ...interface{}) {
	err := fmt.Errorf(str, i...)
	http.Error(w, err.Error(), code)
	s.putResult(authenticationCodeResult{Err: err})
}
