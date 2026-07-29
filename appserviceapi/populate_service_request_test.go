package appserviceapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestPopulateServiceRequestModel_GettersSetters(t *testing.T) {
	model := NewPopulateServiceRequest()

	rel := NewServiceRelation()
	_ = rel.setParent(internal.ToPointer("parent123"))

	err := model.setServiceRelations([]*ServiceRelation{rel})
	require.NoError(t, err)

	got, err := model.GetServiceRelations()
	require.NoError(t, err)
	assert.Equal(t, []*ServiceRelation{rel}, got)
}

func TestPopulateServiceRequestModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *PopulateServiceRequest
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewPopulateServiceRequest(),
		},
		{
			name: "happy path - writes service relations collection",
			model: func() *PopulateServiceRequest {
				m := NewPopulateServiceRequest()
				rel := NewServiceRelation()
				_ = rel.setParent(internal.ToPointer("parent123"))
				_ = m.setServiceRelations([]*ServiceRelation{rel})
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteCollectionOfObjectValues", serviceRelationsKey, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *PopulateServiceRequest {
				m := NewPopulateServiceRequest()
				rel := NewServiceRelation()
				_ = m.setServiceRelations([]*ServiceRelation{rel})
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteCollectionOfObjectValues", serviceRelationsKey, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			if tt.setupMock != nil {
				tt.setupMock(writer)
			}

			err := tt.model.Serialize(writer)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestPopulateServiceRequestModel_GetFieldDeserializers(t *testing.T) {
	model := NewPopulateServiceRequest()
	deserializers := model.GetFieldDeserializers()
	assert.NotNil(t, deserializers[serviceRelationsKey])
	assert.Len(t, deserializers, 1)
}

func TestPopulateServiceRequestModel_GetFieldDeserializers_MalformedInput(t *testing.T) {
	model := NewPopulateServiceRequest()
	deserializers := model.GetFieldDeserializers()

	parseNode := &mocking.MockParseNode{}
	parseNode.On("GetCollectionOfObjectValues", mock.Anything).Return([]serialization.Parsable(nil), errWrite)

	err := deserializers[serviceRelationsKey](parseNode)
	require.ErrorIs(t, err, errWrite)
}

func TestCreatePopulateServiceRequestFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreatePopulateServiceRequestFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

var _ serialization.Parsable = (*ServiceRelation)(nil)

func TestPopulateServiceResultModel_GettersSetters(t *testing.T) {
	model := NewPopulateServiceResult()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"Status", func(v any) error { return model.setStatus(v.(*string)) }, func() (any, error) { return model.GetStatus() }, internal.ToPointer("success")},
		{"Message", func(v any) error { return model.setMessage(v.(*string)) }, func() (any, error) { return model.GetMessage() }, internal.ToPointer("populated")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			require.NoError(t, err)
			got, err := tt.getter()
			require.NoError(t, err)
			assert.Equal(t, tt.value, got)
		})
	}
}

func TestPopulateServiceResultModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *PopulateServiceResult
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewPopulateServiceResult(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *PopulateServiceResult {
				m := NewPopulateServiceResult()
				_ = m.setStatus(internal.ToPointer("success"))
				_ = m.setMessage(internal.ToPointer("populated"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *PopulateServiceResult {
				m := NewPopulateServiceResult()
				_ = m.setStatus(internal.ToPointer("success"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", statusKey, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			if tt.setupMock != nil {
				tt.setupMock(writer)
			}

			err := tt.model.Serialize(writer)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestPopulateServiceResultModel_GetFieldDeserializers(t *testing.T) {
	model := NewPopulateServiceResult()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{statusKey, messageKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 2)
}

func TestCreatePopulateServiceResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreatePopulateServiceResultFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}
