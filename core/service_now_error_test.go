package core

import (
	"errors"
	"testing"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewServiceNowError(t *testing.T) {
	err := NewServiceNowError()
	if err == nil {
		t.Fatal("NewServiceNowError returned nil")
	}
}

func TestCreateServiceNowErrorFromDiscriminatorValue(t *testing.T) {
	res, err := CreateServiceNowErrorFromDiscriminatorValue(nil)
	if err != nil {
		t.Errorf("unexpected err %v", err)
	}
	if res == nil {
		t.Error("returned nil")
	}
}

func TestServiceNowError_Serialize(t *testing.T) {
	err := NewServiceNowError().Serialize(nil)
	require.NoError(t, err)

	var nilE *ServiceNowError
	err = nilE.Serialize(nil)
	assert.NoError(t, err)
}

func TestServiceNowError_GetFieldDeserializers(t *testing.T) {
	deser := NewServiceNowError().GetFieldDeserializers()
	if deser[errorKey] == nil {
		t.Error("missing deserializer")
	}
}

func TestServiceNowError_GetError(t *testing.T) {
	me := NewMainError()
	e := NewServiceNowError()
	_ = e.setError(me)
	var nilE *ServiceNowError

	tests := []struct {
		name     string
		model    *ServiceNowError
		expected MainErrorable
		err      bool
	}{
		{"Ok", e, me, false},
		{"NilE", nilE, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.model.GetError()
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
			if res != tt.expected {
				t.Errorf("got %v, expected %v", res, tt.expected)
			}
		})
	}
}

func TestServiceNowError_setError(t *testing.T) {
	me := NewMainError()
	e := NewServiceNowError()
	var nilE *ServiceNowError

	tests := []struct {
		name  string
		model *ServiceNowError
		val   MainErrorable
		err   bool
	}{
		{"Ok", e, me, false},
		{"NilE", nilE, me, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.model.setError(tt.val)
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
		})
	}
}

func TestCreateSpecificErrorFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name    string
		create  func(serialization.ParseNode) (serialization.Parsable, error)
		wantErr string
	}{
		{
			name:   "BadRequestError",
			create: CreateBadRequestErrorFromDiscriminatorValue,
		},
		{
			name:   "UnauthorizedError",
			create: CreateUnauthorizedErrorFromDiscriminatorValue,
		},
		{
			name:   "ForbiddenError",
			create: CreateForbiddenErrorFromDiscriminatorValue,
		},
		{
			name:   "NotFoundError",
			create: CreateNotFoundErrorFromDiscriminatorValue,
		},
		{
			name:   "TooManyRequestsError",
			create: CreateTooManyRequestsErrorFromDiscriminatorValue,
		},
		{
			name:   "ServerError",
			create: CreateServerErrorFromDiscriminatorValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.create(nil)
			require.NoError(t, err)
			assert.NotNil(t, res)
		})
	}
}

func TestDefaultErrorMapping(t *testing.T) {
	mapping := DefaultErrorMapping()

	tests := []struct {
		key      string
		wantType any
	}{
		{"400", &BadRequestError{}},
		{"401", &UnauthorizedError{}},
		{"403", &ForbiddenError{}},
		{"404", &NotFoundError{}},
		{"429", &TooManyRequestsError{}},
		{"5XX", &ServerError{}},
		{"XXX", &ServiceNowError{}},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			factory, ok := mapping[tt.key]
			require.True(t, ok, "expected mapping for key %q", tt.key)

			res, err := factory(nil)
			require.NoError(t, err)
			assert.IsType(t, tt.wantType, res)
		})
	}
}

func TestServiceNowError_Error(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *ServiceNowError
		expected string
	}{
		{
			name: "message present",
			setup: func() *ServiceNowError {
				e := NewServiceNowError()
				me := NewMainError()
				_ = me.SetMessage(func() *string { s := "bad request"; return &s }())
				_ = e.setError(me)
				return e
			},
			expected: "bad request",
		},
		{
			name: "message nil, detail present",
			setup: func() *ServiceNowError {
				e := NewServiceNowError()
				me := NewMainError()
				_ = me.SetDetail(func() *string { s := "detailed reason"; return &s }())
				_ = e.setError(me)
				return e
			},
			expected: "detailed reason",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tt.setup()
			assert.Equal(t, tt.expected, e.Error())
		})
	}
}

// TODO: improve test table design
func TestServiceNowError_ErrorBranches(t *testing.T) {
	eWrongType := NewServiceNowError()
	require.NoError(t, eWrongType.GetBackingStore().Set(errorKey, 123))
	val, err := eWrongType.GetError()
	assert.Nil(t, val)
	assert.Equal(t, errors.New("cannot convert '123' to type core.MainErrorable"), err)

	eNilBS := &ServiceNowError{BackedModel: &mockNilBSModel{}}
	val, err = eNilBS.GetError()
	assert.Nil(t, val)
	assert.Equal(t, errors.New("store is nil"), err)
}
