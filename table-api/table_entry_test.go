package tableapi

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const keyName = "key1"

func TestTableEntry(t *testing.T) {
	responseJSON, err := json.Marshal(fakeItemResult)
	assert.Nil(t, err)

	var entry TableEntry

	err = json.Unmarshal(responseJSON, &entry)
	assert.Nil(t, err)

	assert.Equal(t, fakeEntry, entry)
}

func TestNewTableEntry(t *testing.T) {
	entry := NewTableEntry()
	assert.IsType(t, tableEntry{}, entry)
	assert.IsType(t, TableEntry{}, entry)
}

func TestTableEntryValueValidKey(t *testing.T) {
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
}

func TestTableEntryValueMissingKey(t *testing.T) {
	entry := TableEntry{
		"key2": "value1",
	}

	value := entry.Value(keyName)

	assert.Nil(t, value)
}

func TestTableEntryKeys(t *testing.T) {
	entry := TableEntry{
		keyName: "value2",
		"key2":  "value1",
	}

	keys := entry.Keys()

	assert.Contains(t, keys, keyName)
	assert.Contains(t, keys, "key2")
}
