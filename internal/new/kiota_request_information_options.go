package internal

import (
	"errors"
	"strings"

	abstractions "github.com/microsoft/kiota-abstractions-go"
)

type KiotaRequestInformationOption = Option[*KiotaRequestInformation]

func WithMethod(method abstractions.HttpMethod) KiotaRequestInformationOption {
	return func(requestInformation *KiotaRequestInformation) error {
		requestInformation.Method = method
		return nil
	}
}

func WithURLTemplate(template string) KiotaRequestInformationOption {
	return func(requestInformation *KiotaRequestInformation) error {
		template = strings.TrimSpace(template)
		if template == "" {
			return errors.New("template is empty")
		}
		requestInformation.UrlTemplate = template
		return nil
	}
}

func WithPathParameters(pathParameters map[string]string) KiotaRequestInformationOption {
	return func(requestInformation *KiotaRequestInformation) error {
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
