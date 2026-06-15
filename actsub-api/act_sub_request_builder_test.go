package actsubapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestActSubRequestBuilder_Hierarchy(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewActSubRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	assert.NotNil(t, builder.Activities())
	assert.NotNil(t, builder.Contexts())
	assert.NotNil(t, builder.Facets())
	assert.NotNil(t, builder.Followings())
	assert.NotNil(t, builder.Preferences())
	assert.NotNil(t, builder.SubObjects())
	assert.NotNil(t, builder.Subscribers())
	assert.NotNil(t, builder.Subscriptions())
	assert.NotNil(t, builder.UserStream())
}

func TestFacetsRequestBuilder_Hierarchy(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewActSubRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	facets := builder.Facets().ByContext("ctx1").ByInstance("inst1")
	assert.Equal(t, "ctx1", facets.GetPathParameters()["activity_context"])
	assert.Equal(t, "inst1", facets.GetPathParameters()["context_instance"])
}

func TestFollowingsRequestBuilder_Hierarchy(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewActSubRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	following := builder.Followings().ByFollower("user1")
	assert.Equal(t, "user1", following.GetPathParameters()["follower"])
}

func TestSubscriptionsRequestBuilder_Hierarchy(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewActSubRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	sub := builder.Subscriptions().BySubscriberId("sub1")
	assert.Equal(t, "sub1", sub.GetPathParameters()["subscriber_id"])

	obj := builder.Subscriptions().ByObjectId("obj1")
	assert.Equal(t, "obj1", obj.GetPathParameters()["sub_obj_id"])

	assert.NotNil(t, obj.IsSubscribed())
	assert.NotNil(t, obj.Subscribe())
	assert.NotNil(t, obj.Unsubscribe())
}

func TestUserStreamRequestBuilder_Hierarchy(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewActSubRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	stream := builder.UserStream().ByProfileId("prof1")
	assert.Equal(t, "prof1", stream.GetPathParameters()["profileId"])
}
