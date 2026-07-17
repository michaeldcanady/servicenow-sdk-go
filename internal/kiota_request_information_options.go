package internal

import (
	"errors"
	"strings"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// KiotaRequestInformationOption
type KiotaRequestInformationOption = Option[*KiotaRequestInformation]

// WithMethod
func WithMethod(method abstractions.HttpMethod) KiotaRequestInformationOption {
	return func(requestInformation *KiotaRequestInformation) error {
		if conversion.IsNil(requestInformation) {
			return snerrors.ErrNilConfig
		}
		requestInformation.Method = method
		return nil
	}
}

// WithURLTemplate
func WithURLTemplate(template string) KiotaRequestInformationOption {
	return func(requestInformation *KiotaRequestInformation) error {
		if conversion.IsNil(requestInformation) {
			return snerrors.ErrNilConfig
		}

		template = strings.TrimSpace(template)
		if template == "" {
			return errors.New("template is empty")
		}
		requestInformation.UrlTemplate = template
		return nil
	}
}

// WithPathParameters
func WithPathParameters(pathParameters map[string]string) KiotaRequestInformationOption {
	return func(requestInformation *KiotaRequestInformation) error {
		if conversion.IsNil(requestInformation) {
			return snerrors.ErrNilConfig
		}

		if pathParameters == nil {
			return snerrors.ErrNilPathParameters
		}
		if len(pathParameters) == 0 {
			return snerrors.ErrEmptyPathParameters
		}
		requestInformation.PathParameters = pathParameters
		return nil
	}
}
