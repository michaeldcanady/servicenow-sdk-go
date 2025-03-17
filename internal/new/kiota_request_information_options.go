package internal

import (
	"errors"
	"strings"

	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// KiotaRequestInformationOption
type KiotaRequestInformationOption = Option[*KiotaRequestInformation]

// WithMethod
func WithMethod(method abstractions.HttpMethod) KiotaRequestInformationOption {
	return func(requestInformation *KiotaRequestInformation) error {
		if IsNil(requestInformation) {
			return errors.New("config is nil")
		}
		requestInformation.Method = method
		return nil
	}
}

// WithURLTemplate
func WithURLTemplate(template string) KiotaRequestInformationOption {
	return func(requestInformation *KiotaRequestInformation) error {
		if IsNil(requestInformation) {
			return errors.New("config is nil")
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
		if IsNil(requestInformation) {
			return errors.New("config is nil")
		}

		if pathParameters == nil {
			return errors.New("pathParameters is nil")
		}
		if len(pathParameters) == 0 {
			return errors.New("pathParameters is empty")
		}
		requestInformation.PathParameters = pathParameters
		return nil
	}
}
