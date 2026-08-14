package appserviceapi

import (
	"testing"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNewDiscoveryPopulationMethod(t *testing.T) {
	model := NewDiscoveryPopulationMethod()
	require.NotNil(t, model)

	var _ DiscoveryPopulationMethod = model
	var _ PopulationMethod = model
}

func TestDiscoveryPopulationMethodModel_GettersSetters(t *testing.T) {
	model := NewDiscoveryPopulationMethod()
	attr := NewDiscoveryPopulationMethodAttributeModel()

	err := model.SetEntryPointID(internal.ToPointer("entry-point-id"))
	require.NoError(t, err)

	gotEntryPointID, err := model.GetEntryPointID()
	require.NoError(t, err)
	assert.Equal(t, internal.ToPointer("entry-point-id"), gotEntryPointID)

	err = model.SetAttributes([]DiscoveryPopulationMethodAttribute{attr})
	require.NoError(t, err)

	gotAttrs, err := model.GetAttributes()
	require.NoError(t, err)
	assert.Equal(t, []DiscoveryPopulationMethodAttribute{attr}, gotAttrs)
}

func TestDiscoveryPopulationMethodModel_AddAttribute(t *testing.T) {
	model := NewDiscoveryPopulationMethod()
	attr := NewDiscoveryPopulationMethodAttributeModel()

	err := model.AddAttribute(attr)
	require.NoError(t, err)

	got, err := model.GetAttributes()
	require.NoError(t, err)
	assert.Equal(t, []DiscoveryPopulationMethodAttribute{attr}, got)
}

// TestDiscoveryPopulationMethodModel_Serialize guards against Serialize
// recursing into itself instead of delegating to the embedded
// PopulationMethod - a prior bug here caused a stack overflow.
func TestDiscoveryPopulationMethodModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *DiscoveryPopulationMethodModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewDiscoveryPopulationMethod(),
		},
		{
			name: "happy path - writes attributes and entry point id",
			model: func() *DiscoveryPopulationMethodModel {
				m := NewDiscoveryPopulationMethod()
				_ = m.SetEntryPointID(internal.ToPointer("entry-point-id"))
				_ = m.SetAttributes([]DiscoveryPopulationMethodAttribute{NewDiscoveryPopulationMethodAttributeModel()})
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteCollectionOfObjectValues", attributesKey, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *DiscoveryPopulationMethodModel {
				m := NewDiscoveryPopulationMethod()
				_ = m.SetEntryPointID(internal.ToPointer("entry-point-id"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", entryPointIDKey, mock.Anything).Return(errWrite)
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

func TestDiscoveryPopulationMethodModel_GetFieldDeserializers(t *testing.T) {
	model := NewDiscoveryPopulationMethod()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{typeKey, attributesKey, entryPointIDKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 3)
}

func TestDiscoveryPopulationMethodModel_GetFieldDeserializers_MalformedInput(t *testing.T) {
	model := NewDiscoveryPopulationMethod()
	deserializers := model.GetFieldDeserializers()

	parseNode := &mocking.MockParseNode{}
	parseNode.On("GetCollectionOfObjectValues", mock.Anything).Return([]serialization.Parsable(nil), errWrite)

	err := deserializers[attributesKey](parseNode)
	require.ErrorIs(t, err, errWrite)
}

func TestCreateDiscoveryPopulationMethodFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateDiscoveryPopulationMethodFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}
