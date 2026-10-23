// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credential

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"hash"
	"net/url"
	"strings"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

// HMACConfig configures HMAC request signing.
func (c HMACConfig) Validate() error {
	if err := c.ProviderConfig.Validate(); err != nil {
		return err
	}
	if strings.TrimSpace(c.Secret) == "" {
		return snerrors.ErrEmptySecret
	}
	if c.Algorithm < HMACAlgorithmSHA256 || c.Algorithm > HMACAlgorithmSHA512 {
		return snerrors.ErrInvalidHMACAlgorithm
	}
	if c.Location != AuthLocationHeader && c.Location != AuthLocationQuery {
		return snerrors.ErrInvalidAuthLocation
	}
	return nil
}

// HMACProvider signs requests with an HMAC signature.
type HMACProvider struct {
	keyID         string
	secret        string
	algorithm     HMACAlgorithm
	location      AuthLocation
	parameterName string
	signer        HMACSigner
	validator     *authentication.AllowedHostsValidator
}

// NewHMACProvider creates an HMAC authentication provider.
func NewHMACProvider(config HMACConfig) (*HMACProvider, error) {
	if config.Algorithm == HMACAlgorithmUnknown {
		return nil, snerrors.ErrInvalidHMACAlgorithm
	}
	if config.Algorithm < HMACAlgorithmSHA256 || config.Algorithm > HMACAlgorithmSHA512 {
		config.Algorithm = HMACAlgorithmSHA256
	}
	if config.Location == AuthLocationUnknown {
		config.Location = AuthLocationHeader
	}
	if strings.TrimSpace(config.ParameterName) == "" {
		config.ParameterName = config.Algorithm.signatureHeader()
	}
	if config.Signer == nil {
		config.Signer = defaultHMACSigner
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	validator, err := newAllowedHostsValidator(config.ProviderConfig)
	if err != nil {
		return nil, err
	}

	return &HMACProvider{
		keyID:         config.KeyID,
		secret:        config.Secret,
		algorithm:     config.Algorithm,
		location:      config.Location,
		parameterName: config.ParameterName,
		signer:        config.Signer,
		validator:     validator,
	}, nil
}

// Initialize is a no-op for HMAC authentication.
func (p *HMACProvider) Initialize(baseURL string) error {
	if p == nil {
		return snerrors.ErrNilProvider
	}
	return nil
}

// AuthenticateRequest signs the request.
func (p *HMACProvider) AuthenticateRequest(ctx context.Context, request *abstractions.RequestInformation, _ map[string]interface{}) error {
	if p == nil {
		return snerrors.ErrNilProvider
	}
	if request == nil {
		return snerrors.ErrNilRequestInfo
	}
	if ctx == nil {
		return snerrors.ErrNilContext
	}

	uri, allowed, err := requestURI(ctx, request, p.validator)
	if err != nil || !allowed {
		return err
	}

	value, err := p.signer(ctx, request, HMACConfig{
		KeyID:     p.keyID,
		Secret:    p.secret,
		Algorithm: p.algorithm,
		Location:  p.location,
		Signer:    p.signer,
	})
	if err != nil {
		return err
	}

	switch p.location {
	case AuthLocationHeader:
		setHeaderIfAbsent(request, p.parameterName, value)
	case AuthLocationQuery:
		query := uri.Query()
		query.Set(p.parameterName, value)
		uri.RawQuery = query.Encode()
		request.SetUri(*uri)
	default:
		return snerrors.ErrInvalidAuthLocation
	}
	return nil
}

func defaultHMACSigner(ctx context.Context, request *abstractions.RequestInformation, config HMACConfig) (string, error) {
	if request == nil {
		return "", snerrors.ErrNilRequestInfo
	}
	if ctx == nil {
		return "", snerrors.ErrNilContext
	}

	uri, err := request.GetUri()
	if err != nil {
		return "", err
	}

	signedURI := cloneURL(uri)
	if config.Location == AuthLocationQuery && config.ParameterName != "" {
		query := signedURI.Query()
		query.Del(config.ParameterName)
		signedURI.RawQuery = query.Encode()
	}

	var body []byte
	if request.Content != nil {
		body = make([]byte, len(request.Content))
		copy(body, request.Content)
	}

	contentType := ""
	if request.Headers != nil {
		if values := request.Headers.Get("Content-Type"); len(values) > 0 {
			contentType = values[0]
		}
	}

	method := ""
	switch request.Method {
	case abstractions.GET:
		method = "GET"
	case abstractions.POST:
		method = "POST"
	case abstractions.PUT:
		method = "PUT"
	case abstractions.PATCH:
		method = "PATCH"
	case abstractions.DELETE:
		method = "DELETE"
	case abstractions.HEAD:
		method = "HEAD"
	case abstractions.OPTIONS:
		method = "OPTIONS"
	default:
		method = "GET"
	}
	canonical := buildCanonical(method, signedURI, body, contentType)
	signature := hmacSum(config.Algorithm, config.Secret, canonical)
	return formatHMACValue(config.KeyID, signature), nil
}

func buildCanonical(method string, uri *url.URL, body []byte, contentType string) string {
	var builder strings.Builder
	builder.WriteString(method)
	builder.WriteString("\n")
	builder.WriteString(uri.Path)
	if uri.RawQuery != "" {
		builder.WriteString("?")
		builder.WriteString(uri.RawQuery)
	}
	builder.WriteString("\n")
	if contentType != "" {
		builder.WriteString(contentType)
		builder.WriteString("\n")
	}
	builder.Write(body)
	return builder.String()
}

func hmacSum(algorithm HMACAlgorithm, secret, message string) string {
	var h hash.Hash
	switch algorithm {
	case HMACAlgorithmSHA384:
		h = hmac.New(sha512.New384, []byte(secret))
	case HMACAlgorithmSHA512:
		h = hmac.New(sha512.New, []byte(secret))
	default:
		h = hmac.New(sha256.New, []byte(secret))
	}
	_, _ = h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func formatHMACValue(keyID, signature string) string {
	if keyID == "" {
		return signature
	}
	return "keyId=" + keyID + ",signature=" + signature
}

func cloneURL(uri *url.URL) *url.URL {
	if uri == nil {
		return &url.URL{}
	}
	copyURI := *uri
	return &copyURI
}
