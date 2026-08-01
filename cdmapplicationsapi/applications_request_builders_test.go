package cdmapplicationsapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// errNetwork is a stand-in for a transport-level error returned by the adapter.
var errNetwork = errors.New("network error")

func TestApplicationsRequestBuilder_Deployables(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	deployablesBuilder := builder.Deployables()
	assert.NotNil(t, deployablesBuilder)
	assert.Equal(t, deployablesURLTemplate, deployablesBuilder.GetURLTemplate())
}

func TestApplicationsRequestBuilder_SharedComponents(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	sharedComponentsBuilder := builder.SharedComponents()
	assert.NotNil(t, sharedComponentsBuilder)
	assert.Equal(t, sharedComponentsURLTemplate, sharedComponentsBuilder.GetURLTemplate())
}

func TestApplicationsRequestBuilder_UploadStatus(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	uploadStatusBuilder := builder.UploadStatus()
	assert.NotNil(t, uploadStatusBuilder)
	assert.Equal(t, uploadStatusURLTemplate, uploadStatusBuilder.GetURLTemplate())
}

func TestDeployablesRequestBuilder_Delete(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewDeployablesRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	appName := "test_app"
	name := "Dep-1"
	config := &DeployablesRequestBuilderDeleteRequestConfiguration{
		QueryParameters: &DeployablesRequestBuilderDeleteQueryParameters{
			AppName: &appName,
			Name:    &name,
		},
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, builder.GetURLTemplate(), builder.GetPathParameters())
	requestInfo.AddQueryParameters(*config.QueryParameters)

	uri, err := requestInfo.GetUri()
	require.NoError(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/deployables?appName=test_app&name=Dep-1", uri.String())
}

func TestSharedComponentsRequestBuilder_Delete(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewSharedComponentsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	appName := "test_app"
	sharedComponentName := "Comp-1"
	config := &SharedComponentsRequestBuilderDeleteRequestConfiguration{
		QueryParameters: &SharedComponentsRequestBuilderDeleteQueryParameters{
			AppName:             &appName,
			SharedComponentName: &sharedComponentName,
		},
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, builder.GetURLTemplate(), builder.GetPathParameters())
	requestInfo.AddQueryParameters(*config.QueryParameters)

	uri, err := requestInfo.GetUri()
	require.NoError(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/shared_components?appName=test_app&sharedComponentName=Comp-1", uri.String())
}

func TestUploadStatusItemRequestBuilder_Get(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewUploadStatusRequestBuilderInternal(map[string]string{
		"baseurl": "https://example.service-now.com",
	}, adapter)

	itemBuilder := builder.ByID("upload123")
	assert.NotNil(t, itemBuilder)

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, itemBuilder.GetURLTemplate(), itemBuilder.GetPathParameters())

	uri, err := requestInfo.GetUri()
	require.NoError(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/upload-status/upload123", uri.String())
}

func TestApplicationsRequestBuilder_SharedLibraries(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	sharedLibrariesBuilder := builder.SharedLibraries()
	assert.NotNil(t, sharedLibrariesBuilder)
	assert.Equal(t, "{+baseurl}/api/sn_cdm/applications/shared_libraries", sharedLibrariesBuilder.GetURLTemplate())
}

func TestApplicationsRequestBuilder_Uploads(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	uploadsBuilder := builder.Uploads()
	assert.NotNil(t, uploadsBuilder)
	assert.Equal(t, "{+baseurl}/api/sn_cdm/applications/uploads", uploadsBuilder.GetURLTemplate())
}

func TestDeployablesRequestBuilder_Put(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewDeployablesRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, builder.GetURLTemplate(), builder.GetPathParameters())
	uri, err := requestInfo.GetUri()
	require.NoError(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/deployables", uri.String())
}

func TestSharedComponentsRequestBuilder_Put(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewSharedComponentsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, builder.GetURLTemplate(), builder.GetPathParameters())
	uri, err := requestInfo.GetUri()
	require.NoError(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/shared_components", uri.String())
}

func TestExportsRequestBuilder_Get(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewDeployablesRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter).Exports()

	appName := "test_app"
	deployableName := "test_dep"
	config := &ExportsRequestBuilderGetRequestConfiguration{
		QueryParameters: &ExportsRequestBuilderGetQueryParameters{
			AppName:        &appName,
			DeployableName: &deployableName,
		},
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, builder.GetURLTemplate(), builder.GetPathParameters())
	requestInfo.AddQueryParameters(*config.QueryParameters)

	uri, err := requestInfo.GetUri()
	require.NoError(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/deployables/exports?appName=test_app&deployableName=test_dep", uri.String())
}

func TestExportItemRequestBuilder(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewDeployablesRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter).Exports().ByID("exp-123")

	statusBuilder := builder.Status()
	assert.NotNil(t, statusBuilder)
	requestInfoStatus := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, statusBuilder.GetURLTemplate(), statusBuilder.GetPathParameters())
	uriStatus, err := requestInfoStatus.GetUri()
	require.NoError(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/deployables/exports/exp-123/status", uriStatus.String())

	contentBuilder := builder.Content()
	assert.NotNil(t, contentBuilder)
	requestInfoContent := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, contentBuilder.GetURLTemplate(), contentBuilder.GetPathParameters())
	uriContent, err := requestInfoContent.GetUri()
	require.NoError(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/deployables/exports/exp-123/content", uriContent.String())
}

func TestSharedLibrariesComponentsApplicationsRequestBuilder_Get(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter).
		SharedLibraries().Components().Applications()

	appName := "test_app"
	sharedCompName := "shared_comp"
	name := "name"
	config := &SharedLibrariesComponentsApplicationsRequestBuilderGetRequestConfiguration{
		QueryParameters: &SharedLibrariesComponentsApplicationsRequestBuilderGetQueryParameters{
			AppName:             &appName,
			SharedComponentName: &sharedCompName,
			Name:                &name,
		},
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, builder.GetURLTemplate(), builder.GetPathParameters())
	requestInfo.AddQueryParameters(*config.QueryParameters)

	uri, err := requestInfo.GetUri()
	require.NoError(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/shared_libraries/components/applications?appName=test_app&sharedComponentName=shared_comp&name=name", uri.String())
}

func TestUploadsComponentsRequestBuilder_Post(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter).
		Uploads().Components()

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, builder.GetURLTemplate(), builder.GetPathParameters())
	uri, err := requestInfo.GetUri()
	require.NoError(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/uploads/components", uri.String())
}

func TestUploadsComponentsVarsRequestBuilder_Post(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter).
		Uploads().Components().Vars()

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, builder.GetURLTemplate(), builder.GetPathParameters())
	uri, err := requestInfo.GetUri()
	require.NoError(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/uploads/components/vars", uri.String())
}

func TestUploadsCollectionsRequestBuilder_Post(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter).
		Uploads().Collections()

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, builder.GetURLTemplate(), builder.GetPathParameters())
	uri, err := requestInfo.GetUri()
	require.NoError(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/uploads/collections", uri.String())
}

func TestUploadsCollectionsFileRequestBuilder_Post(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter).
		Uploads().Collections().File()

	appName := "test_app"
	collectionName := "test_coll"
	config := &UploadsCollectionsFileRequestBuilderPostRequestConfiguration{
		QueryParameters: &UploadsCollectionsFileRequestBuilderPostQueryParameters{
			AppName:        &appName,
			CollectionName: &collectionName,
		},
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, builder.GetURLTemplate(), builder.GetPathParameters())
	requestInfo.AddQueryParameters(*config.QueryParameters)

	uri, err := requestInfo.GetUri()
	require.NoError(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/uploads/collections/file?appName=test_app&collectionName=test_coll", uri.String())
}

func TestUploadsDeployablesFileRequestBuilder_Post(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter).
		Uploads().Deployables().File()

	appName := "test_app"
	deployableName := "test_dep"
	config := &UploadsDeployablesFileRequestBuilderPostRequestConfiguration{
		QueryParameters: &UploadsDeployablesFileRequestBuilderPostQueryParameters{
			AppName:        &appName,
			DeployableName: &deployableName,
		},
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, builder.GetURLTemplate(), builder.GetPathParameters())
	requestInfo.AddQueryParameters(*config.QueryParameters)

	uri, err := requestInfo.GetUri()
	require.NoError(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/uploads/deployables/file?appName=test_app&deployableName=test_dep", uri.String())
}

func TestDeployablesRequestBuilder_NilReceiverGuards(t *testing.T) {
	builders := map[string]*DeployablesRequestBuilder{
		"nil builder":              nil,
		"nil inner RequestBuilder": {},
	}
	for name, builder := range builders {
		t.Run(name, func(t *testing.T) {
			err := builder.Delete(context.Background(), nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)

			resp, err := builder.Put(context.Background(), nil, nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, resp)
		})
	}
}

func TestSharedComponentsRequestBuilder_NilReceiverGuards(t *testing.T) {
	builders := map[string]*SharedComponentsRequestBuilder{
		"nil builder":              nil,
		"nil inner RequestBuilder": {},
	}
	for name, builder := range builders {
		t.Run(name, func(t *testing.T) {
			err := builder.Delete(context.Background(), nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)

			resp, err := builder.Put(context.Background(), nil, nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, resp)
		})
	}
}

func TestUploadStatusItemRequestBuilder_NilReceiverGuards(t *testing.T) {
	builders := map[string]*UploadStatusItemRequestBuilder{
		"nil builder":              nil,
		"nil inner RequestBuilder": {},
	}
	for name, builder := range builders {
		t.Run(name, func(t *testing.T) {
			resp, err := builder.Get(context.Background(), nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, resp)
		})
	}
}

func TestExportsRequestBuilder_NilReceiverGuards(t *testing.T) {
	builders := map[string]*ExportsRequestBuilder{
		"nil builder":              nil,
		"nil inner RequestBuilder": {},
	}
	for name, builder := range builders {
		t.Run(name, func(t *testing.T) {
			resp, err := builder.Get(context.Background(), nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, resp)
		})
	}
}

func TestExportItemStatusRequestBuilder_NilReceiverGuards(t *testing.T) {
	builders := map[string]*ExportItemStatusRequestBuilder{
		"nil builder":              nil,
		"nil inner RequestBuilder": {},
	}
	for name, builder := range builders {
		t.Run(name, func(t *testing.T) {
			resp, err := builder.Get(context.Background(), nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, resp)
		})
	}
}

func TestExportItemContentRequestBuilder_NilReceiverGuards(t *testing.T) {
	builders := map[string]*ExportItemContentRequestBuilder{
		"nil builder":              nil,
		"nil inner RequestBuilder": {},
	}
	for name, builder := range builders {
		t.Run(name, func(t *testing.T) {
			resp, err := builder.Get(context.Background(), nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, resp)
		})
	}
}

func TestSharedLibrariesComponentsApplicationsRequestBuilder_NilReceiverGuards(t *testing.T) {
	builders := map[string]*SharedLibrariesComponentsApplicationsRequestBuilder{
		"nil builder":              nil,
		"nil inner RequestBuilder": {},
	}
	for name, builder := range builders {
		t.Run(name, func(t *testing.T) {
			resp, err := builder.Get(context.Background(), nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, resp)
		})
	}
}

func TestUploadsComponentsRequestBuilder_NilReceiverGuards(t *testing.T) {
	builders := map[string]*UploadsComponentsRequestBuilder{
		"nil builder":              nil,
		"nil inner RequestBuilder": {},
	}
	for name, builder := range builders {
		t.Run(name, func(t *testing.T) {
			resp, err := builder.Post(context.Background(), nil, nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, resp)
		})
	}
}

func TestUploadsComponentsVarsRequestBuilder_NilReceiverGuards(t *testing.T) {
	builders := map[string]*UploadsComponentsVarsRequestBuilder{
		"nil builder":              nil,
		"nil inner RequestBuilder": {},
	}
	for name, builder := range builders {
		t.Run(name, func(t *testing.T) {
			resp, err := builder.Post(context.Background(), nil, nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, resp)
		})
	}
}

func TestUploadsCollectionsRequestBuilder_NilReceiverGuards(t *testing.T) {
	builders := map[string]*UploadsCollectionsRequestBuilder{
		"nil builder":              nil,
		"nil inner RequestBuilder": {},
	}
	for name, builder := range builders {
		t.Run(name, func(t *testing.T) {
			resp, err := builder.Post(context.Background(), nil, nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, resp)
		})
	}
}

func TestUploadsCollectionsFileRequestBuilder_NilReceiverGuards(t *testing.T) {
	builders := map[string]*UploadsCollectionsFileRequestBuilder{
		"nil builder":              nil,
		"nil inner RequestBuilder": {},
	}
	for name, builder := range builders {
		t.Run(name, func(t *testing.T) {
			resp, err := builder.Post(context.Background(), nil, nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, resp)
		})
	}
}

func TestUploadsDeployablesFileRequestBuilder_NilReceiverGuards(t *testing.T) {
	builders := map[string]*UploadsDeployablesFileRequestBuilder{
		"nil builder":              nil,
		"nil inner RequestBuilder": {},
	}
	for name, builder := range builders {
		t.Run(name, func(t *testing.T) {
			resp, err := builder.Post(context.Background(), nil, nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, resp)
		})
	}
}

func TestDeployablesRequestBuilder_NilRequestAdapterGuards(t *testing.T) {
	builder := NewDeployablesRequestBuilderInternal(map[string]string{}, nil)

	err := builder.Delete(context.Background(), nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)

	resp, err := builder.Put(context.Background(), nil, nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestSharedComponentsRequestBuilder_NilRequestAdapterGuards(t *testing.T) {
	builder := NewSharedComponentsRequestBuilderInternal(map[string]string{}, nil)

	err := builder.Delete(context.Background(), nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)

	resp, err := builder.Put(context.Background(), nil, nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestUploadStatusItemRequestBuilder_NilRequestAdapterGuards(t *testing.T) {
	builder := NewUploadStatusItemRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Get(context.Background(), nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestExportsRequestBuilder_NilRequestAdapterGuards(t *testing.T) {
	builder := NewExportsRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Get(context.Background(), nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestExportItemStatusRequestBuilder_NilRequestAdapterGuards(t *testing.T) {
	builder := NewExportItemStatusRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Get(context.Background(), nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestExportItemContentRequestBuilder_NilRequestAdapterGuards(t *testing.T) {
	builder := NewExportItemContentRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Get(context.Background(), nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestSharedLibrariesComponentsApplicationsRequestBuilder_NilRequestAdapterGuards(t *testing.T) {
	builder := NewSharedLibrariesComponentsApplicationsRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Get(context.Background(), nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestUploadsComponentsRequestBuilder_NilRequestAdapterGuards(t *testing.T) {
	builder := NewUploadsComponentsRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Post(context.Background(), nil, nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestUploadsComponentsVarsRequestBuilder_NilRequestAdapterGuards(t *testing.T) {
	builder := NewUploadsComponentsVarsRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Post(context.Background(), nil, nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestUploadsCollectionsRequestBuilder_NilRequestAdapterGuards(t *testing.T) {
	builder := NewUploadsCollectionsRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Post(context.Background(), nil, nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestUploadsCollectionsFileRequestBuilder_NilRequestAdapterGuards(t *testing.T) {
	builder := NewUploadsCollectionsFileRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Post(context.Background(), nil, nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestUploadsDeployablesFileRequestBuilder_NilRequestAdapterGuards(t *testing.T) {
	builder := NewUploadsDeployablesFileRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Post(context.Background(), nil, nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

// ---------------------------------------------------------------------------
// Happy-path / adapter-error coverage for verb methods.
// ---------------------------------------------------------------------------

func TestDeployablesRequestBuilder_Delete_HappyAndError(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
	}{
		{
			name: "happy path",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).Return(errNetwork)
			},
			wantErr: errNetwork,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := mocking.NewMockRequestAdapter()
			tt.setupMock(adapter)
			builder := NewDeployablesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

			err := builder.Delete(context.Background(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
			adapter.AssertExpectations(t)
		})
	}
}

func TestDeployablesRequestBuilder_Put_HappyAndError(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
		wantNil   bool
	}{
		{
			name: "happy path",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowItemResponse[*UploadStatusResult](CreateUploadStatusResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns nil, nil",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := mocking.NewMockRequestAdapter()
			tt.setupMock(adapter)
			builder := NewDeployablesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

			resp, err := builder.Put(context.Background(), NewDeployableUpdateRequest(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			if tt.wantNil {
				assert.Nil(t, resp)
				return
			}
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

// isDefaultErrorMapping reports whether the given argument is a non-nil
// abstractions.ErrorMappings with core.DefaultErrorMapping()'s status-code keys.
func isDefaultErrorMapping(v any) bool {
	mapping, ok := v.(abstractions.ErrorMappings)
	if !ok || mapping == nil {
		return false
	}
	for _, code := range []string{"400", "401", "403", "404", "429", "5XX", "XXX"} {
		if _, ok := mapping[code]; !ok {
			return false
		}
	}
	return len(mapping) == 7
}

// TestDeployablesRequestBuilder_Put_PassesDefaultErrorMapping guards against #565:
// CDM builders previously passed literal nil instead of core.DefaultErrorMapping(),
// so ServiceNow API errors never mapped to a typed core.ServiceNowError.
func TestDeployablesRequestBuilder_Put_PassesDefaultErrorMapping(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.MatchedBy(isDefaultErrorMapping)).
		Return(core.NewBaseServiceNowItemResponse[*UploadStatusResult](CreateUploadStatusResultFromDiscriminatorValue), nil)

	builder := NewDeployablesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	_, err := builder.Put(context.Background(), NewDeployableUpdateRequest(), nil)

	require.NoError(t, err)
	adapter.AssertExpectations(t)
}

func TestSharedComponentsRequestBuilder_Delete_HappyAndError(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
	}{
		{
			name: "happy path",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).Return(errNetwork)
			},
			wantErr: errNetwork,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := mocking.NewMockRequestAdapter()
			tt.setupMock(adapter)
			builder := NewSharedComponentsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

			err := builder.Delete(context.Background(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
			adapter.AssertExpectations(t)
		})
	}
}

func TestSharedComponentsRequestBuilder_Put_HappyAndError(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
		wantNil   bool
	}{
		{
			name: "happy path",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowItemResponse[*UploadStatusResult](CreateUploadStatusResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns nil, nil",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := mocking.NewMockRequestAdapter()
			tt.setupMock(adapter)
			builder := NewSharedComponentsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

			resp, err := builder.Put(context.Background(), NewSharedComponentUpdateRequest(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			if tt.wantNil {
				assert.Nil(t, resp)
				return
			}
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

func TestUploadStatusItemRequestBuilder_Get_HappyAndError(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
		wantNil   bool
	}{
		{
			name: "happy path",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowItemResponse[*UploadStatusResult](CreateUploadStatusResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns nil, nil",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := mocking.NewMockRequestAdapter()
			tt.setupMock(adapter)
			builder := NewUploadStatusRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter).ByID("upload123")

			resp, err := builder.Get(context.Background(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			if tt.wantNil {
				assert.Nil(t, resp)
				return
			}
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

func TestExportsRequestBuilder_Get_HappyAndError(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
		wantNil   bool
	}{
		{
			name: "happy path",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowCollectionResponse[*ExportResult](CreateExportResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns nil, nil",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := mocking.NewMockRequestAdapter()
			tt.setupMock(adapter)
			builder := NewDeployablesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter).Exports()

			resp, err := builder.Get(context.Background(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			if tt.wantNil {
				assert.Nil(t, resp)
				return
			}
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

func TestExportItemStatusRequestBuilder_Get_HappyAndError(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
		wantNil   bool
	}{
		{
			name: "happy path",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowItemResponse[*ExportStatusResult](CreateExportStatusResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns nil, nil",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := mocking.NewMockRequestAdapter()
			tt.setupMock(adapter)
			builder := NewDeployablesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter).Exports().ByID("exp-123").Status()

			resp, err := builder.Get(context.Background(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			if tt.wantNil {
				assert.Nil(t, resp)
				return
			}
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

func TestExportItemContentRequestBuilder_Get_HappyAndError(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
		wantNil   bool
	}{
		{
			name: "happy path",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("SendPrimitive", mock.Anything, mock.Anything, "[]byte", mock.Anything).
					Return([]byte("content"), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("SendPrimitive", mock.Anything, mock.Anything, "[]byte", mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns nil, nil",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("SendPrimitive", mock.Anything, mock.Anything, "[]byte", mock.Anything).
					Return(nil, nil)
			},
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := mocking.NewMockRequestAdapter()
			tt.setupMock(adapter)
			builder := NewDeployablesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter).Exports().ByID("exp-123").Content()

			resp, err := builder.Get(context.Background(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			if tt.wantNil {
				assert.Nil(t, resp)
				return
			}
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

func TestSharedLibrariesComponentsApplicationsRequestBuilder_Get_HappyAndError(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
		wantNil   bool
	}{
		{
			name: "happy path",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowCollectionResponse[*SharedLibraryComponentApplication](CreateSharedLibraryComponentApplicationFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns nil, nil",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := mocking.NewMockRequestAdapter()
			tt.setupMock(adapter)
			builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter).
				SharedLibraries().Components().Applications()

			resp, err := builder.Get(context.Background(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			if tt.wantNil {
				assert.Nil(t, resp)
				return
			}
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

func TestUploadsComponentsRequestBuilder_Post_HappyAndError(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
		wantNil   bool
	}{
		{
			name: "happy path",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowItemResponse[*UploadStatusResult](CreateUploadStatusResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns nil, nil",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := mocking.NewMockRequestAdapter()
			tt.setupMock(adapter)
			builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter).
				Uploads().Components()

			resp, err := builder.Post(context.Background(), NewComponentUploadRequest(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			if tt.wantNil {
				assert.Nil(t, resp)
				return
			}
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

func TestUploadsComponentsVarsRequestBuilder_Post_HappyAndError(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
		wantNil   bool
	}{
		{
			name: "happy path",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowItemResponse[*UploadStatusResult](CreateUploadStatusResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns nil, nil",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := mocking.NewMockRequestAdapter()
			tt.setupMock(adapter)
			builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter).
				Uploads().Components().Vars()

			resp, err := builder.Post(context.Background(), NewComponentVarsUploadRequest(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			if tt.wantNil {
				assert.Nil(t, resp)
				return
			}
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

func TestUploadsCollectionsRequestBuilder_Post_HappyAndError(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
		wantNil   bool
	}{
		{
			name: "happy path",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowItemResponse[*UploadStatusResult](CreateUploadStatusResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns nil, nil",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := mocking.NewMockRequestAdapter()
			tt.setupMock(adapter)
			builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter).
				Uploads().Collections()

			resp, err := builder.Post(context.Background(), NewCollectionUploadRequest(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			if tt.wantNil {
				assert.Nil(t, resp)
				return
			}
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

func TestUploadsCollectionsFileRequestBuilder_Post_HappyAndError(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
		wantNil   bool
	}{
		{
			name: "happy path",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowItemResponse[*UploadStatusResult](CreateUploadStatusResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns nil, nil",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := mocking.NewMockRequestAdapter()
			tt.setupMock(adapter)
			builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter).
				Uploads().Collections().File()

			resp, err := builder.Post(context.Background(), NewMedia("application/json", []byte("data")), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			if tt.wantNil {
				assert.Nil(t, resp)
				return
			}
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

func TestUploadsDeployablesFileRequestBuilder_Post_HappyAndError(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
		wantNil   bool
	}{
		{
			name: "happy path",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowItemResponse[*UploadStatusResult](CreateUploadStatusResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns nil, nil",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := mocking.NewMockRequestAdapter()
			tt.setupMock(adapter)
			builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter).
				Uploads().Deployables().File()

			resp, err := builder.Post(context.Background(), NewMedia("application/json", []byte("data")), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			if tt.wantNil {
				assert.Nil(t, resp)
				return
			}
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}
