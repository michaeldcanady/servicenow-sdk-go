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

func TestServiceModel_GettersSetters(t *testing.T) {
	model := NewService()
	relationship := NewServiceRelationship()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"SysID", func(v any) error { return model.SetSysID(v.(*string)) }, func() (any, error) { return model.GetSysID() }, internal.ToPointer("sys-id")},
		{"Name", func(v any) error { return model.SetName(v.(*string)) }, func() (any, error) { return model.GetName() }, internal.ToPointer("Email_East")},
		{"Number", func(v any) error { return model.SetNumber(v.(*string)) }, func() (any, error) { return model.GetNumber() }, internal.ToPointer("SNSVC0001018")},
		{"Environment", func(v any) error { return model.SetEnvironment(v.(*string)) }, func() (any, error) { return model.GetEnvironment() }, internal.ToPointer("Production")},
		{"Version", func(v any) error { return model.SetVersion(v.(*string)) }, func() (any, error) { return model.GetVersion() }, internal.ToPointer("1.0")},
		{"Relationships", func(v any) error { return model.SetRelationships(v.([]*ServiceRelationship)) }, func() (any, error) { return model.GetRelationships() }, []*ServiceRelationship{relationship}},
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

func TestServiceModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *Service
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewService(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *Service {
				m := NewService()
				_ = m.SetSysID(internal.ToPointer("sys-id"))
				_ = m.SetName(internal.ToPointer("Email_East"))
				_ = m.SetNumber(internal.ToPointer("SNSVC0001018"))
				_ = m.SetEnvironment(internal.ToPointer("Production"))
				_ = m.SetVersion(internal.ToPointer("1.0"))
				_ = m.SetRelationships([]*ServiceRelationship{NewServiceRelationship()})
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteCollectionOfObjectValues", relationshipsKey, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *Service {
				m := NewService()
				_ = m.SetSysID(internal.ToPointer("sys-id"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", sysIDKey, mock.Anything).Return(errWrite)
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

func TestServiceModel_GetFieldDeserializers(t *testing.T) {
	model := NewService()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{sysIDKey, nameKey, numberKey, relationshipsKey, environmentKey, versionKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 6)
}

func TestServiceModel_GetFieldDeserializers_MalformedInput(t *testing.T) {
	model := NewService()
	deserializers := model.GetFieldDeserializers()

	parseNode := &mocking.MockParseNode{}
	parseNode.On("GetCollectionOfObjectValues", mock.Anything).Return([]serialization.Parsable(nil), errWrite)

	err := deserializers[relationshipsKey](parseNode)
	require.ErrorIs(t, err, errWrite)
}

func TestCreateServiceFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateServiceFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}
