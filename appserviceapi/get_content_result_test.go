package appserviceapi

import (
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGetContentResult(t *testing.T) {
	model := NewGetContentResult()
	require.NotNil(t, model)
}

func TestCreateGetContentResultFromDiscriminatorValue(t *testing.T) {
	val, err := CreateGetContentResultFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.IsType(t, &GetContentResult{}, val)
}

func TestGetContentResult_Serialize(t *testing.T) {
	var nilModel *GetContentResult
	require.NoError(t, nilModel.Serialize(nil))

	model := NewGetContentResult()
	require.ErrorIs(t, model.Serialize(nil), snerrors.ErrNilWriter)
}

func TestGetContentResult_GetFieldDeserializers(t *testing.T) {
	model := NewGetContentResult()
	deserializers := model.GetFieldDeserializers()
	assert.NotNil(t, deserializers[sysIDKey])
	assert.Len(t, deserializers, 1)
}
