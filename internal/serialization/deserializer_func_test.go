package serialization

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeserializeStringFunc(t *testing.T) {
	val := "test"
	tests := []struct {
		name    string
		mock    func(*mocking.MockParseNode)
		wantErr bool
		wantVal *string
	}{
		{
			name: "Success",
			mock: func(m *mocking.MockParseNode) {
				m.On("GetStringValue").Return(&val, nil)
			},
			wantErr: false,
			wantVal: &val,
		},
		{
			name: "Error",
			mock: func(m *mocking.MockParseNode) {
				m.On("GetStringValue").Return((*string)(nil), errors.New("err"))
			},
			wantErr: true,
			wantVal: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := mocking.NewMockParseNode()
			tt.mock(node)
			
			var result *string
			setter := func(v *string) error {
				result = v
				return nil
			}
			
			err := DeserializeStringFunc()(setter)(node)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantVal, result)
			}
			node.AssertExpectations(t)
		})
	}
}

func TestDeserializeInt64Func(t *testing.T) {
	val := int64(123)
	tests := []struct {
		name    string
		mock    func(*mocking.MockParseNode)
		wantErr bool
		wantVal *int64
	}{
		{
			name: "Success",
			mock: func(m *mocking.MockParseNode) {
				m.On("GetInt64Value").Return(&val, nil)
			},
			wantErr: false,
			wantVal: &val,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := mocking.NewMockParseNode()
			tt.mock(node)
			
			var result *int64
			setter := func(v *int64) error {
				result = v
				return nil
			}
			
			err := DeserializeInt64Func()(setter)(node)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantVal, result)
			node.AssertExpectations(t)
		})
	}
}

func TestDeserializeCollectionOfObjectValuesFunc(t *testing.T) {
	parsable := mocking.NewMockParsable()
	val := []serialization.Parsable{parsable}
	
	tests := []struct {
		name    string
		mock    func(*mocking.MockParseNode)
		wantErr bool
		wantVal []serialization.Parsable
	}{
		{
			name: "Success",
			mock: func(m *mocking.MockParseNode) {
				m.On("GetCollectionOfObjectValues", mock.Anything).Return(val, nil)
			},
			wantErr: false,
			wantVal: val,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := mocking.NewMockParseNode()
			tt.mock(node)
			
			var result []serialization.Parsable
			setter := func(v []serialization.Parsable) error {
				result = v
				return nil
			}
			
			err := DeserializeCollectionOfObjectValuesFunc[serialization.Parsable](nil)(setter)(node)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantVal, result)
			node.AssertExpectations(t)
		})
	}
}

func TestDeserializeObjectValueFunc(t *testing.T) {
	parsable := mocking.NewMockParsable()
	
	tests := []struct {
		name    string
		mock    func(*mocking.MockParseNode)
		wantErr bool
		wantVal serialization.Parsable
	}{
		{
			name: "Success",
			mock: func(m *mocking.MockParseNode) {
				m.On("GetObjectValue", mock.Anything).Return(parsable, nil)
			},
			wantErr: false,
			wantVal: parsable,
		},
		{
			name: "Nil",
			mock: func(m *mocking.MockParseNode) {
				m.On("GetObjectValue", mock.Anything).Return(nil, nil)
			},
			wantErr: false,
			wantVal: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := mocking.NewMockParseNode()
			tt.mock(node)
			
			var result serialization.Parsable
			setter := func(v serialization.Parsable) error {
				result = v
				return nil
			}
			
			err := DeserializeObjectValueFunc[serialization.Parsable](nil)(setter)(node)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantVal, result)
			node.AssertExpectations(t)
		})
	}
}
