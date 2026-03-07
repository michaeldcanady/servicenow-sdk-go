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

func TestSerializeStringFunc(t *testing.T) {
	val := "test"
	tests := []struct {
		name     string
		accessor ModelAccessor[*string]
		mock     func(*mocking.MockSerializationWriter)
		wantErr  bool
	}{
		{
			name: "Success",
			accessor: func() (*string, error) { return &val, nil },
			mock: func(m *mocking.MockSerializationWriter) {
				m.On("WriteStringValue", "key", &val).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Nil value",
			accessor: func() (*string, error) { return nil, nil },
			mock: func(m *mocking.MockSerializationWriter) {},
			wantErr: false,
		},
		{
			name: "Error",
			accessor: func() (*string, error) { return nil, errors.New("err") },
			mock: func(m *mocking.MockSerializationWriter) {},
			wantErr: true,
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
			name: "Success",
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
			name: "Success",
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

func TestSerializeCollectionOfObjectValuesFunc(t *testing.T) {
	val := []serialization.Parsable{mocking.NewMockParsable()}
	tests := []struct {
		name     string
		accessor ModelAccessor[[]serialization.Parsable]
		mock     func(*mocking.MockSerializationWriter)
		wantErr  bool
	}{
		{
			name: "Success",
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

func TestSerializeStringToTimeFunc(t *testing.T) {
	val := time.Now()
	layout := time.RFC3339
	tests := []struct {
		name     string
		accessor ModelAccessor[*time.Time]
		mock     func(*mocking.MockSerializationWriter)
		wantErr  bool
	}{
		{
			name: "Success",
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
			
			err := SerializeStringToTimeFunc("key", layout)(func() (*time.Time, error) { return &val, nil })(writer)
			assert.NoError(t, err)
			writer.AssertExpectations(t)
		})
	}
}
