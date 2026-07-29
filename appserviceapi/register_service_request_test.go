package appserviceapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRegisterServiceRequestModel_GettersSetters(t *testing.T) {
	model := NewRegisterServiceRequest()
	details := NewBasicDetails()

	err := model.setBasicDetails(details)
	require.NoError(t, err)

	got, err := model.GetBasicDetails()
	require.NoError(t, err)
	assert.Equal(t, details, got)
}

func TestRegisterServiceRequestModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *RegisterServiceRequest
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewRegisterServiceRequest(),
		},
		{
			name: "happy path - writes nested basic details",
			model: func() *RegisterServiceRequest {
				m := NewRegisterServiceRequest()
				_ = m.setBasicDetails(NewBasicDetails())
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteObjectValue", basicDetailsKey, mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "nested object write error propagates",
			model: func() *RegisterServiceRequest {
				m := NewRegisterServiceRequest()
				_ = m.setBasicDetails(NewBasicDetails())
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteObjectValue", basicDetailsKey, mock.Anything, mock.Anything).Return(errWrite)
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

func TestRegisterServiceRequestModel_GetFieldDeserializers(t *testing.T) {
	model := NewRegisterServiceRequest()
	deserializers := model.GetFieldDeserializers()
	assert.NotNil(t, deserializers[basicDetailsKey])
	assert.Len(t, deserializers, 1)
}

func TestRegisterServiceRequestModel_GetFieldDeserializers_MalformedInput(t *testing.T) {
	model := NewRegisterServiceRequest()
	deserializers := model.GetFieldDeserializers()

	parseNode := &mocking.MockParseNode{}
	parseNode.On("GetObjectValue", mock.Anything).Return(nil, errWrite)

	err := deserializers[basicDetailsKey](parseNode)
	require.ErrorIs(t, err, errWrite)
}

func TestCreateRegisterServiceRequestFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateRegisterServiceRequestFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}
