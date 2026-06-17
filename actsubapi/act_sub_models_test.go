package actsubapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestNewActivitySubscriptionModel(t *testing.T) {
	t.Run("SuccessfulCreation", func(t *testing.T) {
		instance := NewActivitySubscriptionModel()
		assert.NotNil(t, instance)
	})
}

func TestCreateActivitySubscriptionModelFromDiscriminatorValue(t *testing.T) {
	t.Run("SuccessfulCreation", func(t *testing.T) {
		instance, err := CreateActivitySubscriptionModelFromDiscriminatorValue(nil)
		assert.NoError(t, err)
		assert.NotNil(t, instance)
	})
}

func TestActivitySubscriptionModel_Serialize(t *testing.T) {
	t.Run("SuccessfulSerialization", func(t *testing.T) {
		instance := NewActivitySubscriptionModel()
		writer := &mocking.MockSerializationWriter{}

		err := instance.Serialize(writer)
		assert.NoError(t, err)
	})
}

func TestActivitySubscriptionModel_GetFieldDeserializers(t *testing.T) {
	t.Run("SuccessfulGet", func(t *testing.T) {
		instance := NewActivitySubscriptionModel()
		deserializers := instance.GetFieldDeserializers()
		assert.NotNil(t, deserializers)
	})
}

func TestActivitySubscriptionModel_GettersSetters(t *testing.T) {
	instance := NewActivitySubscriptionModel()

	activities := []*Activity{NewActivity()}

	tests := []struct {
		name   string
		setter func() error
		getter func() (interface{}, error)
		want   interface{}
	}{
		{
			name: "Message",
			setter: func() error {
				val := "test message"
				return instance.SetMessage(&val)
			},
			getter: func() (interface{}, error) {
				return instance.GetMessage()
			},
			want: "test message",
		},
		{
			name: "Stream",
			setter: func() error {
				val := "test stream"
				return instance.SetStream(&val)
			},
			getter: func() (interface{}, error) {
				return instance.GetStream()
			},
			want: "test stream",
		},
		{
			name: "User",
			setter: func() error {
				val := "test user"
				return instance.SetUser(&val)
			},
			getter: func() (interface{}, error) {
				return instance.GetUser()
			},
			want: "test user",
		},
		{
			name: "Status",
			setter: func() error {
				val := int64(200)
				return instance.SetStatus(&val)
			},
			getter: func() (interface{}, error) {
				return instance.GetStatus()
			},
			want: int64(200),
		},
		{
			name: "Activities",
			setter: func() error {
				return instance.SetActivities(activities)
			},
			getter: func() (interface{}, error) {
				return instance.GetActivities()
			},
			want: activities,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter()
			assert.NoError(t, err)
			got, err := tt.getter()
			assert.NoError(t, err)

			if tt.name == "Activities" {
				assert.Equal(t, tt.want, got)
			} else {
				// Dereference pointer for comparison
				switch v := got.(type) {
				case *string:
					assert.Equal(t, tt.want, *v)
				case *int64:
					assert.Equal(t, tt.want, *v)
				default:
					assert.Equal(t, tt.want, got)
				}
			}
		})
	}
}
