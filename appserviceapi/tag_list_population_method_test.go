package appserviceapi

import (
	"testing"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNewTagListPopulationMethod(t *testing.T) {
	model := NewTagListPopulationMethod()
	require.NotNil(t, model)

	var _ TagListPopulationMethod = model
	var _ PopulationMethod = model
}

func TestTagListPopulationMethodModel_GettersSetters(t *testing.T) {
	model := NewTagListPopulationMethod()
	tag := NewTagListPopulationMethodTag()

	err := model.SetTags([]TagListPopulationMethodTag{tag})
	require.NoError(t, err)

	got, err := model.GetTags()
	require.NoError(t, err)
	assert.Equal(t, []TagListPopulationMethodTag{tag}, got)
}

// TestTagListPopulationMethodModel_Serialize guards against Serialize
// recursing into itself instead of delegating to the embedded
// PopulationMethod - a prior bug here caused a stack overflow.
func TestTagListPopulationMethodModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *TagListPopulationMethodModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewTagListPopulationMethod(),
		},
		{
			name: "happy path - writes tags",
			model: func() *TagListPopulationMethodModel {
				m := NewTagListPopulationMethod()
				_ = m.SetTags([]TagListPopulationMethodTag{NewTagListPopulationMethodTag()})
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteCollectionOfObjectValues", tagsKey, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *TagListPopulationMethodModel {
				m := NewTagListPopulationMethod()
				_ = m.SetTags([]TagListPopulationMethodTag{NewTagListPopulationMethodTag()})
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteCollectionOfObjectValues", tagsKey, mock.Anything).Return(errWrite)
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

func TestTagListPopulationMethodModel_GetFieldDeserializers(t *testing.T) {
	model := NewTagListPopulationMethod()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{typeKey, tagsKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 2)
}

func TestTagListPopulationMethodModel_GetFieldDeserializers_MalformedInput(t *testing.T) {
	model := NewTagListPopulationMethod()
	deserializers := model.GetFieldDeserializers()

	parseNode := &mocking.MockParseNode{}
	parseNode.On("GetCollectionOfObjectValues", mock.Anything).Return([]serialization.Parsable(nil), errWrite)

	err := deserializers[tagsKey](parseNode)
	require.ErrorIs(t, err, errWrite)
}

func TestCreateTagListPopulationMethodFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateTagListPopulationMethodFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}
