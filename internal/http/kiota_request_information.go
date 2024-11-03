package http

import (
	"strings"

	"github.com/google/go-querystring/query"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

type KiotaRequestInformation struct {
	abstractions.RequestInformation
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
