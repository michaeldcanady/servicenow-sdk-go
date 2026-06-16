package internal

import (
	"testing"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
)

func TestNewBaseServiceNowCollectionResponse(t *testing.T) {
	res := NewBaseServiceNowCollectionResponse[serialization.Parsable](nil)
	assert.NotNil(t, res)
}

func TestBaseServiceNowCollectionResponse_Serialize(t *testing.T) {
	res := NewBaseServiceNowCollectionResponse[serialization.Parsable](nil)
	err := res.Serialize(nil)
	assert.NoError(t, err)

	var nilR *BaseServiceNowCollectionResponse[serialization.Parsable]
	err = nilR.Serialize(nil)
	assert.NoError(t, err)
}

func TestBaseServiceNowCollectionResponse_GetFieldDeserializers(t *testing.T) {
	res := NewBaseServiceNowCollectionResponse[*MainError](CreateMainErrorFromDiscriminatorValue)
	deser := res.GetFieldDeserializers()
	assert.NotNil(t, deser[resultKey])
	assert.NotNil(t, deser[nextKey])
	assert.NotNil(t, deser[previousKey])
	assert.NotNil(t, deser[firstKey])
	assert.NotNil(t, deser[lastKey])
}

func TestBaseServiceNowCollectionResponse_GettersSetters(t *testing.T) {
	res := NewBaseServiceNowCollectionResponse[*MainError](nil)
	link := "https://example.com"
	
	_ = res.SetNextLink(&link)
	_ = res.SetPreviousLink(&link)
	_ = res.SetFirstLink(&link)
	_ = res.SetLastLink(&link)
	
	l, _ := res.GetNextLink()
	assert.Equal(t, &link, l)
	l, _ = res.GetPreviousLink()
	assert.Equal(t, &link, l)
	l, _ = res.GetFirstLink()
	assert.Equal(t, &link, l)
	l, _ = res.GetLastLink()
	assert.Equal(t, &link, l)
}

func TestBaseServiceNowCollectionResponse_Result(t *testing.T) {
	res := NewBaseServiceNowCollectionResponse[*MainError](nil)
	items := []*MainError{NewMainError()}
	_ = res.setResult(items)
	
	r, _ := res.GetResult()
	assert.Equal(t, items, r)
}

func TestBaseServiceNowCollectionResponse_ErrorBranches(t *testing.T) {
	var nilR *BaseServiceNowCollectionResponse[*MainError]
	
	r, err := nilR.GetResult()
	assert.NoError(t, err)
	assert.Nil(t, r)
	
	l, err := nilR.GetNextLink()
	assert.NoError(t, err)
	assert.Nil(t, l)
	
	err = nilR.SetNextLink(nil)
	assert.NoError(t, err)
}

