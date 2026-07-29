package cdmchangesetapi

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

func TestChangesetResultModel_GettersSetters(t *testing.T) {
	model := NewChangesetResult()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"AutoValidate", func(v interface{}) error { return model.setAutoValidate(v.(*bool)) }, func() (interface{}, error) { return model.GetAutoValidate() }, internal.ToPointer(true)},
		{"CdmApplication", func(v interface{}) error { return model.setCdmApplication(v.(*Reference)) }, func() (interface{}, error) { return model.GetCdmApplication() }, NewReference()},
		{"CommittedAt", func(v interface{}) error { return model.setCommittedAt(v.(*string)) }, func() (interface{}, error) { return model.GetCommittedAt() }, internal.ToPointer("2023-01-01 12:00:00")},
		{"CommittedBy", func(v interface{}) error { return model.setCommittedBy(v.(*Reference)) }, func() (interface{}, error) { return model.GetCommittedBy() }, NewReference()},
		{"Description", func(v interface{}) error { return model.setDescription(v.(*string)) }, func() (interface{}, error) { return model.GetDescription() }, internal.ToPointer("changeset desc")},
		{"LastConflictDetectionTime", func(v interface{}) error { return model.setLastConflictDetectionTime(v.(*int64)) }, func() (interface{}, error) { return model.GetLastConflictDetectionTime() }, internal.ToPointer(int64(123456789))},
		{"Number", func(v interface{}) error { return model.setNumber(v.(*string)) }, func() (interface{}, error) { return model.GetNumber() }, internal.ToPointer("CHG001")},
		{"PublishOption", func(v interface{}) error { return model.setPublishOption(v.(*string)) }, func() (interface{}, error) { return model.GetPublishOption() }, internal.ToPointer("all")},
		{"State", func(v interface{}) error { return model.setState(v.(*string)) }, func() (interface{}, error) { return model.GetState() }, internal.ToPointer("committed")},
		{"SysID", func(v interface{}) error { return model.setSysID(v.(*string)) }, func() (interface{}, error) { return model.GetSysID() }, internal.ToPointer("sys-id")},
		{"Title", func(v interface{}) error { return model.setTitle(v.(*string)) }, func() (interface{}, error) { return model.GetTitle() }, internal.ToPointer("Changeset Title")},
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

func TestChangesetActivityResultModel_GettersSetters(t *testing.T) {
	model := NewChangesetActivityResult()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"ChangesetId", func(v interface{}) error { return model.setChangesetID(v.(*Reference)) }, func() (interface{}, error) { return model.GetChangesetID() }, NewReference()},
		{"Conflict", func(v interface{}) error { return model.setConflict(v.(*bool)) }, func() (interface{}, error) { return model.GetConflict() }, internal.ToPointer(false)},
		{"NamePath", func(v interface{}) error { return model.setNamePath(v.(*string)) }, func() (interface{}, error) { return model.GetNamePath() }, internal.ToPointer("/path")},
		{"NewName", func(v interface{}) error { return model.setNewName(v.(*string)) }, func() (interface{}, error) { return model.GetNewName() }, internal.ToPointer("new")},
		{"OldName", func(v interface{}) error { return model.setOldName(v.(*string)) }, func() (interface{}, error) { return model.GetOldName() }, internal.ToPointer("old")},
		{"NewValue", func(v interface{}) error { return model.setNewValue(v.(*string)) }, func() (interface{}, error) { return model.GetNewValue() }, internal.ToPointer("new-val")},
		{"OldValue", func(v interface{}) error { return model.setOldValue(v.(*string)) }, func() (interface{}, error) { return model.GetOldValue() }, internal.ToPointer("old-val")},
		{"Type", func(v interface{}) error { return model.setType(v.(*string)) }, func() (interface{}, error) { return model.GetType() }, internal.ToPointer("type")},
		{"Secure", func(v interface{}) error { return model.setSecure(v.(*bool)) }, func() (interface{}, error) { return model.GetSecure() }, internal.ToPointer(true)},
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

func TestCreateChangesetResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateChangesetResultFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCommitStatusResultModel_GettersSetters(t *testing.T) {
	model := NewCommitStatusResult()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"State", func(v interface{}) error { return model.setState(v.(*string)) }, func() (interface{}, error) { return model.GetState() }, internal.ToPointer("completed")},
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

func TestImpactedSharedComponentResultModel_GettersSetters(t *testing.T) {
	model := NewImpactedSharedComponentResult()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"CdmSharedLibrary", func(v interface{}) error { return model.setCdmSharedLibrary(v.(*string)) }, func() (interface{}, error) { return model.GetCdmSharedLibrary() }, internal.ToPointer("lib-id")},
		{"Description", func(v interface{}) error { return model.setDescription(v.(*string)) }, func() (interface{}, error) { return model.GetDescription() }, internal.ToPointer("desc")},
		{"Name", func(v interface{}) error { return model.setName(v.(*string)) }, func() (interface{}, error) { return model.GetName() }, internal.ToPointer("name")},
		{"Node", func(v interface{}) error { return model.setNode(v.(*string)) }, func() (interface{}, error) { return model.GetNode() }, internal.ToPointer("node-id")},
		{"NodeMain", func(v interface{}) error { return model.setNodeMain(v.(*string)) }, func() (interface{}, error) { return model.GetNodeMain() }, internal.ToPointer("main-id")},
		{"State", func(v interface{}) error { return model.setState(v.(*string)) }, func() (interface{}, error) { return model.GetState() }, internal.ToPointer("active")},
		{"SysCreatedBy", func(v interface{}) error { return model.setSysCreatedBy(v.(*string)) }, func() (interface{}, error) { return model.GetSysCreatedBy() }, internal.ToPointer("admin")},
		{"SysCreatedOn", func(v interface{}) error { return model.setSysCreatedOn(v.(*string)) }, func() (interface{}, error) { return model.GetSysCreatedOn() }, internal.ToPointer("2023-01-01")},
		{"SysID", func(v interface{}) error { return model.setSysID(v.(*string)) }, func() (interface{}, error) { return model.GetSysID() }, internal.ToPointer("sys-id")},
		{"SysUpdatedBy", func(v interface{}) error { return model.setSysUpdatedBy(v.(*string)) }, func() (interface{}, error) { return model.GetSysUpdatedBy() }, internal.ToPointer("admin")},
		{"SysUpdatedOn", func(v interface{}) error { return model.setSysUpdatedOn(v.(*string)) }, func() (interface{}, error) { return model.GetSysUpdatedOn() }, internal.ToPointer("2023-01-01")},
		{"VersionCounter", func(v interface{}) error { return model.setVersionCounter(v.(*int32)) }, func() (interface{}, error) { return model.GetVersionCounter() }, internal.ToPointer(int32(1))},
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

func TestCreateCommitStatusResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateCommitStatusResultFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestImpactedDeployableResultModel_GettersSetters(t *testing.T) {
	model := NewImpactedDeployableResult()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"CdiCount", func(v interface{}) error { return model.setCdiCount(v.(*int32)) }, func() (interface{}, error) { return model.GetCdiCount() }, internal.ToPointer(int32(1))},
		{"CdiUsage", func(v interface{}) error { return model.setCdiUsage(v.(*string)) }, func() (interface{}, error) { return model.GetCdiUsage() }, internal.ToPointer("usage")},
		{"CdmApp", func(v interface{}) error { return model.setCdmApp(v.(*Reference)) }, func() (interface{}, error) { return model.GetCdmApp() }, NewReference()},
		{"CdmCi", func(v interface{}) error { return model.setCdmCi(v.(*Reference)) }, func() (interface{}, error) { return model.GetCdmCi() }, NewReference()},
		{"Description", func(v interface{}) error { return model.setDescription(v.(*string)) }, func() (interface{}, error) { return model.GetDescription() }, internal.ToPointer("desc")},
		{"EnvironmentType", func(v interface{}) error { return model.setEnvironmentType(v.(*string)) }, func() (interface{}, error) { return model.GetEnvironmentType() }, internal.ToPointer("prod")},
		{"Name", func(v interface{}) error { return model.setName(v.(*string)) }, func() (interface{}, error) { return model.GetName() }, internal.ToPointer("name")},
		{"Node", func(v interface{}) error { return model.setNode(v.(*Reference)) }, func() (interface{}, error) { return model.GetNode() }, NewReference()},
		{"SnapshotVersionCounter", func(v interface{}) error { return model.setSnapshotVersionCounter(v.(*int32)) }, func() (interface{}, error) { return model.GetSnapshotVersionCounter() }, internal.ToPointer(int32(1))},
		{"State", func(v interface{}) error { return model.setState(v.(*string)) }, func() (interface{}, error) { return model.GetState() }, internal.ToPointer("active")},
		{"SysID", func(v interface{}) error { return model.setSysID(v.(*string)) }, func() (interface{}, error) { return model.GetSysID() }, internal.ToPointer("sys-id")},
		{"SysCreatedBy", func(v interface{}) error { return model.setSysCreatedBy(v.(*string)) }, func() (interface{}, error) { return model.GetSysCreatedBy() }, internal.ToPointer("admin")},
		{"SysCreatedOn", func(v interface{}) error { return model.setSysCreatedOn(v.(*string)) }, func() (interface{}, error) { return model.GetSysCreatedOn() }, internal.ToPointer("2023-01-01")},
		{"SysUpdatedBy", func(v interface{}) error { return model.setSysUpdatedBy(v.(*string)) }, func() (interface{}, error) { return model.GetSysUpdatedBy() }, internal.ToPointer("admin")},
		{"SysUpdatedOn", func(v interface{}) error { return model.setSysUpdatedOn(v.(*string)) }, func() (interface{}, error) { return model.GetSysUpdatedOn() }, internal.ToPointer("2023-01-01")},
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

func TestImpactedDeployableBySysIdResultModel_GettersSetters(t *testing.T) {
	model := NewImpactedDeployableBySysIDResult()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"ChangesetId", func(v interface{}) error { return model.setChangesetID(v.(*string)) }, func() (interface{}, error) { return model.GetChangesetID() }, internal.ToPointer("chg-id")},
		{"Conflict", func(v interface{}) error { return model.setConflict(v.(*bool)) }, func() (interface{}, error) { return model.GetConflict() }, internal.ToPointer(true)},
		{"ConflictType", func(v interface{}) error { return model.setConflictType(v.(*string)) }, func() (interface{}, error) { return model.GetConflictType() }, internal.ToPointer("type")},
		{"Description", func(v interface{}) error { return model.setDescription(v.(*string)) }, func() (interface{}, error) { return model.GetDescription() }, internal.ToPointer("desc")},
		{"EffectiveFrom", func(v interface{}) error { return model.setEffectiveFrom(v.(*string)) }, func() (interface{}, error) { return model.GetEffectiveFrom() }, internal.ToPointer("2023-01-01")},
		{"EffectiveTo", func(v interface{}) error { return model.setEffectiveTo(v.(*string)) }, func() (interface{}, error) { return model.GetEffectiveTo() }, internal.ToPointer("2023-12-31")},
		{"Level", func(v interface{}) error { return model.setLevel(v.(*int32)) }, func() (interface{}, error) { return model.GetLevel() }, internal.ToPointer(int32(1))},
		{"LinkedTo", func(v interface{}) error { return model.setLinkedTo(v.(*string)) }, func() (interface{}, error) { return model.GetLinkedTo() }, internal.ToPointer("link")},
		{"MainId", func(v interface{}) error { return model.setMainID(v.(*string)) }, func() (interface{}, error) { return model.GetMainID() }, internal.ToPointer("main-id")},
		{"MainIdEncoded", func(v interface{}) error { return model.setMainIDEncoded(v.(*string)) }, func() (interface{}, error) { return model.GetMainIDEncoded() }, internal.ToPointer("encoded")},
		{"Name", func(v interface{}) error { return model.setName(v.(*string)) }, func() (interface{}, error) { return model.GetName() }, internal.ToPointer("name")},
		{"NodeClassifier", func(v interface{}) error { return model.setNodeClassifier(v.(*string)) }, func() (interface{}, error) { return model.GetNodeClassifier() }, internal.ToPointer("classifier")},
		{"Status", func(v interface{}) error { return model.setStatus(v.(*string)) }, func() (interface{}, error) { return model.GetStatus() }, internal.ToPointer("status")},
		{"SysID", func(v interface{}) error { return model.setSysID(v.(*string)) }, func() (interface{}, error) { return model.GetSysID() }, internal.ToPointer("sys-id")},
		{"Type", func(v interface{}) error { return model.setType(v.(*string)) }, func() (interface{}, error) { return model.GetType() }, internal.ToPointer("type")},
		{"Value", func(v interface{}) error { return model.setValue(v.(*string)) }, func() (interface{}, error) { return model.GetValue() }, internal.ToPointer("value")},
		{"SecureValue", func(v interface{}) error { return model.setSecureValue(v.(*string)) }, func() (interface{}, error) { return model.GetSecureValue() }, internal.ToPointer("secure")},
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

func TestCreateImpactedDeployableResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateImpactedDeployableResultFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateImpactedDeployableBySysIdResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateImpactedDeployableBySysIDResultFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

// ---------------------------------------------------------------------------
// Reference
// ---------------------------------------------------------------------------

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

func TestReferenceModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *Reference
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewReference(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *Reference {
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
			model: func() *Reference {
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

func TestCreateReferenceFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateReferenceFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

// ---------------------------------------------------------------------------
// ChangesetResult - Serialize / GetFieldDeserializers
// ---------------------------------------------------------------------------

func TestChangesetResultModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *ChangesetResult
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewChangesetResult(),
		},
		{
			name: "happy path - writes all fields including nested references",
			model: func() *ChangesetResult {
				m := NewChangesetResult()
				_ = m.setAutoValidate(internal.ToPointer(true))
				_ = m.setCdmApplication(NewReference())
				_ = m.setCommittedAt(internal.ToPointer("2023-01-01"))
				_ = m.setCommittedBy(NewReference())
				_ = m.setDescription(internal.ToPointer("desc"))
				_ = m.setLastConflictDetectionTime(internal.ToPointer(int64(123)))
				_ = m.setNumber(internal.ToPointer("CHG001"))
				_ = m.setPublishOption(internal.ToPointer("all"))
				_ = m.setState(internal.ToPointer("committed"))
				_ = m.setSysID(internal.ToPointer("sys-id"))
				_ = m.setTitle(internal.ToPointer("title"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteBoolValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteInt64Value", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteObjectValue", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *ChangesetResult {
				m := NewChangesetResult()
				_ = m.setAutoValidate(internal.ToPointer(true))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteBoolValue", autoValidateKey, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
		{
			name: "nested object write error propagates",
			model: func() *ChangesetResult {
				m := NewChangesetResult()
				_ = m.setCdmApplication(NewReference())
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteObjectValue", cdmApplicationKey, mock.Anything, mock.Anything).Return(errWrite)
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

func TestChangesetResultModel_GetFieldDeserializers(t *testing.T) {
	model := NewChangesetResult()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{
		autoValidateKey, cdmApplicationKey, committedAtKey, committedByKey, descriptionKey,
		lastConflictDetectionTimeKey, numberKey, publishOptionKey, stateKey, sysIDKey, titleKey,
	} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 11)
}

// ---------------------------------------------------------------------------
// ChangesetActivityResult - Serialize / GetFieldDeserializers / discriminator
// ---------------------------------------------------------------------------

func TestChangesetActivityResultModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *ChangesetActivityResult
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewChangesetActivityResult(),
		},
		{
			name: "happy path - writes all fields including nested reference",
			model: func() *ChangesetActivityResult {
				m := NewChangesetActivityResult()
				_ = m.setChangesetID(NewReference())
				_ = m.setConflict(internal.ToPointer(false))
				_ = m.setNamePath(internal.ToPointer("/path"))
				_ = m.setNewName(internal.ToPointer("new"))
				_ = m.setOldName(internal.ToPointer("old"))
				_ = m.setNewValue(internal.ToPointer("new-val"))
				_ = m.setOldValue(internal.ToPointer("old-val"))
				_ = m.setType(internal.ToPointer("type"))
				_ = m.setSecure(internal.ToPointer(true))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteObjectValue", changesetIDKey, mock.Anything, mock.Anything).Return(nil)
				w.On("WriteBoolValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "nested object write error propagates",
			model: func() *ChangesetActivityResult {
				m := NewChangesetActivityResult()
				_ = m.setChangesetID(NewReference())
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteObjectValue", changesetIDKey, mock.Anything, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
		{
			name: "write error propagates",
			model: func() *ChangesetActivityResult {
				m := NewChangesetActivityResult()
				_ = m.setConflict(internal.ToPointer(true))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteBoolValue", conflictKey, mock.Anything).Return(errWrite)
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

func TestChangesetActivityResultModel_GetFieldDeserializers(t *testing.T) {
	model := NewChangesetActivityResult()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{
		changesetIDKey, conflictKey, namePathKey, newNameKey, oldNameKey,
		newValueKey, oldValueKey, typeKey, secureKey,
	} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 9)
}

func TestCreateChangesetActivityResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateChangesetActivityResultFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

// ---------------------------------------------------------------------------
// CommitStatusResult - Serialize / GetFieldDeserializers
// ---------------------------------------------------------------------------

func TestCommitStatusResultModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *CommitStatusResult
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewCommitStatusResult(),
		},
		{
			name: "happy path - writes state",
			model: func() *CommitStatusResult {
				m := NewCommitStatusResult()
				_ = m.setState(internal.ToPointer("completed"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", stateKey, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *CommitStatusResult {
				m := NewCommitStatusResult()
				_ = m.setState(internal.ToPointer("completed"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", stateKey, mock.Anything).Return(errWrite)
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

func TestCommitStatusResultModel_GetFieldDeserializers(t *testing.T) {
	model := NewCommitStatusResult()
	deserializers := model.GetFieldDeserializers()
	assert.Len(t, deserializers, 1)
	assert.NotNil(t, deserializers[stateKey])
}

// ---------------------------------------------------------------------------
// ImpactedDeployableResult - Serialize / GetFieldDeserializers
// ---------------------------------------------------------------------------

func TestImpactedDeployableResultModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *ImpactedDeployableResult
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewImpactedDeployableResult(),
		},
		{
			name: "happy path - writes all fields including nested references",
			model: func() *ImpactedDeployableResult {
				m := NewImpactedDeployableResult()
				_ = m.setCdiCount(internal.ToPointer(int32(1)))
				_ = m.setCdiUsage(internal.ToPointer("usage"))
				_ = m.setCdmApp(NewReference())
				_ = m.setCdmCi(NewReference())
				_ = m.setDescription(internal.ToPointer("desc"))
				_ = m.setEnvironmentType(internal.ToPointer("prod"))
				_ = m.setName(internal.ToPointer("name"))
				_ = m.setNode(NewReference())
				_ = m.setSnapshotVersionCounter(internal.ToPointer(int32(1)))
				_ = m.setState(internal.ToPointer("active"))
				_ = m.setSysID(internal.ToPointer("sys-id"))
				_ = m.setSysCreatedBy(internal.ToPointer("admin"))
				_ = m.setSysCreatedOn(internal.ToPointer("2023-01-01"))
				_ = m.setSysUpdatedBy(internal.ToPointer("admin"))
				_ = m.setSysUpdatedOn(internal.ToPointer("2023-01-01"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteInt32Value", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteObjectValue", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *ImpactedDeployableResult {
				m := NewImpactedDeployableResult()
				_ = m.setCdiCount(internal.ToPointer(int32(1)))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteInt32Value", cdiCountKey, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
		{
			name: "nested object write error propagates",
			model: func() *ImpactedDeployableResult {
				m := NewImpactedDeployableResult()
				_ = m.setCdmApp(NewReference())
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteObjectValue", cdmAppKey, mock.Anything, mock.Anything).Return(errWrite)
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

func TestImpactedDeployableResultModel_GetFieldDeserializers(t *testing.T) {
	model := NewImpactedDeployableResult()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{
		cdiCountKey, cdiUsageKey, cdmAppKey, cdmCiKey, descriptionKey, environmentTypeKey,
		nameKey, nodeKey, snapshotVersionCounterKey, stateKey, sysIDKey, sysCreatedByKey,
		sysCreatedOnKey, sysUpdatedByKey, sysUpdatedOnKey,
	} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 15)
}

// ---------------------------------------------------------------------------
// ImpactedDeployableBySysIDResult - Serialize / GetFieldDeserializers
// ---------------------------------------------------------------------------

func TestImpactedDeployableBySysIDResultModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *ImpactedDeployableBySysIDResult
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewImpactedDeployableBySysIDResult(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *ImpactedDeployableBySysIDResult {
				m := NewImpactedDeployableBySysIDResult()
				_ = m.setChangesetID(internal.ToPointer("chg-id"))
				_ = m.setConflict(internal.ToPointer(true))
				_ = m.setConflictType(internal.ToPointer("type"))
				_ = m.setDescription(internal.ToPointer("desc"))
				_ = m.setEffectiveFrom(internal.ToPointer("2023-01-01"))
				_ = m.setEffectiveTo(internal.ToPointer("2023-12-31"))
				_ = m.setLevel(internal.ToPointer(int32(1)))
				_ = m.setLinkedTo(internal.ToPointer("link"))
				_ = m.setMainID(internal.ToPointer("main-id"))
				_ = m.setMainIDEncoded(internal.ToPointer("encoded"))
				_ = m.setName(internal.ToPointer("name"))
				_ = m.setNodeClassifier(internal.ToPointer("classifier"))
				_ = m.setStatus(internal.ToPointer("status"))
				_ = m.setSysID(internal.ToPointer("sys-id"))
				_ = m.setType(internal.ToPointer("type"))
				_ = m.setValue(internal.ToPointer("value"))
				_ = m.setSecureValue(internal.ToPointer("secure"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteBoolValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteInt32Value", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *ImpactedDeployableBySysIDResult {
				m := NewImpactedDeployableBySysIDResult()
				_ = m.setChangesetID(internal.ToPointer("chg-id"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", changesetIDKey, mock.Anything).Return(errWrite)
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

func TestImpactedDeployableBySysIDResultModel_GetFieldDeserializers(t *testing.T) {
	model := NewImpactedDeployableBySysIDResult()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{
		changesetIDKey, conflictKey, conflictTypeKey, descriptionKey, effectiveFromKey,
		effectiveToKey, levelKey, linkedToKey, mainIDKey, mainIDEncodedKey, nameKey,
		nodeClassifierKey, statusKey, sysIDKey, typeKey, valueKey, secureValueKey,
	} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 17)
}

// ---------------------------------------------------------------------------
// ImpactedSharedComponentResult - Serialize / GetFieldDeserializers / discriminator
// ---------------------------------------------------------------------------

func TestImpactedSharedComponentResultModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *ImpactedSharedComponentResult
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewImpactedSharedComponentResult(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *ImpactedSharedComponentResult {
				m := NewImpactedSharedComponentResult()
				_ = m.setCdmSharedLibrary(internal.ToPointer("lib-id"))
				_ = m.setDescription(internal.ToPointer("desc"))
				_ = m.setName(internal.ToPointer("name"))
				_ = m.setNode(internal.ToPointer("node-id"))
				_ = m.setNodeMain(internal.ToPointer("main-id"))
				_ = m.setState(internal.ToPointer("active"))
				_ = m.setSysCreatedBy(internal.ToPointer("admin"))
				_ = m.setSysCreatedOn(internal.ToPointer("2023-01-01"))
				_ = m.setSysID(internal.ToPointer("sys-id"))
				_ = m.setSysUpdatedBy(internal.ToPointer("admin"))
				_ = m.setSysUpdatedOn(internal.ToPointer("2023-01-01"))
				_ = m.setVersionCounter(internal.ToPointer(int32(1)))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteInt32Value", versionCounterKey, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *ImpactedSharedComponentResult {
				m := NewImpactedSharedComponentResult()
				_ = m.setCdmSharedLibrary(internal.ToPointer("lib-id"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", cdmSharedLibraryKey, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
		{
			name: "int32 write error propagates",
			model: func() *ImpactedSharedComponentResult {
				m := NewImpactedSharedComponentResult()
				_ = m.setVersionCounter(internal.ToPointer(int32(1)))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteInt32Value", versionCounterKey, mock.Anything).Return(errWrite)
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

func TestImpactedSharedComponentResultModel_GetFieldDeserializers(t *testing.T) {
	model := NewImpactedSharedComponentResult()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{
		cdmSharedLibraryKey, descriptionKey, nameKey, nodeKey, nodeMainKey, stateKey,
		sysCreatedByKey, sysCreatedOnKey, sysIDKey, sysUpdatedByKey, sysUpdatedOnKey, versionCounterKey,
	} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 12)
}

func TestCreateImpactedSharedComponentResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateImpactedSharedComponentResultFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

// ---------------------------------------------------------------------------
// Response discriminator factories
// ---------------------------------------------------------------------------

func TestResponses_CreateFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name    string
		factory func() (any, error)
	}{
		{"ChangesetsResponse", func() (any, error) { return CreateChangesetsResponseFromDiscriminatorValue(nil) }},
		{"ChangesetActivityResponse", func() (any, error) { return CreateChangesetActivityResponseFromDiscriminatorValue(nil) }},
		{"CommitStatusResponse", func() (any, error) { return CreateCommitStatusResponseFromDiscriminatorValue(nil) }},
		{"ImpactedSharedComponentsResponse", func() (any, error) { return CreateImpactedSharedComponentsResponseFromDiscriminatorValue(nil) }},
		{"ImpactedDeployablesResponse", func() (any, error) { return CreateImpactedDeployablesResponseFromDiscriminatorValue(nil) }},
		{"ImpactedDeployablesBySysIDResponse", func() (any, error) { return CreateImpactedDeployablesBySysIDResponseFromDiscriminatorValue(nil) }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsable, err := tt.factory()
			require.NoError(t, err)
			assert.NotNil(t, parsable)
		})
	}
}
