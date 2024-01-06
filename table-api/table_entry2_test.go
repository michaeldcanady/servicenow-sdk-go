package tableapi

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Not functional
func TestTableEntry2_UnmarshallJSON(t *testing.T) {
	responseJSON, err := json.Marshal(fakeResultItem)
	assert.Nil(t, err)

	var entry TableEntry2

	err = json.Unmarshal(responseJSON, &entry)
	assert.Nil(t, err)

	assert.Equal(t, fakeEntry, entry)
}

func TestTableEntry2_MarshalJSON(t *testing.T) {

}

func TestNewTableEntry2(t *testing.T) {
	entry := NewTableEntry2()

	assert.IsType(t, &TableEntry2{}, entry)
}

func TestTableEntry2_Value(t *testing.T) {

	tests := []test[interface{}]{
		{
			title: "ValidKey_DisplayValueTrue_ExcludeReferenceLink",
			value: &TableEntry2{
				value: map[string]TableValue2{keyName: {
					DisplayValue: &DataValue{value: "Lenovo"},
				}},
			},
			expected: &TableValue2{
				Link:         "",
				Value:        nil,
				DisplayValue: &DataValue{value: "Lenovo"},
			},
			expectErr: false,
			err:       nil,
		},
		{
			title: "ValidKey",
			value: &TableEntry2{
				value: map[string]TableValue2{keyName: {
					Link:  "https://instance.servicenow.com/api/now/table/cmdb_ci/55b35562c0a8010e01cff22378e0aea9",
					Value: &DataValue{value: "55b35562c0a8010e01cff22378e0aea9"},
				}},
			},
			expected: &TableValue2{
				Link:  "https://instance.servicenow.com/api/now/table/cmdb_ci/55b35562c0a8010e01cff22378e0aea9",
				Value: &DataValue{value: "55b35562c0a8010e01cff22378e0aea9"},
			},
			expectErr: false,
			err:       nil,
		},
		{
			title: "ValidKey_DisplayValueTrue",
			value: &TableEntry2{
				value: map[string]TableValue2{keyName: {
					Link:         "https://instance.servicenow.com/api/now/table/cmdb_ci/55b35562c0a8010e01cff22378e0aea9",
					DisplayValue: &DataValue{value: "Lenovo"},
				}},
			},
			expected: &TableValue2{
				Link:         "https://instance.servicenow.com/api/now/table/cmdb_ci/55b35562c0a8010e01cff22378e0aea9",
				DisplayValue: &DataValue{value: "Lenovo"},
			},
			expectErr: false,
			err:       nil,
		},
		{
			title: "ValidKey_DisplayValueAll",
			value: &TableEntry2{
				value: map[string]TableValue2{keyName: {
					Link:         "https://instance.servicenow.com/api/now/table/cmdb_ci/55b35562c0a8010e01cff22378e0aea9",
					DisplayValue: &DataValue{value: "Lenovo"},
					Value:        &DataValue{value: "55b35562c0a8010e01cff22378e0aea9"},
				}},
			},
			expected: &TableValue2{
				Link:         "https://instance.servicenow.com/api/now/table/cmdb_ci/55b35562c0a8010e01cff22378e0aea9",
				DisplayValue: &DataValue{value: "Lenovo"},
				Value:        &DataValue{value: "55b35562c0a8010e01cff22378e0aea9"},
			},
			expectErr: false,
			err:       nil,
		},
	}

	for _, tt := range tests {

		value := tt.value.(*TableEntry2).Value(keyName)
		assert.NotNil(t, value)
		assert.Equal(t, value, tt.expected)
	}
}

func TestTableEntry2_Keys(t *testing.T) {
	entry := &TableEntry2{
		value: map[string]TableValue2{
			keyName: {
				Link:         "https://instance.servicenow.com/api/now/table/cmdb_ci/55b35562c0a8010e01cff22378e0aea9",
				DisplayValue: &DataValue{value: "Lenovo"},
				Value:        &DataValue{value: "55b35562c0a8010e01cff22378e0aea9"},
			},
			"key2": {
				DisplayValue: &DataValue{value: "Lenovo"},
				Value:        &DataValue{value: "55b35562c0a8010e01cff22378e0aea9"},
			},
		},
	}

	keys := entry.Keys()

	assert.Contains(t, keys, keyName)
	assert.Contains(t, keys, "key2")
}

func TestTableEntry2_Len(t *testing.T) {
	entry := &TableEntry2{
		value: map[string]TableValue2{
			keyName: {
				Link:         "https://instance.servicenow.com/api/now/table/cmdb_ci/55b35562c0a8010e01cff22378e0aea9",
				DisplayValue: &DataValue{value: "Lenovo"},
				Value:        &DataValue{value: "55b35562c0a8010e01cff22378e0aea9"},
			},
			"key2": {
				DisplayValue: &DataValue{value: "Lenovo"},
				Value:        &DataValue{value: "55b35562c0a8010e01cff22378e0aea9"},
			},
		},
	}

	assert.Equal(t, 2, entry.Len())
}
