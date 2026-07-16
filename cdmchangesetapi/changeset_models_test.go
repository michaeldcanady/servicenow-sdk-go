package cdmchangesetapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/stretchr/testify/assert"
)

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
		{"SysId", func(v interface{}) error { return model.setSysId(v.(*string)) }, func() (interface{}, error) { return model.GetSysId() }, internal.ToPointer("sys-id")},
		{"Title", func(v interface{}) error { return model.setTitle(v.(*string)) }, func() (interface{}, error) { return model.GetTitle() }, internal.ToPointer("Changeset Title")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			assert.NoError(t, err)
			got, err := tt.getter()
			assert.NoError(t, err)
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
		{"ChangesetId", func(v interface{}) error { return model.setChangesetId(v.(*Reference)) }, func() (interface{}, error) { return model.GetChangesetId() }, NewReference()},
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
			assert.NoError(t, err)
			got, err := tt.getter()
			assert.NoError(t, err)
			assert.Equal(t, tt.value, got)
		})
	}
}

func TestCreateChangesetResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateChangesetResultFromDiscriminatorValue(nil)
	assert.NoError(t, err)
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
			assert.NoError(t, err)
			got, err := tt.getter()
			assert.NoError(t, err)
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
		{"SysId", func(v interface{}) error { return model.setSysId(v.(*string)) }, func() (interface{}, error) { return model.GetSysId() }, internal.ToPointer("sys-id")},
		{"SysUpdatedBy", func(v interface{}) error { return model.setSysUpdatedBy(v.(*string)) }, func() (interface{}, error) { return model.GetSysUpdatedBy() }, internal.ToPointer("admin")},
		{"SysUpdatedOn", func(v interface{}) error { return model.setSysUpdatedOn(v.(*string)) }, func() (interface{}, error) { return model.GetSysUpdatedOn() }, internal.ToPointer("2023-01-01")},
		{"VersionCounter", func(v interface{}) error { return model.setVersionCounter(v.(*int32)) }, func() (interface{}, error) { return model.GetVersionCounter() }, internal.ToPointer(int32(1))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			assert.NoError(t, err)
			got, err := tt.getter()
			assert.NoError(t, err)
			assert.Equal(t, tt.value, got)
		})
	}
}

func TestCreateCommitStatusResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateCommitStatusResultFromDiscriminatorValue(nil)
	assert.NoError(t, err)
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
		{"SysId", func(v interface{}) error { return model.setSysId(v.(*string)) }, func() (interface{}, error) { return model.GetSysId() }, internal.ToPointer("sys-id")},
		{"SysCreatedBy", func(v interface{}) error { return model.setSysCreatedBy(v.(*string)) }, func() (interface{}, error) { return model.GetSysCreatedBy() }, internal.ToPointer("admin")},
		{"SysCreatedOn", func(v interface{}) error { return model.setSysCreatedOn(v.(*string)) }, func() (interface{}, error) { return model.GetSysCreatedOn() }, internal.ToPointer("2023-01-01")},
		{"SysUpdatedBy", func(v interface{}) error { return model.setSysUpdatedBy(v.(*string)) }, func() (interface{}, error) { return model.GetSysUpdatedBy() }, internal.ToPointer("admin")},
		{"SysUpdatedOn", func(v interface{}) error { return model.setSysUpdatedOn(v.(*string)) }, func() (interface{}, error) { return model.GetSysUpdatedOn() }, internal.ToPointer("2023-01-01")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			assert.NoError(t, err)
			got, err := tt.getter()
			assert.NoError(t, err)
			assert.Equal(t, tt.value, got)
		})
	}
}

func TestImpactedDeployableBySysIdResultModel_GettersSetters(t *testing.T) {
	model := NewImpactedDeployableBySysIdResult()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"ChangesetId", func(v interface{}) error { return model.setChangesetId(v.(*string)) }, func() (interface{}, error) { return model.GetChangesetId() }, internal.ToPointer("chg-id")},
		{"Conflict", func(v interface{}) error { return model.setConflict(v.(*bool)) }, func() (interface{}, error) { return model.GetConflict() }, internal.ToPointer(true)},
		{"ConflictType", func(v interface{}) error { return model.setConflictType(v.(*string)) }, func() (interface{}, error) { return model.GetConflictType() }, internal.ToPointer("type")},
		{"Description", func(v interface{}) error { return model.setDescription(v.(*string)) }, func() (interface{}, error) { return model.GetDescription() }, internal.ToPointer("desc")},
		{"EffectiveFrom", func(v interface{}) error { return model.setEffectiveFrom(v.(*string)) }, func() (interface{}, error) { return model.GetEffectiveFrom() }, internal.ToPointer("2023-01-01")},
		{"EffectiveTo", func(v interface{}) error { return model.setEffectiveTo(v.(*string)) }, func() (interface{}, error) { return model.GetEffectiveTo() }, internal.ToPointer("2023-12-31")},
		{"Level", func(v interface{}) error { return model.setLevel(v.(*int32)) }, func() (interface{}, error) { return model.GetLevel() }, internal.ToPointer(int32(1))},
		{"LinkedTo", func(v interface{}) error { return model.setLinkedTo(v.(*string)) }, func() (interface{}, error) { return model.GetLinkedTo() }, internal.ToPointer("link")},
		{"MainId", func(v interface{}) error { return model.setMainId(v.(*string)) }, func() (interface{}, error) { return model.GetMainId() }, internal.ToPointer("main-id")},
		{"MainIdEncoded", func(v interface{}) error { return model.setMainIdEncoded(v.(*string)) }, func() (interface{}, error) { return model.GetMainIdEncoded() }, internal.ToPointer("encoded")},
		{"Name", func(v interface{}) error { return model.setName(v.(*string)) }, func() (interface{}, error) { return model.GetName() }, internal.ToPointer("name")},
		{"NodeClassifier", func(v interface{}) error { return model.setNodeClassifier(v.(*string)) }, func() (interface{}, error) { return model.GetNodeClassifier() }, internal.ToPointer("classifier")},
		{"Status", func(v interface{}) error { return model.setStatus(v.(*string)) }, func() (interface{}, error) { return model.GetStatus() }, internal.ToPointer("status")},
		{"SysId", func(v interface{}) error { return model.setSysId(v.(*string)) }, func() (interface{}, error) { return model.GetSysId() }, internal.ToPointer("sys-id")},
		{"Type", func(v interface{}) error { return model.setType(v.(*string)) }, func() (interface{}, error) { return model.GetType() }, internal.ToPointer("type")},
		{"Value", func(v interface{}) error { return model.setValue(v.(*string)) }, func() (interface{}, error) { return model.GetValue() }, internal.ToPointer("value")},
		{"SecureValue", func(v interface{}) error { return model.setSecureValue(v.(*string)) }, func() (interface{}, error) { return model.GetSecureValue() }, internal.ToPointer("secure")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			assert.NoError(t, err)
			got, err := tt.getter()
			assert.NoError(t, err)
			assert.Equal(t, tt.value, got)
		})
	}
}

func TestCreateImpactedDeployableResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateImpactedDeployableResultFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateImpactedDeployableBySysIdResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateImpactedDeployableBySysIdResultFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, parsable)
}
