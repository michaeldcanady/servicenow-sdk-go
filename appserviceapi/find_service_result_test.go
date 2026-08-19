package appserviceapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestFindServiceResultModel_GettersSetters(t *testing.T) {
	model := NewFindServiceResult()
	service := NewService()

	err := model.SetServices([]*Service{service})
	require.NoError(t, err)

	got, err := model.GetServices()
	require.NoError(t, err)
	assert.Equal(t, []*Service{service}, got)
}

func TestFindServiceResultModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *FindServiceResult
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewFindServiceResult(),
		},
		{
			name: "happy path - writes services collection",
			model: func() *FindServiceResult {
				m := NewFindServiceResult()
				_ = m.SetServices([]*Service{NewService()})
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteCollectionOfObjectValues", servicesKey, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *FindServiceResult {
				m := NewFindServiceResult()
				_ = m.SetServices([]*Service{NewService()})
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteCollectionOfObjectValues", servicesKey, mock.Anything).Return(errWrite)
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

func TestFindServiceResultModel_GetFieldDeserializers(t *testing.T) {
	model := NewFindServiceResult()
	deserializers := model.GetFieldDeserializers()
	assert.NotNil(t, deserializers[servicesKey])
	assert.Len(t, deserializers, 1)
}

func TestFindServiceResultModel_GetFieldDeserializers_MalformedInput(t *testing.T) {
	model := NewFindServiceResult()
	deserializers := model.GetFieldDeserializers()

	parseNode := &mocking.MockParseNode{}
	parseNode.On("GetCollectionOfObjectValues", mock.Anything).Return([]serialization.Parsable(nil), errWrite)

	err := deserializers[servicesKey](parseNode)
	require.ErrorIs(t, err, errWrite)
}

func TestCreateFindServiceResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateFindServiceResultFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}
