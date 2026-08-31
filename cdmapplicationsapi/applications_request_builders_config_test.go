// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package cdmapplicationsapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const configTestBaseURL = "https://example.service-now.com"

func configTestPathParameters() map[string]string {
	return map[string]string{"baseurl": configTestBaseURL}
}

// configTestHeaders returns a populated header set, to drive the config.Headers branch.
func configTestHeaders() *abstractions.RequestHeaders {
	headers := abstractions.NewRequestHeaders()
	headers.Add("X-Test", "value")

	return headers
}

// configTestOptions returns a request option ready to be stored on a RequestInformation, to
// drive the config.Options branch.
func configTestOptions() []abstractions.RequestOption {
	option := mocking.NewMockRequestOption()
	option.On("GetKey").Return(abstractions.RequestOptionKey{Key: "test-option"})

	return []abstractions.RequestOption{option}
}

// newConfigTestAdapter returns an adapter that serializes bodies and answers both Send and
// SendNoContent successfully, so every verb reaches the end of its happy path.
func newConfigTestAdapter(response any) *mocking.MockRequestAdapter {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(response, nil)
	adapter.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	adapter.On("SendPrimitive", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]byte("content"), nil)

	return adapter
}

func uploadStatusResponse() any {
	return core.NewBaseServiceNowItemResponse[*UploadStatusResultModel](CreateUploadStatusResultFromDiscriminatorValue)
}

func exportsResponse() any {
	return core.NewBaseServiceNowCollectionResponse[*ExportResult](CreateExportResultFromDiscriminatorValue)
}

func exportStatusResponse() any {
	return core.NewBaseServiceNowItemResponse[*ExportStatusResult](CreateExportStatusResultFromDiscriminatorValue)
}

func sharedLibraryApplicationsResponse() any {
	return core.NewBaseServiceNowCollectionResponse[*SharedLibraryComponentApplication](CreateSharedLibraryComponentApplicationFromDiscriminatorValue)
}

// TestRequestBuilders_ConfigurationIsApplied covers the inline request-configuration handling in
// every verb of this package. These builders apply Headers, Options and QueryParameters by hand
// rather than through abstractions.ConfigureRequestInformation, so each branch is only reached
// when a config carrying that field is supplied — the existing happy-path tests all pass nil.
func TestRequestBuilders_ConfigurationIsApplied(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name string
		// response is what the adapter's Send returns; the verbs assert it to their own
		// response interface, so a collection endpoint needs a collection response.
		response func() any
		call     func(adapter *mocking.MockRequestAdapter) error
	}{
		{
			name: "Deployables Delete",
			call: func(adapter *mocking.MockRequestAdapter) error {
				return NewDeployablesRequestBuilderInternal(configTestPathParameters(), adapter).
					Delete(ctx, &DeployablesRequestBuilderDeleteRequestConfiguration{
						Headers:         configTestHeaders(),
						Options:         configTestOptions(),
						QueryParameters: &DeployablesRequestBuilderDeleteQueryParameters{},
					})
			},
		},
		{
			name: "Deployables Put",
			call: func(adapter *mocking.MockRequestAdapter) error {
				_, err := NewDeployablesRequestBuilderInternal(configTestPathParameters(), adapter).
					Put(ctx, &DeployablesRequestBuilderPutRequestConfiguration{
						Headers: configTestHeaders(),
						Options: configTestOptions(),
					})

				return err
			},
		},
		{
			name: "SharedComponents Delete",
			call: func(adapter *mocking.MockRequestAdapter) error {
				return NewSharedComponentsRequestBuilderInternal(configTestPathParameters(), adapter).
					Delete(ctx, &SharedComponentsRequestBuilderDeleteRequestConfiguration{
						Headers:         configTestHeaders(),
						Options:         configTestOptions(),
						QueryParameters: &SharedComponentsRequestBuilderDeleteQueryParameters{},
					})
			},
		},
		{
			name: "SharedComponents Put",
			call: func(adapter *mocking.MockRequestAdapter) error {
				_, err := NewSharedComponentsRequestBuilderInternal(configTestPathParameters(), adapter).
					Put(ctx, &SharedComponentsRequestBuilderPutRequestConfiguration{
						Headers: configTestHeaders(),
						Options: configTestOptions(),
					})

				return err
			},
		},
		{
			name: "UploadStatusItem Get",
			call: func(adapter *mocking.MockRequestAdapter) error {
				_, err := NewUploadStatusItemRequestBuilderInternal(configTestPathParameters(), adapter).
					Get(ctx, &UploadStatusItemRequestBuilderGetRequestConfiguration{
						Headers: configTestHeaders(),
						Options: configTestOptions(),
					})

				return err
			},
		},
		{
			name:     "Exports Get",
			response: exportsResponse,
			call: func(adapter *mocking.MockRequestAdapter) error {
				_, err := NewExportsRequestBuilderInternal(configTestPathParameters(), adapter).
					Get(ctx, &ExportsRequestBuilderGetRequestConfiguration{
						Headers:         configTestHeaders(),
						Options:         configTestOptions(),
						QueryParameters: &ExportsRequestBuilderGetQueryParameters{},
					})

				return err
			},
		},
		{
			name:     "ExportItemStatus Get",
			response: exportStatusResponse,
			call: func(adapter *mocking.MockRequestAdapter) error {
				_, err := NewExportItemStatusRequestBuilderInternal(configTestPathParameters(), adapter).
					Get(ctx, &ExportItemStatusRequestBuilderGetRequestConfiguration{
						Headers: configTestHeaders(),
						Options: configTestOptions(),
					})

				return err
			},
		},
		{
			name: "ExportItemContent Get",
			call: func(adapter *mocking.MockRequestAdapter) error {
				_, err := NewExportItemContentRequestBuilderInternal(configTestPathParameters(), adapter).
					Get(ctx, &ExportItemContentRequestBuilderGetRequestConfiguration{
						Headers: configTestHeaders(),
						Options: configTestOptions(),
					})

				return err
			},
		},
		{
			name:     "SharedLibrariesComponentsApplications Get",
			response: sharedLibraryApplicationsResponse,
			call: func(adapter *mocking.MockRequestAdapter) error {
				_, err := NewSharedLibrariesComponentsApplicationsRequestBuilderInternal(configTestPathParameters(), adapter).
					Get(ctx, &SharedLibrariesComponentsApplicationsRequestBuilderGetRequestConfiguration{
						Headers:         configTestHeaders(),
						Options:         configTestOptions(),
						QueryParameters: &SharedLibrariesComponentsApplicationsRequestBuilderGetQueryParameters{},
					})

				return err
			},
		},
		{
			name: "UploadsComponents Post",
			call: func(adapter *mocking.MockRequestAdapter) error {
				_, err := NewUploadsComponentsRequestBuilderInternal(configTestPathParameters(), adapter).
					Post(ctx, NewComponentUploadRequest(), &UploadsComponentsRequestBuilderPostRequestConfiguration{
						Headers: configTestHeaders(),
						Options: configTestOptions(),
					})

				return err
			},
		},
		{
			name: "UploadsComponentsVars Post",
			call: func(adapter *mocking.MockRequestAdapter) error {
				_, err := NewUploadsComponentsVarsRequestBuilderInternal(configTestPathParameters(), adapter).
					Post(ctx, NewComponentVarsUploadRequest(), &UploadsComponentsVarsRequestBuilderPostRequestConfiguration{
						Headers: configTestHeaders(),
						Options: configTestOptions(),
					})

				return err
			},
		},
		{
			name: "UploadsCollections Post",
			call: func(adapter *mocking.MockRequestAdapter) error {
				_, err := NewUploadsCollectionsRequestBuilderInternal(configTestPathParameters(), adapter).
					Post(ctx, NewCollectionUploadRequest(), &UploadsCollectionsRequestBuilderPostRequestConfiguration{
						Headers: configTestHeaders(),
						Options: configTestOptions(),
					})

				return err
			},
		},
		{
			name: "UploadsCollectionsFile Post",
			call: func(adapter *mocking.MockRequestAdapter) error {
				_, err := NewUploadsCollectionsFileRequestBuilderInternal(configTestPathParameters(), adapter).
					Post(ctx, NewMedia("application/zip", []byte("zip")), &UploadsCollectionsFileRequestBuilderPostRequestConfiguration{
						Headers: configTestHeaders(),
						Options: configTestOptions(),
						QueryParameters: &UploadsCollectionsFileRequestBuilderPostQueryParameters{
							AppName:        internal.ToPointer("app"),
							CollectionName: internal.ToPointer("collection"),
						},
					})

				return err
			},
		},
		{
			name: "UploadsDeployablesFile Post",
			call: func(adapter *mocking.MockRequestAdapter) error {
				_, err := NewUploadsDeployablesFileRequestBuilderInternal(configTestPathParameters(), adapter).
					Post(ctx, NewMedia("application/zip", []byte("zip")), &UploadsDeployablesFileRequestBuilderPostRequestConfiguration{
						Headers:         configTestHeaders(),
						Options:         configTestOptions(),
						QueryParameters: &UploadsDeployablesFileRequestBuilderPostQueryParameters{},
					})

				return err
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := test.response
			if response == nil {
				response = uploadStatusResponse
			}

			require.NoError(t, test.call(newConfigTestAdapter(response())))
		})
	}
}

// newFailingBodyAdapter returns an adapter whose serialization writer fails on write. A
// body-write failure is the only way the body-carrying verbs can fail between their nil guard
// and the Send call.
func newFailingBodyAdapter() *mocking.MockRequestAdapter {
	writer := mocking.NewMockSerializationWriter()
	writer.On("WriteObjectValue", "", mock.Anything, mock.Anything).Return(errNetwork)
	writer.On("Close").Return(nil)

	factory := mocking.NewMockSerializationWriterFactory()
	factory.On("GetSerializationWriter", "application/json").Return(writer, nil)

	adapter := mocking.NewMockRequestAdapter()
	adapter.On("GetSerializationWriterFactory").Return(factory)

	return adapter
}

// TestRequestBuilders_BodySerializationFailure covers the SetContentFromParsable error branch of
// every verb that carries a parsable request body.
func TestRequestBuilders_BodySerializationFailure(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name string
		call func(adapter *mocking.MockRequestAdapter) (any, error)
	}{
		{
			name: "UploadsComponents Post",
			call: func(adapter *mocking.MockRequestAdapter) (any, error) {
				return NewUploadsComponentsRequestBuilderInternal(configTestPathParameters(), adapter).
					Post(ctx, NewComponentUploadRequest(), nil)
			},
		},
		{
			name: "UploadsComponentsVars Post",
			call: func(adapter *mocking.MockRequestAdapter) (any, error) {
				return NewUploadsComponentsVarsRequestBuilderInternal(configTestPathParameters(), adapter).
					Post(ctx, NewComponentVarsUploadRequest(), nil)
			},
		},
		{
			name: "UploadsCollections Post",
			call: func(adapter *mocking.MockRequestAdapter) (any, error) {
				return NewUploadsCollectionsRequestBuilderInternal(configTestPathParameters(), adapter).
					Post(ctx, NewCollectionUploadRequest(), nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			adapter := newFailingBodyAdapter()

			response, err := test.call(adapter)

			require.ErrorIs(t, err, errNetwork)
			assert.Nil(t, response)
			adapter.AssertNotCalled(t, "Send")
		})
	}
}
