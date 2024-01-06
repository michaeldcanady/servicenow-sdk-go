package tableapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableValue2(t *testing.T) {

	tests := []test[*TableValue2]{
		{
			title: "",
			value: []byte(`{
                "link": "https://xxxxxxx.service-now.com/api/now/table/core_company/e540eeee1b22b8504b0dcb36464bcbb7",
                "value": "e540eeee1b22b8504b0dcb36464bcbb7"
            }`),
			expected: &TableValue2{
				Value:        &DataValue{value: "e540eeee1b22b8504b0dcb36464bcbb7"},
				Link:         "https://xxxxxxx.service-now.com/api/now/table/core_company/e540eeee1b22b8504b0dcb36464bcbb7",
				DisplayValue: nil,
			},
			expectErr: false,
			err:       nil,
		},
		{
			title: "",
			value: []byte(`{
                "link": "https://xxxxxxx.service-now.com/api/now/table/core_company/e540eeee1b22b8504b0dcb36464bcbb7",
                "display_value": "Lenovo"
            }`),
			expected: &TableValue2{
				DisplayValue: &DataValue{"Lenovo"},
				Link:         "https://xxxxxxx.service-now.com/api/now/table/core_company/e540eeee1b22b8504b0dcb36464bcbb7",
				Value:        nil,
			},
			expectErr: false,
			err:       nil,
		},
		{
			title: "",
			value: []byte(`"Lenovo"`),
			expected: &TableValue2{
				DisplayValue: nil,
				Link:         "",
				Value:        &DataValue{"Lenovo"},
			},
			expectErr: false,
			err:       nil,
		},
		{
			title:     "",
			value:     []byte(`\{"name": "Alice", "age": 25, 'gender': "female"}`),
			expected:  nil,
			expectErr: true,
			err:       nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {

			var value = &TableValue2{}

			err := value.Unmarshal(tt.value.([]byte))

			if tt.expectErr {
				assert.Error(t, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, value)
		})
	}
}

func TestTableValue2_isEmpty(t *testing.T) {
	tests := []test[bool]{
		{
			title:    "",
			value:    &TableValue2{},
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {

			got := tt.value.(*TableValue2).isEmpty()
			assert.Equal(t, tt.expected, got)
		})
	}
}
