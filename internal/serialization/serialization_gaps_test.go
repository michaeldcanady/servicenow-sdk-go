package serialization

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// errRead is a stand-in for a read failure from a parse node.
var errRead = errors.New("read failed")

// ---------------------------------------------------------------------------
// Previously untested primitive deserializers
// ---------------------------------------------------------------------------

func TestDeserializeTimeFunc(t *testing.T) {
	expected := time.Date(2026, 7, 29, 12, 30, 0, 0, time.UTC)

	t.Run("value is read", func(t *testing.T) {
		node := mocking.NewMockParseNode()
		node.On("GetTimeValue").Return(&expected, nil)

		var result *time.Time
		require.NoError(t, DeserializeTimeFunc(func(v *time.Time) error { result = v; return nil })(node))

		require.NotNil(t, result)
		assert.Equal(t, expected, *result)
	})

	t.Run("read error propagates", func(t *testing.T) {
		node := mocking.NewMockParseNode()
		node.On("GetTimeValue").Return((*time.Time)(nil), errRead)

		require.ErrorIs(t, DeserializeTimeFunc(func(*time.Time) error { return nil })(node), errRead)
	})
}

func TestDeserializeTimeOnlyFunc(t *testing.T) {
	expected, err := serialization.ParseTimeOnly("12:30:00")
	require.NoError(t, err)

	t.Run("value is read", func(t *testing.T) {
		node := mocking.NewMockParseNode()
		node.On("GetTimeOnlyValue").Return(expected, nil)

		var result *serialization.TimeOnly
		require.NoError(t, DeserializeTimeOnlyFunc(func(v *serialization.TimeOnly) error { result = v; return nil })(node))

		assert.Equal(t, expected, result)
	})

	t.Run("read error propagates", func(t *testing.T) {
		node := mocking.NewMockParseNode()
		node.On("GetTimeOnlyValue").Return((*serialization.TimeOnly)(nil), errRead)

		require.ErrorIs(t, DeserializeTimeOnlyFunc(func(*serialization.TimeOnly) error { return nil })(node), errRead)
	})
}

func TestDeserializeDateOnlyFunc(t *testing.T) {
	expected, err := serialization.ParseDateOnly("2026-07-29")
	require.NoError(t, err)

	t.Run("value is read", func(t *testing.T) {
		node := mocking.NewMockParseNode()
		node.On("GetDateOnlyValue").Return(expected, nil)

		var result *serialization.DateOnly
		require.NoError(t, DeserializeDateOnlyFunc(func(v *serialization.DateOnly) error { result = v; return nil })(node))

		assert.Equal(t, expected, result)
	})

	t.Run("read error propagates", func(t *testing.T) {
		node := mocking.NewMockParseNode()
		node.On("GetDateOnlyValue").Return((*serialization.DateOnly)(nil), errRead)

		require.ErrorIs(t, DeserializeDateOnlyFunc(func(*serialization.DateOnly) error { return nil })(node), errRead)
	})
}

func TestDeserializeUUIDFunc(t *testing.T) {
	expected := uuid.MustParse("d71f7935-c0a8-0167-0080-2b64c67c11c6")

	t.Run("value is read", func(t *testing.T) {
		node := mocking.NewMockParseNode()
		node.On("GetUUIDValue").Return(&expected, nil)

		var result *uuid.UUID
		require.NoError(t, DeserializeUUIDFunc(func(v *uuid.UUID) error { result = v; return nil })(node))

		require.NotNil(t, result)
		assert.Equal(t, expected, *result)
	})

	t.Run("read error propagates", func(t *testing.T) {
		node := mocking.NewMockParseNode()
		node.On("GetUUIDValue").Return((*uuid.UUID)(nil), errRead)

		require.ErrorIs(t, DeserializeUUIDFunc(func(*uuid.UUID) error { return nil })(node), errRead)
	})
}

// ---------------------------------------------------------------------------
// Deserializer error branches
// ---------------------------------------------------------------------------

func TestDeserializeCollectionOfObjectValuesFunc_ReadError(t *testing.T) {
	node := mocking.NewMockParseNode()
	node.On("GetCollectionOfObjectValues", mock.Anything).Return([]serialization.Parsable(nil), errRead)

	err := DeserializeCollectionOfObjectValuesFunc[*mocking.MockParsable](
		nil,
		func([]*mocking.MockParsable) error { return nil },
	)(node)

	require.ErrorIs(t, err, errRead)
}

func TestDeserializeObjectValueFunc_ReadError(t *testing.T) {
	node := mocking.NewMockParseNode()
	node.On("GetObjectValue", mock.Anything).Return(nil, errRead)

	err := DeserializeObjectValueFunc[*mocking.MockParsable](
		nil,
		func(*mocking.MockParsable) error { return nil },
	)(node)

	require.ErrorIs(t, err, errRead)
}

func TestDeserializeEnumFunc_Failures(t *testing.T) {
	t.Run("read error propagates", func(t *testing.T) {
		node := mocking.NewMockParseNode()
		node.On("GetEnumValue", mock.Anything).Return(nil, errRead)

		err := DeserializeEnumFunc[mockEnum](nil, func(*mockEnum) error { return nil })(node)

		require.ErrorIs(t, err, errRead)
	})

	t.Run("a value that is neither T nor *T is rejected", func(t *testing.T) {
		node := mocking.NewMockParseNode()
		node.On("GetEnumValue", mock.Anything).Return("not-an-enum", nil)

		err := DeserializeEnumFunc[mockEnum](nil, func(*mockEnum) error { return nil })(node)

		require.ErrorContains(t, err, "unexpected type from enum factory")
	})
}

// ---------------------------------------------------------------------------
// Serializer nil-skip branches
// ---------------------------------------------------------------------------

// TestSerializeMutatedStringFunc_NilMutatedValue covers the branch where the mutator yields nil:
// nothing is written at all, rather than an empty string.
func TestSerializeMutatedStringFunc_NilMutatedValue(t *testing.T) {
	writer := mocking.NewMockSerializationWriter()

	err := SerializeMutatedStringFunc("key",
		func(int) (*string, error) { return nil, nil },
		func() (int, error) { return 0, nil },
	)(writer)

	require.NoError(t, err)
	writer.AssertNotCalled(t, "WriteStringValue", mock.Anything, mock.Anything)
}

// TestSerializeStringToXFunc_NilSkips covers the nil-value branch of each "serialize as string"
// helper. A nil source value must produce no write at all.
func TestSerializeStringToXFunc_NilSkips(t *testing.T) {
	tests := []struct {
		name      string
		serialize func() WriterFunc
	}{
		{
			name: "SerializeStringToBoolFunc",
			serialize: func() WriterFunc {
				return SerializeStringToBoolFunc("key", func() (*bool, error) { return nil, nil })
			},
		},
		{
			name: "SerializeStringToFloat64Func",
			serialize: func() WriterFunc {
				return SerializeStringToFloat64Func("key", func() (*float64, error) { return nil, nil })
			},
		},
		{
			name: "SerializeStringToInt64Func",
			serialize: func() WriterFunc {
				return SerializeStringToInt64Func("key", func() (*int64, error) { return nil, nil })
			},
		},
		{
			name: "SerializeStringToTimeFunc",
			serialize: func() WriterFunc {
				return SerializeStringToTimeFunc("key", time.RFC3339, func() (*time.Time, error) { return nil, nil })
			},
		},
		{
			name: "SerializeStringToSliceFunc",
			serialize: func() WriterFunc {
				return SerializeStringToSliceFunc("key", ",", func() ([]string, error) { return nil, nil })
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()

			require.NoError(t, test.serialize()(writer))

			writer.AssertNotCalled(t, "WriteStringValue", mock.Anything, mock.Anything)
		})
	}
}

// TestSerializeCollectionOfObjectValuesFunc_NilSkips covers the nil-collection branch.
func TestSerializeCollectionOfObjectValuesFunc_NilSkips(t *testing.T) {
	writer := mocking.NewMockSerializationWriter()

	err := SerializeCollectionOfObjectValuesFunc[*mocking.MockParsable]("key",
		func() ([]*mocking.MockParsable, error) { return nil, nil },
	)(writer)

	require.NoError(t, err)
	writer.AssertNotCalled(t, "WriteCollectionOfObjectValues", mock.Anything, mock.Anything)
}

// TestSerializeEnumFunc_NilSkips covers the nil-enum branch, where the mutator short-circuits
// before ever calling String() on a nil pointer.
func TestSerializeEnumFunc_NilSkips(t *testing.T) {
	writer := mocking.NewMockSerializationWriter()

	err := SerializeEnumFunc[mockEnum, *mockEnum]("key",
		func() (*mockEnum, error) { return nil, nil },
	)(writer)

	require.NoError(t, err)
	writer.AssertNotCalled(t, "WriteStringValue", mock.Anything, mock.Anything)
}
