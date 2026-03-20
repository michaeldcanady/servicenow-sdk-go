package kiota

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeserializeMutatedStringFunc(t *testing.T) {
	val := "123"
	mutated := 123
	tests := []struct {
		name    string
		mock    func(*mocking.MockParseNode)
		mutator Mutator[*string, int]
		wantErr bool
		wantVal int
	}{
		{
			name: "Success",
			mock: func(m *mocking.MockParseNode) {
				m.On("GetStringValue").Return(&val, nil)
			},
			mutator: func(s *string) (int, error) { return mutated, nil },
			wantErr: false,
			wantVal: mutated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := mocking.NewMockParseNode()
			tt.mock(node)
			var result int
			setter := func(v int) error {
				result = v
				return nil
			}
			err := DeserializeMutatedStringFunc(tt.mutator)(setter)(node)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantVal, result)
		})
	}
}

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

			err := DeserializeStringFunc(setter)(node)
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

			err := DeserializeInt64Func(setter)(node)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantVal, result)
			node.AssertExpectations(t)
		})
	}
}

func TestDeserializeInt32Func(t *testing.T) {
	val := int32(123)
	tests := []struct {
		name    string
		mock    func(*mocking.MockParseNode)
		wantErr bool
		wantVal *int32
	}{
		{
			name: "Success",
			mock: func(m *mocking.MockParseNode) {
				m.On("GetInt32Value").Return(&val, nil)
			},
			wantErr: false,
			wantVal: &val,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := mocking.NewMockParseNode()
			tt.mock(node)
			var result *int32
			setter := func(v *int32) error {
				result = v
				return nil
			}
			err := DeserializeInt32Func()(setter)(node)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantVal, result)
		})
	}
}

func TestDeserializeBoolFunc(t *testing.T) {
	val := true
	tests := []struct {
		name    string
		mock    func(*mocking.MockParseNode)
		wantErr bool
		wantVal *bool
	}{
		{
			name: "Success",
			mock: func(m *mocking.MockParseNode) {
				m.On("GetBoolValue").Return(&val, nil)
			},
			wantErr: false,
			wantVal: &val,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := mocking.NewMockParseNode()
			tt.mock(node)
			var result *bool
			setter := func(v *bool) error {
				result = v
				return nil
			}
			err := DeserializeBoolFunc()(setter)(node)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantVal, result)
		})
	}
}

func TestDeserializeFloat64Func(t *testing.T) {
	val := float64(1.23)
	tests := []struct {
		name    string
		mock    func(*mocking.MockParseNode)
		wantErr bool
		wantVal *float64
	}{
		{
			name: "Success",
			mock: func(m *mocking.MockParseNode) {
				m.On("GetFloat64Value").Return(&val, nil)
			},
			wantErr: false,
			wantVal: &val,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := mocking.NewMockParseNode()
			tt.mock(node)
			var result *float64
			setter := func(v *float64) error {
				result = v
				return nil
			}
			err := DeserializeFloat64Func()(setter)(node)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantVal, result)
		})
	}
}

func TestDeserializeFloat32Func(t *testing.T) {
	val := float32(1.23)
	tests := []struct {
		name    string
		mock    func(*mocking.MockParseNode)
		wantErr bool
		wantVal *float32
	}{
		{
			name: "Success",
			mock: func(m *mocking.MockParseNode) {
				m.On("GetFloat32Value").Return(&val, nil)
			},
			wantErr: false,
			wantVal: &val,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := mocking.NewMockParseNode()
			tt.mock(node)
			var result *float32
			setter := func(v *float32) error {
				result = v
				return nil
			}
			err := DeserializeFloat32Func()(setter)(node)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantVal, result)
		})
	}
}

func TestDeserializeByteArrayFunc(t *testing.T) {
	val := []byte{1, 2, 3}
	tests := []struct {
		name    string
		mock    func(*mocking.MockParseNode)
		wantErr bool
		wantVal []byte
	}{
		{
			name: "Success",
			mock: func(m *mocking.MockParseNode) {
				m.On("GetByteArrayValue").Return(val, nil)
			},
			wantErr: false,
			wantVal: val,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := mocking.NewMockParseNode()
			tt.mock(node)
			var result []byte
			setter := func(v []byte) error {
				result = v
				return nil
			}
			err := DeserializeByteArrayFunc()(setter)(node)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantVal, result)
		})
	}
}

func TestDeserializeMutatedByteArrayFunc(t *testing.T) {
	val := []byte{1, 2, 3}
	mutated := "123"
	tests := []struct {
		name    string
		mock    func(*mocking.MockParseNode)
		mutator Mutator[[]byte, string]
		wantErr bool
		wantVal string
	}{
		{
			name: "Success",
			mock: func(m *mocking.MockParseNode) {
				m.On("GetByteArrayValue").Return(val, nil)
			},
			mutator: func(b []byte) (string, error) { return mutated, nil },
			wantErr: false,
			wantVal: mutated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := mocking.NewMockParseNode()
			tt.mock(node)
			var result string
			setter := func(v string) error {
				result = v
				return nil
			}
			err := DeserializeMutatedByteArrayFunc(tt.mutator)(setter)(node)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantVal, result)
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

func TestDeserializeEnumFunc(t *testing.T) {
	enumVal := mockEnumVal
	tests := []struct {
		name    string
		mock    func(*mocking.MockParseNode)
		wantErr bool
		wantVal *mockEnum
	}{
		{
			name: "Success Value",
			mock: func(m *mocking.MockParseNode) {
				m.On("GetEnumValue", mock.Anything).Return(enumVal, nil)
			},
			wantErr: false,
			wantVal: &enumVal,
		},
		{
			name: "Success Pointer",
			mock: func(m *mocking.MockParseNode) {
				m.On("GetEnumValue", mock.Anything).Return(&enumVal, nil)
			},
			wantErr: false,
			wantVal: &enumVal,
		},
		{
			name: "Nil",
			mock: func(m *mocking.MockParseNode) {
				m.On("GetEnumValue", mock.Anything).Return(nil, nil)
			},
			wantErr: false,
			wantVal: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := mocking.NewMockParseNode()
			tt.mock(node)
			var result *mockEnum
			setter := func(v *mockEnum) error {
				result = v
				return nil
			}
			err := DeserializeEnumFunc[mockEnum](nil)(setter)(node)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantVal, result)
		})
	}
}

func TestDeserializeAnyFunc(t *testing.T) {
	val := "any"
	tests := []struct {
		name    string
		mock    func(*mocking.MockParseNode)
		wantErr bool
		wantVal any
	}{
		{
			name: "Success",
			mock: func(m *mocking.MockParseNode) {
				m.On("GetRawValue").Return(val, nil)
			},
			wantErr: false,
			wantVal: val,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := mocking.NewMockParseNode()
			tt.mock(node)
			var result any
			setter := func(v any) error {
				result = v
				return nil
			}
			err := DeserializeAnyFunc()(setter)(node)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantVal, result)
		})
	}
}

func TestDeserializeMutatedAnyFunc(t *testing.T) {
	val := "123"
	mutated := 123
	tests := []struct {
		name    string
		mock    func(*mocking.MockParseNode)
		mutator Mutator[any, int]
		wantErr bool
		wantVal int
	}{
		{
			name: "Success",
			mock: func(m *mocking.MockParseNode) {
				m.On("GetRawValue").Return(val, nil)
			},
			mutator: func(a any) (int, error) { return mutated, nil },
			wantErr: false,
			wantVal: mutated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := mocking.NewMockParseNode()
			tt.mock(node)
			var result int
			setter := func(v int) error {
				result = v
				return nil
			}
			err := DeserializeMutatedAnyFunc(tt.mutator)(setter)(node)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantVal, result)
		})
	}
}

func TestDeserializeISODurationFunc(t *testing.T) {
	val := serialization.NewDuration(0, 0, 0, 0, 0, 1, 0)
	tests := []struct {
		name    string
		mock    func(*mocking.MockParseNode)
		wantErr bool
		wantVal *serialization.ISODuration
	}{
		{
			name: "Success",
			mock: func(m *mocking.MockParseNode) {
				m.On("GetISODurationValue").Return(val, nil)
			},
			wantErr: false,
			wantVal: val,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := mocking.NewMockParseNode()
			tt.mock(node)
			var result *serialization.ISODuration
			setter := func(v *serialization.ISODuration) error {
				result = v
				return nil
			}
			err := DeserializeISODurationFunc()(setter)(node)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantVal, result)
		})
	}
}
