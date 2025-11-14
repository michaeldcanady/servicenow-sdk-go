package tableapi

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTableRecordFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				parseNode := mocking.NewMockParseNode()
				parseNode.On("GetRawValue").Return(map[string]any{"key1": true}, nil)

				parsable, err := CreateTableRecordFromDiscriminatorValue(parseNode)

				assert.NotNil(t, parsable)
				assert.Nil(t, err)
				parseNode.AssertExpectations(t)
			},
		},
		{
			name: "Value retrieval error",
			test: func(t *testing.T) {
				parseNode := mocking.NewMockParseNode()
				parseNode.On("GetRawValue").Return(nil, errors.New("retrieval error"))

				parsable, err := CreateTableRecordFromDiscriminatorValue(parseNode)

				assert.Nil(t, parsable)
				assert.Equal(t, errors.New("retrieval error"), err)
				parseNode.AssertExpectations(t)
			},
		},
		{
			name: "Unsupported type",
			test: func(t *testing.T) {
				parseNode := mocking.NewMockParseNode()
				parseNode.On("GetRawValue").Return("key1", nil)

				parsable, err := CreateTableRecordFromDiscriminatorValue(parseNode)

				assert.Nil(t, parsable)
				assert.Equal(t, errors.New("unsupported type string"), err)
				parseNode.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRecordElementParser(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestNewTableRecord(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func keys[Map ~map[K]V, K comparable, V any](m Map) []K {
	keys := make([]K, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	return keys
}

func TestTableRecord_GetFieldDeserializers(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "key verification",
			test: func(t *testing.T) {
				expectedKeys := []string{"test", "test1"}
				record := &TableRecord{
					keys: expectedKeys,
				}

				deserializers := record.GetFieldDeserializers()
				assert.Equal(t, expectedKeys, keys(deserializers))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestTableRecord_Serialize(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "error",
			test: func(t *testing.T) {
				writer := mocking.NewMockSerializationWriter()

				record := (*TableRecord)(nil)

				err := record.Serialize(writer)

				assert.Equal(t, errors.New("unimplemented"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestTableRecord_Get(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "",
			test: func(t *testing.T) {
				result := NewRecordElement()

				store := mocking.NewMockBackingStore()
				store.On("Get", "Test").Return(result, nil)

				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return(store)

				record := &TableRecord{
					keys:  make([]string, 0),
					Model: innerModel,
				}

				elem, err := record.Get("Test")

				assert.Nil(t, err)
				assert.Equal(t, result, elem)
				innerModel.AssertExpectations(t)
				store.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestTableRecord_SetElement(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := NewRecordElement()

				store := mocking.NewMockBackingStore()
				store.On("Set", "Test", input).Return(nil)

				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return(store)

				record := &TableRecord{
					keys:  make([]string, 0),
					Model: innerModel,
				}

				err := record.SetElement("Test", input)

				assert.Nil(t, err)
				innerModel.AssertExpectations(t)
				store.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestTableRecord_SetValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				store.On("Set", "Test", mock.AnythingOfType("*tableapi.RecordElement")).Return(nil)

				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return(store)

				record := &TableRecord{
					keys:  make([]string, 0),
					Model: innerModel,
				}

				err := record.SetValue("Test", "value")

				assert.Nil(t, err)
				innerModel.AssertExpectations(t)
				store.AssertExpectations(t)
			},
		},
		{
			name: "Unsupported",
			test: func(t *testing.T) {
				record := &TableRecord{
					keys:  make([]string, 0),
					Model: nil,
				}

				err := record.SetValue("Test", make(chan int))

				assert.Equal(t, errors.New("unsupported kind chan"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestTableRecord_HasAttribute(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "key exists",
			test: func(t *testing.T) {
				record := &TableRecord{
					keys: []string{"key1"},
				}
				assert.True(t, record.HasAttribute("key1"))
			},
		},
		{
			name: "key doesn't exist",
			test: func(t *testing.T) {
				record := &TableRecord{
					keys: []string{"key1"},
				}
				assert.False(t, record.HasAttribute("key2"))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
