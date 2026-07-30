package caseapi

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// errWrite is a stand-in for an error a serialization.SerializationWriter can
// return from a Write* call.
var errWrite = errors.New("write error")

func TestReferenceModel_GettersSetters(t *testing.T) {
	model := NewReference()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"Link", func(v any) error { return model.setLink(v.(*string)) }, func() (any, error) { return model.GetLink() }, internal.ToPointer("https://example.com")},
		{"Value", func(v any) error { return model.setValue(v.(*string)) }, func() (any, error) { return model.GetValue() }, internal.ToPointer("test-value")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			require.NoError(t, err)
			got, err := tt.getter()
			require.NoError(t, err)
			assert.Equal(t, tt.value, got)
		})
	}
}

func TestCaseResultModel_GettersSetters(t *testing.T) {
	model := NewCaseResult()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"SysID", func(v any) error { return model.setSysID(v.(*string)) }, func() (any, error) { return model.GetSysID() }, internal.ToPointer("sys-id")},
		{"Number", func(v any) error { return model.setNumber(v.(*string)) }, func() (any, error) { return model.GetNumber() }, internal.ToPointer("CASE001")},
		{"ShortDescription", func(v any) error { return model.setShortDescription(v.(*string)) }, func() (any, error) { return model.GetShortDescription() }, internal.ToPointer("Short desc")},
		{"Description", func(v any) error { return model.setDescription(v.(*string)) }, func() (any, error) { return model.GetDescription() }, internal.ToPointer("Full desc")},
		{"State", func(v any) error { return model.setState(v.(*string)) }, func() (any, error) { return model.GetState() }, internal.ToPointer("10")},
		{"Priority", func(v any) error { return model.setPriority(v.(*string)) }, func() (any, error) { return model.GetPriority() }, internal.ToPointer("1")},
		{"Category", func(v any) error { return model.setCategory(v.(*string)) }, func() (any, error) { return model.GetCategory() }, internal.ToPointer("inquiry")},
		{"AssignmentGroup", func(v any) error { return model.setAssignmentGroup(v.(Reference)) }, func() (any, error) { return model.GetAssignmentGroup() }, NewReference()},
		{"AssignedTo", func(v any) error { return model.setAssignedTo(v.(Reference)) }, func() (any, error) { return model.GetAssignedTo() }, NewReference()},
		{"Contact", func(v any) error { return model.setContact(v.(Reference)) }, func() (any, error) { return model.GetContact() }, NewReference()},
		{"Account", func(v any) error { return model.setAccount(v.(Reference)) }, func() (any, error) { return model.GetAccount() }, NewReference()},
		{"SysCreatedOn", func(v any) error { return model.setSysCreatedOn(v.(*string)) }, func() (any, error) { return model.GetSysCreatedOn() }, internal.ToPointer("2023-01-01 12:00:00")},
		{"SysUpdatedOn", func(v any) error { return model.setSysUpdatedOn(v.(*string)) }, func() (any, error) { return model.GetSysUpdatedOn() }, internal.ToPointer("2023-01-01 13:00:00")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			require.NoError(t, err)
			got, err := tt.getter()
			require.NoError(t, err)
			assert.Equal(t, tt.value, got)
		})
	}
}

func TestActivitiesResultModel_GettersSetters(t *testing.T) {
	model := NewActivitiesResult()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"SysID", func(v any) error { return model.setSysID(v.(*string)) }, func() (any, error) { return model.GetSysID() }, internal.ToPointer("sys-id")},
		{"Type", func(v any) error { return model.setType(v.(*string)) }, func() (any, error) { return model.GetType() }, internal.ToPointer("work_notes")},
		{"Value", func(v any) error { return model.setValue(v.(*string)) }, func() (any, error) { return model.GetValue() }, internal.ToPointer("test value")},
		{"User", func(v any) error { return model.setUser(v.(*string)) }, func() (any, error) { return model.GetUser() }, internal.ToPointer("admin")},
		{"SysCreatedOn", func(v any) error { return model.setSysCreatedOn(v.(*string)) }, func() (any, error) { return model.GetSysCreatedOn() }, internal.ToPointer("2023-01-01 12:00:00")},
		{"FieldName", func(v any) error { return model.setFieldName(v.(*string)) }, func() (any, error) { return model.GetFieldName() }, internal.ToPointer("work_notes")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			require.NoError(t, err)
			got, err := tt.getter()
			require.NoError(t, err)
			assert.Equal(t, tt.value, got)
		})
	}
}

func TestFieldValuesResultModel_GettersSetters(t *testing.T) {
	model := NewFieldValuesResult()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"Label", func(v any) error { return model.setLabel(v.(*string)) }, func() (any, error) { return model.GetLabel() }, internal.ToPointer("label")},
		{"Value", func(v any) error { return model.setValue(v.(*string)) }, func() (any, error) { return model.GetValue() }, internal.ToPointer("value")},
		{"Sequence", func(v any) error { return model.setSequence(v.(*int32)) }, func() (any, error) { return model.GetSequence() }, internal.ToPointer(int32(1))},
		{"DependentValue", func(v any) error { return model.setDependentValue(v.(*string)) }, func() (any, error) { return model.GetDependentValue() }, internal.ToPointer("dep")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			require.NoError(t, err)
			got, err := tt.getter()
			require.NoError(t, err)
			assert.Equal(t, tt.value, got)
		})
	}
}

// ---------------------------------------------------------------------------
// Discriminator factories (models + responses)
// ---------------------------------------------------------------------------

func TestCreateFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name    string
		factory func() (any, error)
	}{
		{"Reference", func() (any, error) { return CreateReferenceFromDiscriminatorValue(nil) }},
		{"CaseResult", func() (any, error) { return CreateCaseResultFromDiscriminatorValue(nil) }},
		{"ActivitiesResult", func() (any, error) { return CreateActivitiesResultFromDiscriminatorValue(nil) }},
		{"FieldValuesResult", func() (any, error) { return CreateFieldValuesResultFromDiscriminatorValue(nil) }},
		{"CaseCollectionResponse", func() (any, error) { return CreateCaseCollectionResponseFromDiscriminatorValue(nil) }},
		{"CaseItemResponse", func() (any, error) { return CreateCaseItemResponseFromDiscriminatorValue(nil) }},
		{"ActivitiesResponse", func() (any, error) { return CreateActivitiesResponseFromDiscriminatorValue(nil) }},
		{"FieldValuesResponse", func() (any, error) { return CreateFieldValuesResponseFromDiscriminatorValue(nil) }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsable, err := tt.factory()
			require.NoError(t, err)
			assert.NotNil(t, parsable)
		})
	}
}

// ---------------------------------------------------------------------------
// Serialize / GetFieldDeserializers
// ---------------------------------------------------------------------------

func TestReferenceModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *ReferenceModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name: "happy path - writes all fields",
			model: func() *ReferenceModel {
				m := NewReference()
				_ = m.setLink(internal.ToPointer("link"))
				_ = m.setValue(internal.ToPointer("value"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", linkKey, mock.Anything).Return(nil)
				w.On("WriteStringValue", valueKey, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *ReferenceModel {
				m := NewReference()
				_ = m.setLink(internal.ToPointer("link"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", linkKey, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			if tt.setupMock != nil {
				tt.setupMock(writer)
			}

			err := tt.model.Serialize(writer)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestReferenceModel_GetFieldDeserializers(t *testing.T) {
	model := NewReference()
	deserializers := model.GetFieldDeserializers()
	assert.Len(t, deserializers, 2)
	assert.NotNil(t, deserializers[linkKey])
	assert.NotNil(t, deserializers[valueKey])
}

func TestCaseResultModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *CaseResultModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewCaseResult(),
		},
		{
			name: "happy path - writes all fields including nested references",
			model: func() *CaseResultModel {
				m := NewCaseResult()
				_ = m.setSysID(internal.ToPointer("sys-id"))
				_ = m.setNumber(internal.ToPointer("CASE001"))
				_ = m.setShortDescription(internal.ToPointer("short"))
				_ = m.setDescription(internal.ToPointer("desc"))
				_ = m.setState(internal.ToPointer("10"))
				_ = m.setPriority(internal.ToPointer("1"))
				_ = m.setCategory(internal.ToPointer("inquiry"))
				_ = m.setAssignmentGroup(NewReference())
				_ = m.setAssignedTo(NewReference())
				_ = m.setContact(NewReference())
				_ = m.setAccount(NewReference())
				_ = m.setSysCreatedOn(internal.ToPointer("2023-01-01"))
				_ = m.setSysUpdatedOn(internal.ToPointer("2023-01-02"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteObjectValue", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *CaseResultModel {
				m := NewCaseResult()
				_ = m.setSysID(internal.ToPointer("sys-id"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", sysIDKey, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
		{
			name: "nested object write error propagates",
			model: func() *CaseResultModel {
				m := NewCaseResult()
				_ = m.setAssignmentGroup(NewReference())
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteObjectValue", assignmentGroupKey, mock.Anything, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			if tt.setupMock != nil {
				tt.setupMock(writer)
			}

			err := tt.model.Serialize(writer)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestCaseResultModel_GetFieldDeserializers(t *testing.T) {
	model := NewCaseResult()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{
		sysIDKey, numberKey, shortDescriptionKey, descriptionKey, stateKey,
		priorityKey, categoryKey, assignmentGroupKey, assignedToKey, contactKey,
		accountKey, sysCreatedOnKey, sysUpdatedOnKey,
	} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 13)
}

func TestActivitiesResultModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *ActivitiesResultModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewActivitiesResult(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *ActivitiesResultModel {
				m := NewActivitiesResult()
				_ = m.setSysID(internal.ToPointer("sys-id"))
				_ = m.setType(internal.ToPointer("work_notes"))
				_ = m.setValue(internal.ToPointer("value"))
				_ = m.setUser(internal.ToPointer("admin"))
				_ = m.setSysCreatedOn(internal.ToPointer("2023-01-01"))
				_ = m.setFieldName(internal.ToPointer("work_notes"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *ActivitiesResultModel {
				m := NewActivitiesResult()
				_ = m.setSysID(internal.ToPointer("sys-id"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", sysIDKey, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			if tt.setupMock != nil {
				tt.setupMock(writer)
			}

			err := tt.model.Serialize(writer)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestActivitiesResultModel_GetFieldDeserializers(t *testing.T) {
	model := NewActivitiesResult()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{sysIDKey, typeKey, valueKey, userKey, sysCreatedOnKey, fieldNameKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 6)
}

func TestFieldValuesResultModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *FieldValuesResultModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewFieldValuesResult(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *FieldValuesResultModel {
				m := NewFieldValuesResult()
				_ = m.setLabel(internal.ToPointer("label"))
				_ = m.setValue(internal.ToPointer("value"))
				_ = m.setSequence(internal.ToPointer(int32(1)))
				_ = m.setDependentValue(internal.ToPointer("dep"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteInt32Value", sequenceKey, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *FieldValuesResultModel {
				m := NewFieldValuesResult()
				_ = m.setLabel(internal.ToPointer("label"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", labelKey, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
		{
			name: "int32 write error propagates",
			model: func() *FieldValuesResultModel {
				m := NewFieldValuesResult()
				_ = m.setSequence(internal.ToPointer(int32(1)))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteInt32Value", sequenceKey, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			if tt.setupMock != nil {
				tt.setupMock(writer)
			}

			err := tt.model.Serialize(writer)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestFieldValuesResultModel_GetFieldDeserializers(t *testing.T) {
	model := NewFieldValuesResult()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{labelKey, valueKey, sequenceKey, dependentValueKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 4)
}
