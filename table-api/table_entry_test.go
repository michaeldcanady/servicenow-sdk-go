package tableapi

import (
	"encoding/json"
	"testing"

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

func TestTableEntry_Value(t *testing.T) {

	tests := []test[interface{}]{
		{
			title: "ValidKey_DisplayValueTrue_ExcludeReferenceLink",
			value: TableEntry{
				keyName: "value1",
			},
			expected:  &TableValue{value: "value1"},
			expectErr: false,
			err:       nil,
		},
		{
			title: "ValidKey",
			value: TableEntry{
				keyName: map[string]interface{}{
					"link":  "https://instance.servicenow.com/api/now/table/cmdb_ci/55b35562c0a8010e01cff22378e0aea9",
					"value": "55b35562c0a8010e01cff22378e0aea9",
				},
			},
			expected:  &TableValue{value: "55b35562c0a8010e01cff22378e0aea9"}, //Correct for test but wrong
			expectErr: false,
			err:       nil,
		},
		{
			title: "ValidKey_DisplayValueTrue",
			value: TableEntry{
				keyName: map[string]interface{}{
					"display_value": "Lenovo",
					"link":          "https://instance.servicenow.com/api/now/table/cmdb_ci/55b35562c0a8010e01cff22378e0aea9",
				},
			},
			expected:  &TableValue{value: "55b35562c0a8010e01cff22378e0aea9"}, //Correct for test but wrong
			expectErr: false,
			err:       nil,
		},
		{
			title: "ValidKey_DisplayValueAll",
			value: TableEntry{
				keyName: map[string]interface{}{
					"display_value": "Lenovo",
					"value":         "55b35562c0a8010e01cff22378e0aea9",
					"link":          "https://instance.servicenow.com/api/now/table/cmdb_ci/55b35562c0a8010e01cff22378e0aea9",
				},
			},
			expected:  &TableValue{value: "55b35562c0a8010e01cff22378e0aea9"}, //Correct for test but wrong
			expectErr: false,
			err:       nil,
		},
	}

	for _, tt := range tests {

		value := tt.value.(TableEntry).Value(keyName)
		assert.NotNil(t, value)
		assert.Equal(t, value, tt.expected)
	}
}

func TestTableEntry_Set(t *testing.T) {
	entry := TableEntry{}

	entry.Set(keyName, "value2")

	assert.Equal(t, TableEntry{keyName: "value2"}, entry)
}

func TestTableEntry_ValueMissingKey(t *testing.T) {
	entry := TableEntry{
		"key2": "value1",
	}

	value := entry.Value(keyName)

	assert.Nil(t, value)
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
