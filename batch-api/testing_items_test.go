package batchapi

import (
	"net/http"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
)

type MockClient struct{}

func (c *MockClient) Send(requestInfo internal.RequestInformation, errorMapping internal.ErrorMapping) (*http.Response, error) {
	req, err := requestInfo.ToRequest()
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
