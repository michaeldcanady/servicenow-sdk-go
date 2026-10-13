// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableRecord_GetSysID(t *testing.T) {
	tests := []struct {
		name     string
		record   func() *TableRecord
		expected *string
		wantErr  bool
	}{
		{
			name: "returns the sys_id when set",
			record: func() *TableRecord {
				record := NewTableRecord()
				require.NoError(t, record.SetValue("sys_id", "d71f7935c0a8016700802b64c67c11c6"))

				return record
			},
			expected: internal.ToPointer("d71f7935c0a8016700802b64c67c11c6"),
		},
		{
			// Get returns a nil element for an absent key, and GetValue rejects a nil
			// receiver, so an absent sys_id surfaces as an error rather than a nil result.
			name: "errors when sys_id was never set",
			record: func() *TableRecord {
				return NewTableRecord()
			},
			wantErr: true,
		},
		{
			name: "returns nil when sys_id holds an explicit nil",
			record: func() *TableRecord {
				record := NewTableRecord()
				require.NoError(t, record.SetValue("sys_id", nil))

				return record
			},
			expected: nil,
		},
		{
			// GetStringValue rejects a value that is not a string, so a numeric sys_id surfaces
			// as an error rather than being coerced.
			name: "errors when sys_id is not a string",
			record: func() *TableRecord {
				record := NewTableRecord()
				require.NoError(t, record.SetValue("sys_id", int32(42)))

				return record
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sysID, err := test.record().GetSysID()

			if test.wantErr {
				require.Error(t, err)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, test.expected, sysID)
		})
	}
}

func TestRecordElementParserFromRaw_Gaps(t *testing.T) {
	t.Run("link of the wrong type is rejected", func(t *testing.T) {
		element, err := recordElementParserFromRaw(map[string]any{
			recordValueKey: "value",
			recordLinkKey:  "not-a-string-pointer",
		})

		require.ErrorContains(t, err, "link is not *string")
		assert.Nil(t, element)
	})

	t.Run("map with display value, value and link", func(t *testing.T) {
		element, err := recordElementParserFromRaw(map[string]any{
			recordDisplayValueKey: "Display",
			recordValueKey:        "raw",
			recordLinkKey:         internal.ToPointer("https://example.com/link"),
		})

		require.NoError(t, err)
		require.NotNil(t, element)

		link, err := element.GetLink()
		require.NoError(t, err)
		require.NotNil(t, link)
		assert.Equal(t, "https://example.com/link", *link)
	})
}

// mustElementValue builds an ElementValue for a test table, failing the test if the visitor
// cannot handle the value.
func mustElementValue(t *testing.T, val any) *ElementValue {
	t.Helper()

	elementValue, err := NewElementValue(val)
	require.NoError(t, err)

	return elementValue
}

func TestElementValue_GetCollectionOfPrimitiveValues_Gaps(t *testing.T) {
	tests := []struct {
		name       string
		value      *ElementValue
		targetType Primitive
		expected   []any
		wantErrMsg string
	}{
		{
			name:       "nil element value returns nil",
			value:      nil,
			targetType: PrimitiveString,
			expected:   nil,
		},
		{
			name:       "unknown target type is rejected",
			value:      mustElementValue(t, "x"),
			targetType: PrimitiveUnknown,
			wantErrMsg: "target type can't be",
		},
		{
			name:       "non-collection value is rejected",
			value:      mustElementValue(t, "x"),
			targetType: PrimitiveString,
			wantErrMsg: "val is not a collection",
		},
		{
			name:       "collection of strings is converted element by element",
			value:      mustElementValue(t, []any{internal.ToPointer("a"), internal.ToPointer("b")}),
			targetType: PrimitiveString,
			expected:   []any{internal.ToPointer("a"), internal.ToPointer("b")},
		},
		{
			// A nil entry in the source collection survives as a nil slot rather than
			// short-circuiting the whole conversion.
			name:       "nil entries become nil slots",
			value:      mustElementValue(t, []any{internal.ToPointer("a"), nil}),
			targetType: PrimitiveString,
			expected:   []any{internal.ToPointer("a"), nil},
		},
		{
			// An entry that is not the requested primitive fails the whole conversion.
			name:       "entry of the wrong primitive type fails the conversion",
			value:      mustElementValue(t, []any{internal.ToPointer("a"), internal.ToPointer(true)}),
			targetType: PrimitiveString,
			wantErrMsg: "not",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			collection, err := test.value.GetCollectionOfPrimitiveValues(test.targetType)

			if test.wantErrMsg != "" {
				require.ErrorContains(t, err, test.wantErrMsg)
				assert.Nil(t, collection)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, test.expected, collection)
		})
	}
}
