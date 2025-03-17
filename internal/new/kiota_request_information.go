package internal

import (
	"strings"

	"github.com/google/go-querystring/query"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// KiotaRequestInformation
type KiotaRequestInformation struct {
	*abstractions.RequestInformation
}

// NewRequestInformation
func NewRequestInformation(opts ...KiotaRequestInformationOption) (*KiotaRequestInformation, error) {
	defaultReqInfo := &KiotaRequestInformation{abstractions.NewRequestInformation()}
	if err := ApplyOptions(defaultReqInfo, opts...); err != nil {
		return nil, err
	}
	return defaultReqInfo, nil
}

func NewRequestInformationWithMethodAndURLTemplateAndPathParameters(method abstractions.HttpMethod, urlTemplate string, pathParameters map[string]string) *KiotaRequestInformation { // nocov // there's no real functionality to test
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(method, urlTemplate, pathParameters)
	return &KiotaRequestInformation{RequestInformation: requestInfo}
}

// AddQueryParameters adds the query parameters to the request by reading the properties from the provided object.
func (request *KiotaRequestInformation) AddQueryParameters(source any) {
	if source == nil || request == nil {
		return
	}

	values, err := query.Values(source)
	if err != nil {
		// -_-
		panic(err)
	}
	for key, values := range values {
		request.QueryParameters[key] = strings.Join(values, ",")
		tmp := make([]any, len(values))
		for i, v := range values {
			tmp[i] = v
		}
		request.QueryParametersAny[key] = tmp
	}
}
