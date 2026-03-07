package policyapi

import (
	"testing"

	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewRef(t *testing.T) {
	ref := NewRef()
	assert.NotNil(t, ref)
}

func TestCreateRefFromDiscriminatorValue(t *testing.T) {
	ref, err := CreateRefFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, ref)
}

func TestRef_GettersAndSetters(t *testing.T) {
	link := "http://example.com"
	value := "some-value"

	tests := []struct {
		name     string
		setter   func(*Ref, *string) error
		getter   func(*Ref) (*string, error)
		val      *string
		expected *string
	}{
		{
			name:     "Link",
			setter:   func(r *Ref, v *string) error { return r.SetLink(v) },
			getter:   func(r *Ref) (*string, error) { return r.GetLink() },
			val:      &link,
			expected: &link,
		},
		{
			name:     "Value",
			setter:   func(r *Ref, v *string) error { return r.SetValue(v) },
			getter:   func(r *Ref) (*string, error) { return r.GetValue() },
			val:      &value,
			expected: &value,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRef()
			err := tt.setter(r, tt.val)
			assert.NoError(t, err)

			res, err := tt.getter(r)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestRef_Serialize(t *testing.T) {
	writer := mocking.NewMockSerializationWriter()
	writer.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)

	ref := NewRef()
	_ = ref.SetLink(newInternal.ToPointer("link"))
	_ = ref.SetValue(newInternal.ToPointer("value"))

	err := ref.Serialize(writer)
	assert.NoError(t, err)

	var nilRef *Ref
	err = nilRef.Serialize(writer)
	assert.NoError(t, err)
}

func TestRef_GetFieldDeserializers(t *testing.T) {
	ref := NewRef()
	deser := ref.GetFieldDeserializers()
	assert.NotNil(t, deser[refLinkKey])
	assert.NotNil(t, deser[refValueKey])
}
