package appserviceapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestServiceRelationshipModel_GettersSetters(t *testing.T) {
	model := NewServiceRelationship()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"BusinessApp", func(v any) error { return model.SetBusinessApp(v.([]string)) }, func() (any, error) { return model.GetBusinessApp() }, []string{"app-sys-id"}},
		{"BusinessServiceOffering", func(v any) error { return model.SetBusinessServiceOffering(v.([]string)) }, func() (any, error) { return model.GetBusinessServiceOffering() }, []string{"offering-sys-id"}},
		{"TechnicalServiceOffering", func(v any) error { return model.SetTechnicalServiceOffering(v.([]string)) }, func() (any, error) { return model.GetTechnicalServiceOffering() }, []string{"tech-offering-sys-id"}},
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

func TestServiceRelationshipModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *ServiceRelationship
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewServiceRelationship(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *ServiceRelationship {
				m := NewServiceRelationship()
				_ = m.SetBusinessApp([]string{"app-sys-id"})
				_ = m.SetBusinessServiceOffering([]string{"offering-sys-id"})
				_ = m.SetTechnicalServiceOffering([]string{"tech-offering-sys-id"})
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteCollectionOfStringValues", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *ServiceRelationship {
				m := NewServiceRelationship()
				_ = m.SetBusinessApp([]string{"app-sys-id"})
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteCollectionOfStringValues", businessAppKey, mock.Anything).Return(errWrite)
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

func TestServiceRelationshipModel_GetFieldDeserializers(t *testing.T) {
	model := NewServiceRelationship()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{businessAppKey, businessServiceOfferingKey, technicalServiceOfferingKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 3)
}

func TestServiceRelationshipModel_GetFieldDeserializers_MalformedInput(t *testing.T) {
	model := NewServiceRelationship()
	deserializers := model.GetFieldDeserializers()

	parseNode := &mocking.MockParseNode{}
	parseNode.On("GetCollectionOfPrimitiveValues", mock.Anything).Return([]interface{}(nil), errWrite)

	for _, key := range []string{businessAppKey, businessServiceOfferingKey, technicalServiceOfferingKey} {
		err := deserializers[key](parseNode)
		require.ErrorIs(t, err, errWrite)
	}
}

func TestCreateServiceRelationshipFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateServiceRelationshipFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}
