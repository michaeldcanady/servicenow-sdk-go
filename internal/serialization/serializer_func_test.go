package serialization

import (
	"errors"
	"testing"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSerialize(t *testing.T) {
	writer := mocking.NewMockSerializationWriter()

	s1 := func(sw serialization.SerializationWriter) error { return nil }
	s2 := func(sw serialization.SerializationWriter) error { return errors.New("err") }

	err := Serialize(writer, s1)
	assert.NoError(t, err)

	err = Serialize(writer, s1, s2)
	assert.Error(t, err)
}

func TestSerializeMutatedStringFunc(t *testing.T) {
	val := 123
	mutated := "123"
	tests := []struct {
		name     string
		accessor ModelAccessor[int]
		mutator  Mutator[int, *string]
		mock     func(*mocking.MockSerializationWriter)
		wantErr  bool
	}{
		{
			name:     "Success",
			accessor: func() (int, error) { return val, nil },
			mutator:  func(v int) (*string, error) { return &mutated, nil },
			mock: func(m *mocking.MockSerializationWriter) {
				m.On("WriteStringValue", "key", &mutated).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			tt.mock(writer)
			err := SerializeMutatedStringFunc("key", tt.mutator)(tt.accessor)(writer)
			assert.NoError(t, err)
			writer.AssertExpectations(t)
		})
	}
}

func TestSerializeStringFunc(t *testing.T) {
	val := "test"
	tests := []struct {
		name     string
		accessor ModelAccessor[*string]
		mock     func(*mocking.MockSerializationWriter)
		wantErr  bool
	}{
		{
			name:     "Success",
			accessor: func() (*string, error) { return &val, nil },
			mock: func(m *mocking.MockSerializationWriter) {
				m.On("WriteStringValue", "key", &val).Return(nil)
			},
			wantErr: false,
		},
		{
			name:     "Nil value",
			accessor: func() (*string, error) { return nil, nil },
			mock:     func(m *mocking.MockSerializationWriter) {},
			wantErr:  false,
		},
		{
			name:     "Error",
			accessor: func() (*string, error) { return nil, errors.New("err") },
			mock:     func(m *mocking.MockSerializationWriter) {},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			tt.mock(writer)

			err := SerializeStringFunc("key")(tt.accessor)(writer)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			writer.AssertExpectations(t)
		})
	}
}

func TestSerializeBoolFunc(t *testing.T) {
	val := true
	tests := []struct {
		name     string
		accessor ModelAccessor[*bool]
		mock     func(*mocking.MockSerializationWriter)
		wantErr  bool
	}{
		{
			name:     "Success",
			accessor: func() (*bool, error) { return &val, nil },
			mock: func(m *mocking.MockSerializationWriter) {
				m.On("WriteBoolValue", "key", &val).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			tt.mock(writer)

			err := SerializeBoolFunc("key")(tt.accessor)(writer)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			writer.AssertExpectations(t)
		})
	}
}

func TestSerializeInt64Func(t *testing.T) {
	val := int64(123)
	tests := []struct {
		name     string
		accessor ModelAccessor[*int64]
		mock     func(*mocking.MockSerializationWriter)
		wantErr  bool
	}{
		{
			name:     "Success",
			accessor: func() (*int64, error) { return &val, nil },
			mock: func(m *mocking.MockSerializationWriter) {
				m.On("WriteInt64Value", "key", &val).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			tt.mock(writer)

			err := SerializeInt64Func("key")(tt.accessor)(writer)
			assert.NoError(t, err)
			writer.AssertExpectations(t)
		})
	}
}

func TestSerializeInt32Func(t *testing.T) {
	val := int32(123)
	tests := []struct {
		name     string
		accessor ModelAccessor[*int32]
		mock     func(*mocking.MockSerializationWriter)
		wantErr  bool
	}{
		{
			name:     "Success",
			accessor: func() (*int32, error) { return &val, nil },
			mock: func(m *mocking.MockSerializationWriter) {
				m.On("WriteInt32Value", "key", &val).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			tt.mock(writer)

			err := SerializeInt32Func("key")(tt.accessor)(writer)
			assert.NoError(t, err)
			writer.AssertExpectations(t)
		})
	}
}

func TestSerializeFloat64Func(t *testing.T) {
	val := float64(1.23)
	tests := []struct {
		name     string
		accessor ModelAccessor[*float64]
		mock     func(*mocking.MockSerializationWriter)
		wantErr  bool
	}{
		{
			name:     "Success",
			accessor: func() (*float64, error) { return &val, nil },
			mock: func(m *mocking.MockSerializationWriter) {
				m.On("WriteFloat64Value", "key", &val).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			tt.mock(writer)

			err := SerializeFloat64Func("key")(tt.accessor)(writer)
			assert.NoError(t, err)
			writer.AssertExpectations(t)
		})
	}
}

func TestSerializeFloat32Func(t *testing.T) {
	val := float32(1.23)
	tests := []struct {
		name     string
		accessor ModelAccessor[*float32]
		mock     func(*mocking.MockSerializationWriter)
		wantErr  bool
	}{
		{
			name:     "Success",
			accessor: func() (*float32, error) { return &val, nil },
			mock: func(m *mocking.MockSerializationWriter) {
				m.On("WriteFloat32Value", "key", &val).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			tt.mock(writer)

			err := SerializeFloat32Func("key")(tt.accessor)(writer)
			assert.NoError(t, err)
			writer.AssertExpectations(t)
		})
	}
}

func TestSerializeTimeFunc(t *testing.T) {
	val := time.Now()
	tests := []struct {
		name     string
		accessor ModelAccessor[*time.Time]
		mock     func(*mocking.MockSerializationWriter)
		wantErr  bool
	}{
		{
			name:     "Success",
			accessor: func() (*time.Time, error) { return &val, nil },
			mock: func(m *mocking.MockSerializationWriter) {
				m.On("WriteTimeValue", "key", &val).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			tt.mock(writer)

			err := SerializeTimeFunc("key")(tt.accessor)(writer)
			assert.NoError(t, err)
			writer.AssertExpectations(t)
		})
	}
}

func TestSerializeObjectValueFunc(t *testing.T) {
	val := mocking.NewMockParsable()
	tests := []struct {
		name     string
		accessor ModelAccessor[serialization.Parsable]
		mock     func(*mocking.MockSerializationWriter)
		wantErr  bool
	}{
		{
			name:     "Success",
			accessor: func() (serialization.Parsable, error) { return val, nil },
			mock: func(m *mocking.MockSerializationWriter) {
				m.On("WriteObjectValue", "key", val, []serialization.Parsable(nil)).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			tt.mock(writer)

			err := SerializeObjectValueFunc[serialization.Parsable]("key")(tt.accessor)(writer)
			assert.NoError(t, err)
			writer.AssertExpectations(t)
		})
	}
}

func TestSerializeByteArrayFunc(t *testing.T) {
	val := []byte{1, 2, 3}
	tests := []struct {
		name     string
		accessor ModelAccessor[[]byte]
		mock     func(*mocking.MockSerializationWriter)
		wantErr  bool
	}{
		{
			name:     "Success",
			accessor: func() ([]byte, error) { return val, nil },
			mock: func(m *mocking.MockSerializationWriter) {
				m.On("WriteByteArrayValue", "key", val).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			tt.mock(writer)

			err := SerializeByteArrayFunc("key")(tt.accessor)(writer)
			assert.NoError(t, err)
			writer.AssertExpectations(t)
		})
	}
}

func TestSerializeCollectionOfObjectValuesFunc(t *testing.T) {
	val := []serialization.Parsable{mocking.NewMockParsable()}
	tests := []struct {
		name     string
		accessor ModelAccessor[[]serialization.Parsable]
		mock     func(*mocking.MockSerializationWriter)
		wantErr  bool
	}{
		{
			name:     "Success",
			accessor: func() ([]serialization.Parsable, error) { return val, nil },
			mock: func(m *mocking.MockSerializationWriter) {
				m.On("WriteCollectionOfObjectValues", "key", mock.Anything).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			tt.mock(writer)

			err := SerializeCollectionOfObjectValuesFunc[serialization.Parsable]("key")(tt.accessor)(writer)
			assert.NoError(t, err)
			writer.AssertExpectations(t)
		})
	}
}

func TestSerializeCollectionOfStringValuesFunc(t *testing.T) {
	val := []string{"a", "b"}
	tests := []struct {
		name     string
		accessor ModelAccessor[[]string]
		mock     func(*mocking.MockSerializationWriter)
		wantErr  bool
	}{
		{
			name:     "Success",
			accessor: func() ([]string, error) { return val, nil },
			mock: func(m *mocking.MockSerializationWriter) {
				m.On("WriteCollectionOfStringValues", "key", val).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			tt.mock(writer)

			err := SerializeCollectionOfStringValuesFunc("key")(tt.accessor)(writer)
			assert.NoError(t, err)
			writer.AssertExpectations(t)
		})
	}
}

type mockEnum int

const mockEnumVal mockEnum = 1

func (e *mockEnum) String() string { return "mock" }

func TestSerializeEnumFunc(t *testing.T) {
	val := mockEnumVal
	tests := []struct {
		name     string
		accessor ModelAccessor[*mockEnum]
		mock     func(*mocking.MockSerializationWriter)
		wantErr  bool
	}{
		{
			name:     "Success",
			accessor: func() (*mockEnum, error) { return &val, nil },
			mock: func(m *mocking.MockSerializationWriter) {
				s := "mock"
				m.On("WriteStringValue", "key", &s).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			tt.mock(writer)

			err := SerializeEnumFunc[mockEnum]("key")(tt.accessor)(writer)
			assert.NoError(t, err)
			writer.AssertExpectations(t)
		})
	}
}

func TestSerializeStringToBoolFunc(t *testing.T) {
	val := true
	tests := []struct {
		name     string
		accessor ModelAccessor[*bool]
		mock     func(*mocking.MockSerializationWriter)
		wantErr  bool
	}{
		{
			name:     "Success",
			accessor: func() (*bool, error) { return &val, nil },
			mock: func(m *mocking.MockSerializationWriter) {
				s := "true"
				m.On("WriteStringValue", "key", &s).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			tt.mock(writer)

			err := SerializeStringToBoolFunc("key")(tt.accessor)(writer)
			assert.NoError(t, err)
			writer.AssertExpectations(t)
		})
	}
}

func TestSerializeStringToFloat64Func(t *testing.T) {
	val := 1.23
	tests := []struct {
		name     string
		accessor ModelAccessor[*float64]
		mock     func(*mocking.MockSerializationWriter)
		wantErr  bool
	}{
		{
			name:     "Success",
			accessor: func() (*float64, error) { return &val, nil },
			mock: func(m *mocking.MockSerializationWriter) {
				s := "1.23"
				m.On("WriteStringValue", "key", &s).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			tt.mock(writer)

			err := SerializeStringToFloat64Func("key")(tt.accessor)(writer)
			assert.NoError(t, err)
			writer.AssertExpectations(t)
		})
	}
}

func TestSerializeStringToInt64Func(t *testing.T) {
	val := int64(123)
	tests := []struct {
		name     string
		accessor ModelAccessor[*int64]
		mock     func(*mocking.MockSerializationWriter)
		wantErr  bool
	}{
		{
			name:     "Success",
			accessor: func() (*int64, error) { return &val, nil },
			mock: func(m *mocking.MockSerializationWriter) {
				s := "123"
				m.On("WriteStringValue", "key", &s).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			tt.mock(writer)

			err := SerializeStringToInt64Func("key")(tt.accessor)(writer)
			assert.NoError(t, err)
			writer.AssertExpectations(t)
		})
	}
}

func TestSerializeStringToTimeFunc(t *testing.T) {
	val := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	layout := time.RFC3339
	tests := []struct {
		name     string
		accessor ModelAccessor[*time.Time]
		mock     func(*mocking.MockSerializationWriter)
		wantErr  bool
	}{
		{
			name:     "Success",
			accessor: func() (*time.Time, error) { return &val, nil },
			mock: func(m *mocking.MockSerializationWriter) {
				str := val.Format(layout)
				m.On("WriteStringValue", "key", &str).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			tt.mock(writer)

			err := SerializeStringToTimeFunc("key", layout)(tt.accessor)(writer)
			assert.NoError(t, err)
			writer.AssertExpectations(t)
		})
	}
}

func TestSerializeISODurationFunc(t *testing.T) {
	val := serialization.NewDuration(0, 0, 0, 0, 0, 1, 0)
	tests := []struct {
		name     string
		accessor ModelAccessor[*serialization.ISODuration]
		mock     func(*mocking.MockSerializationWriter)
		wantErr  bool
	}{
		{
			name:     "Success",
			accessor: func() (*serialization.ISODuration, error) { return val, nil },
			mock: func(m *mocking.MockSerializationWriter) {
				m.On("WriteISODurationValue", "key", val).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			tt.mock(writer)

			err := SerializeISODurationFunc("key")(tt.accessor)(writer)
			assert.NoError(t, err)
			writer.AssertExpectations(t)
		})
	}
}

func TestSerializeAnyFunc(t *testing.T) {
	val := "any"
	tests := []struct {
		name     string
		accessor ModelAccessor[any]
		mock     func(*mocking.MockSerializationWriter)
		wantErr  bool
	}{
		{
			name:     "Success",
			accessor: func() (any, error) { return val, nil },
			mock: func(m *mocking.MockSerializationWriter) {
				m.On("WriteAnyValue", "key", mock.Anything).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			tt.mock(writer)

			err := SerializeAnyFunc("key")(tt.accessor)(writer)
			assert.NoError(t, err)
			writer.AssertExpectations(t)
		})
	}
}

func TestSerializeStringToSliceFunc(t *testing.T) {
	val := []string{"a", "b"}
	tests := []struct {
		name     string
		accessor ModelAccessor[[]string]
		mock     func(*mocking.MockSerializationWriter)
		wantErr  bool
	}{
		{
			name:     "Success",
			accessor: func() ([]string, error) { return val, nil },
			mock: func(m *mocking.MockSerializationWriter) {
				s := "a,b"
				m.On("WriteStringValue", "key", &s).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			tt.mock(writer)

			err := SerializeStringToSliceFunc("key", ",")(tt.accessor)(writer)
			assert.NoError(t, err)
			writer.AssertExpectations(t)
		})
	}
}
