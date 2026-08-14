package appserviceapi

import (
	"testing"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNewCmdbGroupBasedPopulationMethod(t *testing.T) {
	model := NewCmdbGroupBasedPopulationMethod()
	require.NotNil(t, model)

	var _ CmdbGroupBasedPopulationMethod = model
	var _ PopulationMethod = model
}

func TestCmdbGroupBasedPopulationMethodModel_GettersSetters(t *testing.T) {
	model := NewCmdbGroupBasedPopulationMethod()

	err := model.SetGroupID(internal.ToPointer("group-id"))
	require.NoError(t, err)

	got, err := model.GetGroupID()
	require.NoError(t, err)
	assert.Equal(t, internal.ToPointer("group-id"), got)
}

// TestCmdbGroupBasedPopulationMethodModel_Serialize guards against Serialize
// recursing into itself instead of delegating to the embedded
// PopulationMethod - a prior bug here caused a stack overflow.
func TestCmdbGroupBasedPopulationMethodModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *CmdbGroupBasedPopulationMethodModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewCmdbGroupBasedPopulationMethod(),
		},
		{
			name: "happy path - writes group id",
			model: func() *CmdbGroupBasedPopulationMethodModel {
				m := NewCmdbGroupBasedPopulationMethod()
				_ = m.SetGroupID(internal.ToPointer("group-id"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", groupIDKey, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *CmdbGroupBasedPopulationMethodModel {
				m := NewCmdbGroupBasedPopulationMethod()
				_ = m.SetGroupID(internal.ToPointer("group-id"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", groupIDKey, mock.Anything).Return(errWrite)
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

			done := make(chan struct{})
			var err error
			go func() {
				defer close(done)
				err = tt.model.Serialize(writer)
			}()

			select {
			case <-done:
			case <-time.After(5 * time.Second):
				t.Fatal("Serialize did not return - likely infinite recursion")
			}

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestCmdbGroupBasedPopulationMethodModel_GetFieldDeserializers(t *testing.T) {
	model := NewCmdbGroupBasedPopulationMethod()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{typeKey, groupIDKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 2)
}

func TestCreateCmdbGroupBasedPopulationMethodFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateCmdbGroupBasedPopulationMethodFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}
