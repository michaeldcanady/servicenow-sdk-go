package tableapi

import (
	"encoding/json"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
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
	entry := TableEntry{
		keyName: "value1",
	}

	tests := []internal.Test[*TableValue]{
		{
			Title:    "ValidKey",
			Input:    keyName,
			Expected: &TableValue{value: "value1"},
		},
		{
			Title: "ValidKey",
			Prepare: func() {
				entry = TableEntry{
					keyName: map[string]interface{}{
						"link":  "https://instance.servicenow.com/api/now/table/cmdb_ci/55b35562c0a8010e01cff22378e0aea9",
						"value": "55b35562c0a8010e01cff22378e0aea9",
					},
				}
			},
			Input:    keyName,
			Expected: &TableValue{value: "55b35562c0a8010e01cff22378e0aea9"},
		},
		{
			Title: "MissingKey",
			Prepare: func() {
				entry = TableEntry{
					"key2": "value1",
				}
			},
			Input: keyName,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Title, func(t *testing.T) {
			if tt.Prepare != nil {
				tt.Prepare()
			}

			value := entry.Value(tt.Input.(string))

			assert.Equal(t, tt.Expected, value)
		})
	}
}

func TestTableEntry_Set(t *testing.T) {
	entry := TableEntry{}

	tests := []internal.Test[TableEntry]{
		{
			Title:    "",
			Input:    []interface{}{keyName, "value2"},
			Expected: TableEntry{keyName: "value2"},
			Cleanup: func() {
				entry = TableEntry{}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Title, func(t *testing.T) {
			if tt.Prepare != nil {
				tt.Prepare()
			}

			keyName := tt.Input.([]interface{})[0]
			value := tt.Input.([]interface{})[1]

			entry.Set(keyName.(string), value)

			assert.Equal(t, tt.Expected, entry)

			if tt.Cleanup != nil {
				tt.Cleanup()
			}
		})
	}
}

func TestTableEntry_Keys(t *testing.T) {
	// Define a slice of test cases
	tests := []internal.Test[[]string]{
		{
			Title:    "OneKey",
			Input:    TableEntry{keyName: "value2"},
			Expected: []string{keyName},
		},
		{
			Title:    "TwoKeys",
			Input:    TableEntry{keyName: "value2", "key2": "value1"},
			Expected: []string{keyName, "key2"},
		},
	}

	// Iterate over the test cases
	for _, tt := range tests {
		t.Run(tt.Title, func(t *testing.T) {
			// Call the prepare function if it exists
			if tt.Prepare != nil {
				tt.Prepare()
			}

			// Call the function under test with the input
			keys := tt.Input.(TableEntry).Keys()

			// Compare the output with the expected value
			for _, key := range tt.Expected {
				assert.Contains(t, keys, key)
			}

			// Call the cleanup function if it exists
			if tt.Cleanup != nil {
				tt.Cleanup()
			}
		})
	}
}

func TestTableEntry_Len(t *testing.T) {
	// Define a slice of test cases
	tests := []internal.Test[int]{
		{
			Title:    "EmptyEntry",
			Input:    TableEntry{},
			Expected: 0,
		},
		{
			Title:    "OneKeyEntry",
			Input:    TableEntry{keyName: "value2"},
			Expected: 1,
		},
		{
			Title:    "TwoKeysEntry",
			Input:    TableEntry{keyName: "value2", "key2": "value1"},
			Expected: 2,
		},
	}

	// Iterate over the test cases
	for _, tt := range tests {
		t.Run(tt.Title, func(t *testing.T) {
			// Call the prepare function if it exists
			if tt.Prepare != nil {
				tt.Prepare()
			}

			// Call the function under test with the input
			length := tt.Input.(TableEntry).Len()

			// Compare the output with the expected value
			assert.Equal(t, length, tt.Expected)

			// Call the cleanup function if it exists
			if tt.Cleanup != nil {
				tt.Cleanup()
			}
		})
	}
}
