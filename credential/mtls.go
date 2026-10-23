// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credential

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
)

// NewMTLSHTTPClient creates an HTTP client configured for mutual TLS.
func NewMTLSHTTPClient(config MTLSConfig) (*http.Client, error) {
	if config.ClientCertificate == nil {
		return nil, errors.New("client certificate is required")
	}

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{*config.ClientCertificate},
		RootCAs:            config.RootCAs,
		InsecureSkipVerify: config.InsecureSkipVerify, //nolint:gosec // Caller must explicitly opt in.
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
		DialContext: (&net.Dialer{
			Timeout:   30,
			KeepAlive: 30,
		}).DialContext,
	}

	return &http.Client{
		Transport: transport,
	}, nil
}

// NewMTLSHTTPClientWithBase returns a standard MTLS client. Use it with
// WithHTTPClient on the request adapter or the OAuth HTTP client when the
// transport requires client certificates.
func NewMTLSHTTPClientWithBase(config MTLSConfig, base *http.Transport) (*http.Client, error) {
	client, err := NewMTLSHTTPClient(config)
	if err != nil {
		return nil, err
	}
	if base != nil {
		transport := base.Clone()
		transport.TLSClientConfig = client.Transport.(*http.Transport).TLSClientConfig
		client.Transport = transport
	}
	return client, nil
}

// Validate checks the MTLS configuration.
func (c MTLSConfig) Validate() error {
	if c.ClientCertificate == nil {
		return errors.New("client certificate is required")
	}
	if c.InsecureSkipVerify {
		fmt.Println("warning: TLS certificate verification is disabled")
	}
	return nil
}
