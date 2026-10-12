// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/browser"
)

const (
	// defaultAuthorizationCodePort is the default port for the local redirect server.
	defaultAuthorizationCodePort = 5001
	// authorizationCodeTimeout bounds the whole interactive code acquisition.
	authorizationCodeTimeout = 100 * time.Second
	// authorizationCodeShutdownTimeout bounds the redirect server shutdown.
	authorizationCodeShutdownTimeout = 5 * time.Second
)

// authorizationCodeSource is the minimum a client must provide for the
// interactive authorization code flow to run: build the authorization URL and
// exchange the returned code for a token.
type authorizationCodeSource interface {
	getAuthorizationURL(redirectURI, state string, scopes []string) (string, error)
	acquireTokenByCode(ctx context.Context, code, redirectURI, state string) (*AccessToken, error)
}

// authorizationCodeFlow runs the interactive authorization code redirect dance:
// it generates a state value, starts the local redirect server, opens the
// authorization URL, waits for the callback, and exchanges the returned code.
//
// It holds the machinery shared by every credential that authenticates through a
// browser redirect - the ServiceNow authorization code credential and the
// external OIDC credential - so the flow is implemented once.
type authorizationCodeFlow struct {
	port           int
	stateGenerator func() string
	urlOpener      func(string) error
	serverFactory  ServerFactory
	mutex          sync.Mutex
}

// newAuthorizationCodeFlow builds a flow, applying the same defaults the
// credential constructors use for an unset port, state generator, URL opener,
// or server factory.
func newAuthorizationCodeFlow(port int, stateGenerator func() string, urlOpener func(string) error, serverFactory ServerFactory) *authorizationCodeFlow {
	if port == 0 {
		port = defaultAuthorizationCodePort
	}
	if stateGenerator == nil {
		stateGenerator = uuid.NewString
	}
	if urlOpener == nil {
		urlOpener = browser.OpenURL
	}
	if serverFactory == nil {
		serverFactory = defaultServerFactory
	}

	return &authorizationCodeFlow{
		port:           port,
		stateGenerator: stateGenerator,
		urlOpener:      urlOpener,
		serverFactory:  serverFactory,
	}
}

// acquire runs one interactive authorization code exchange against the source,
// serialized so only one exchange runs at a time.
func (f *authorizationCodeFlow) acquire(ctx context.Context, source authorizationCodeSource) (token *AccessToken, err error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	state := f.stateGenerator()

	server, err := f.serverFactory(state, f.port)
	if err != nil {
		return nil, err
	}

	// Create a context with timeout for the entire code acquisition process.
	authCtx, cancelAuth := context.WithTimeout(ctx, authorizationCodeTimeout)
	defer cancelAuth()

	shutdownCtx, cancelShutdown := context.WithTimeout(authCtx, authorizationCodeShutdownTimeout)
	defer cancelShutdown()

	defer func() {
		if shutdownErr := server.Shutdown(shutdownCtx); shutdownErr != nil && err == nil {
			err = shutdownErr
		}
	}()

	authURL, err := source.getAuthorizationURL(server.GetAddr(), state, nil)
	if err != nil {
		return nil, err
	}

	if err := f.urlOpener(authURL); err != nil {
		return nil, err
	}

	code, _, err := server.Result(shutdownCtx)
	if err != nil {
		return nil, err
	}

	token, err = source.acquireTokenByCode(authCtx, code, server.GetAddr(), state)
	if err != nil {
		return nil, err
	}
	return token, nil
}
