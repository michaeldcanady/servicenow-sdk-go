package authentication

import (
	"context"
	"fmt"
	"html"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var convertFragmentToQueryStringPage = []byte(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <script>
        window.onload = function() {
            // Parse the URL hash to get the access token
            const hash = window.location.hash;
            const params = new URLSearchParams(hash.substring(1));
            const accessToken = params.get('access_token');
            const state = params.get('state');
            const expiresIn = params.get('expires_in');

            var newUrl = window.location.origin + window.location.pathname + "?" + params;

            // Resend the request with the new URL
            window.location.href = newUrl;
        };
    </script>
</head>
<body>
</body>
</html>
`)

// Result is the result from the redirect.
type Result struct {
	// AccessToken is the code sent by the authority server.
	AccessToken string
	//
	ExpiresIn string
	// Err is set if there was an error.
	Err error
}

// authenticationTokenRedirectServer is an HTTP server.
type authenticationTokenRedirectServer struct {
	// Addr is the address the server is listening on.
	Addr      string
	resultCh  chan Result
	s         *http.Server
	reqState  string
	callCount int
	mu        sync.Mutex
}

// NewAuthenticationTokenRedirectServer creates a local HTTP server and starts it.
func NewAuthenticationTokenRedirectServer(reqState string, port int) (*authenticationTokenRedirectServer, error) {
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

	serv := &authenticationTokenRedirectServer{
		Addr:     fmt.Sprintf("http://localhost:%s", portStr),
		s:        &http.Server{Addr: "localhost:0", ReadHeaderTimeout: time.Second},
		reqState: reqState,
		resultCh: make(chan Result, 1),
	}
	serv.s.Handler = http.HandlerFunc(serv.handler)

	if err := serv.start(l); err != nil {
		return nil, err
	}

	return serv, nil
}

func (s *authenticationTokenRedirectServer) start(l net.Listener) error {
	go func() {
		err := s.s.Serve(l)
		if err != nil {
			select {
			case s.resultCh <- Result{Err: err}:
			default:
			}
		}
	}()

	return nil
}

// Result gets the result of the redirect operation. Once a single result is returned, the server
// is shutdown. ctx deadline will be honored.
func (s *authenticationTokenRedirectServer) Result(ctx context.Context) Result {
	select {
	case <-ctx.Done():
		return Result{Err: ctx.Err()}
	case r := <-s.resultCh:
		return r
	}
}

// Shutdown shuts down the server.
func (s *authenticationTokenRedirectServer) Shutdown() {
	// Note: You might get clever and think you can do this in handler() as a defer, you can't.
	_ = s.s.Shutdown(context.Background())
}

func (s *authenticationTokenRedirectServer) putResult(r Result) {
	select {
	case s.resultCh <- r:
	default:
	}
}

func (s *authenticationTokenRedirectServer) handler(w http.ResponseWriter, r *http.Request) {
	// TODO: service now sends auth info as fragment but fragments aren't supported by golang... -_-
	//https://stackoverflow.com/questions/25489843/http-server-get-url-fragment
	q := r.URL.Query()

	s.mu.Lock()
	s.callCount++
	callCount := s.callCount
	s.mu.Unlock()

	// initial request
	if len(q) == 0 {
		_, _ = w.Write(convertFragmentToQueryStringPage)
		return
	}

	if callCount < 2 {
		_, _ = w.Write(convertFragmentToQueryStringPage)
		return
	}

	headerErr := q.Get("error")
	if headerErr != "" {
		desc := html.EscapeString(q.Get("error_description"))
		// Note: It is a little weird we handle some errors by not going to the failPage. If they all should,
		// change this to s.error() and make s.error() write the failPage instead of an error code.
		_, _ = w.Write([]byte(fmt.Sprintf(failPage, headerErr, desc)))
		s.putResult(Result{Err: fmt.Errorf(desc)})
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

	token := q.Get("access_token")
	if token == "" {
		s.error(w, http.StatusInternalServerError, "authorization code missing in query string")
		return
	}

	expiresIn := q.Get("expires_in")
	if expiresIn == "" {
		s.error(w, http.StatusInternalServerError, "expires in missing in query string")
		return
	}

	_, _ = w.Write(okPage)
	s.putResult(Result{AccessToken: token, ExpiresIn: expiresIn})
}

func (s *authenticationTokenRedirectServer) error(w http.ResponseWriter, code int, str string, i ...interface{}) {
	err := fmt.Errorf(str, i...)
	http.Error(w, err.Error(), code)
	s.putResult(Result{Err: err})
}
