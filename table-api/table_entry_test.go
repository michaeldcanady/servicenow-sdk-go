package tableapi

import (
	"encoding/json"
	"errors"
	"testing"

	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/stretchr/testify/assert"
)

const keyName = "key1"

func TestTableEntry(t *testing.T) {
	responseJSON, err := json.Marshal(fakeResultItem)
	assert.Nil(t, err)

	var entry TableEntry

	err = json.Unmarshal(responseJSON, &entry)
	assert.Nil(t, err)

	assert.Equal(t, fakeEntry, entry)
}

func TestNewTableEntry(t *testing.T) {
	entry := NewTableEntry()

	assert.IsType(t, TableEntry{}, entry)
}

func TestTableEntry_Set(t *testing.T) {
	entry := TableEntry{}

	entry.Set(keyName, "value2")

	assert.Equal(t, TableEntry{keyName: "value2"}, entry)
}

func TestTableEntry_Value(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "ValidKey",
			test: func(t *testing.T) {
				entry := TableEntry{
					keyName: "value1",
				}

				value := entry.Value(keyName)

				assert.NotNil(t, value)
				assert.Equal(t, value, &TableValue{value: "value1"})

				entry = TableEntry{
					keyName: map[string]interface{}{
						"link":  "https://instance.servicenow.com/api/now/table/cmdb_ci/55b35562c0a8010e01cff22378e0aea9",
						"value": "55b35562c0a8010e01cff22378e0aea9",
					},
				}

				value = entry.Value(keyName)

				assert.NotNil(t, value)
				assert.Equal(t, value, &TableValue{value: "55b35562c0a8010e01cff22378e0aea9"})
			},
		},
		{
			name: "MissingKey",
			test: func(t *testing.T) {
				entry := TableEntry{
					"key2": "value1",
				}

				value := entry.Value(keyName)

				assert.Nil(t, value)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestTableEntry_DisplayValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "ValidKey",
			test: func(t *testing.T) {
				entry := TableEntry{
					keyName: "value1",
				}

				value := entry.DisplayValue(keyName)

				assert.NotNil(t, value)
				assert.Equal(t, value, &TableValue{value: "value1"})

				entry = TableEntry{
					keyName: map[string]interface{}{
						linkKey:         "https://instance.servicenow.com/api/now/table/cmdb_ci/55b35562c0a8010e01cff22378e0aea9",
						displayValueKey: "55b35562c0a8010e01cff22378e0aea9",
					},
				}

				value = entry.DisplayValue(keyName)

				assert.NotNil(t, value)
				assert.Equal(t, value, &TableValue{value: "55b35562c0a8010e01cff22378e0aea9"})
			},
		},
		{
			name: "MissingKey",
			test: func(t *testing.T) {
				entry := TableEntry{
					"key2": "value1",
				}

				value := entry.DisplayValue(keyName)

				assert.Nil(t, value)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestTableEntry_Link(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "ValidKey",
			test: func(t *testing.T) {
				entry := TableEntry{
					keyName: "value1",
				}

				value, err := entry.Link(keyName)

				assert.Nil(t, value)
				assert.Nil(t, err)

				entry = TableEntry{
					keyName: map[string]interface{}{
						linkKey:         "https://instance.servicenow.com/api/now/table/cmdb_ci/55b35562c0a8010e01cff22378e0aea9",
						displayValueKey: "55b35562c0a8010e01cff22378e0aea9",
					},
				}

				value, err = entry.Link(keyName)

				assert.NotNil(t, value)
				assert.Equal(t, value, internal.ToPointer("https://instance.servicenow.com/api/now/table/cmdb_ci/55b35562c0a8010e01cff22378e0aea9"))
				assert.Nil(t, err)
			},
		},
		{
			name: "MissingKey",
			test: func(t *testing.T) {
				entry := TableEntry{
					"key2": "value1",
				}

				value, err := entry.Link(keyName)

				assert.Nil(t, value)
				assert.Nil(t, err)

				entry = TableEntry{
					keyName: map[string]interface{}{
						valueKey:        "55b35562c0a8010e01cff22378e0aea9",
						displayValueKey: "55b35562c0a8010e01cff22378e0aea9",
					},
				}

				value, err = entry.Link(keyName)

				assert.Nil(t, value)
				assert.Nil(t, err)
			},
		},
		{
			name: "Wrong link type",
			test: func(t *testing.T) {
				entry := TableEntry{
					keyName: map[string]interface{}{
						linkKey:         true,
						displayValueKey: "55b35562c0a8010e01cff22378e0aea9",
					},
				}

				value, err := entry.Link(keyName)

				assert.Nil(t, value)
				assert.Equal(t, errors.New("link is not string"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestTableEntry_Keys(t *testing.T) {
	entry := TableEntry{
		keyName: "value2",
		"key2":  "value1",
	}

	keys := entry.Keys()

	assert.Contains(t, keys, keyName)
	assert.Contains(t, keys, "key2")
}

func TestTableEntry_Len(t *testing.T) {
	entry := TableEntry{
		keyName: "value2",
		"key2":  "value1",
	}

	assert.Len(t, entry, 2)
}
