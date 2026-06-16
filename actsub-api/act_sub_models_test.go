package actsubapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestNewActivitySubscriptionModel(t *testing.T) {
	instance := NewActivitySubscriptionModel()
	assert.NotNil(t, instance)
}

func TestCreateActivitySubscriptionModelFromDiscriminatorValue(t *testing.T) {
	instance, err := CreateActivitySubscriptionModelFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, instance)
}

func TestActivitySubscriptionModel_Serialize(t *testing.T) {
	instance := NewActivitySubscriptionModel()
	writer := &mocking.MockSerializationWriter{}

	err := instance.Serialize(writer)
	assert.NoError(t, err)
}

func TestActivitySubscriptionModel_GetFieldDeserializers(t *testing.T) {
	instance := NewActivitySubscriptionModel()
	deserializers := instance.GetFieldDeserializers()
	assert.NotNil(t, deserializers)
}

func TestActivitySubscriptionModel_GettersSetters(t *testing.T) {
	instance := NewActivitySubscriptionModel()

	message := "test message"
	err := instance.SetMessage(&message)
	assert.NoError(t, err)
	resMessage, err := instance.GetMessage()
	assert.NoError(t, err)
	assert.Equal(t, &message, resMessage)

	stream := "test stream"
	err = instance.SetStream(&stream)
	assert.NoError(t, err)
	resStream, err := instance.GetStream()
	assert.NoError(t, err)
	assert.Equal(t, &stream, resStream)

	user := "test user"
	err = instance.SetUser(&user)
	assert.NoError(t, err)
	resUser, err := instance.GetUser()
	assert.NoError(t, err)
	assert.Equal(t, &user, resUser)

	status := int64(200)
	err = instance.SetStatus(&status)
	assert.NoError(t, err)
	resStatus, err := instance.GetStatus()
	assert.NoError(t, err)
	assert.Equal(t, &status, resStatus)

	activities := []*Activity{NewActivity()}
	err = instance.SetActivities(activities)
	assert.NoError(t, err)
	resActivities, err := instance.GetActivities()
	assert.NoError(t, err)
	assert.Equal(t, activities, resActivities)
}
