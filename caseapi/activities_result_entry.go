package caseapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	attachmentKey      = "attachment"
	containsCodeKey    = "contains_code"
	elementKey         = "element"
	fieldLabelKey      = "field_label"
	initialsKey        = "initials"
	loginNameKey       = "login_name"
	nameKey            = "name"
	sysCreatedOnAdjKey = "sys_created_on_adjusted"
	userSysIdKey       = "user_sys_id"
)

var _ ActivitiesResultEntry = (*ActivitiesResultEntryModel)(nil)

// ActivitiesResultEntry represents a single activity stream entry from a CSM case.
type ActivitiesResultEntry interface {
	serialization.Parsable
	kiotaStore.BackedModel

	// GetAttachment returns the description of a file attached to the CSM case in the activity stream entry.
	GetAttachment() (ActivitiesResultEntryAttachment, error)
	// SetAttachment sets the description of a file attached to the CSM case in the activity stream entry.
	SetAttachment(ActivitiesResultEntryAttachment) error

	// GetContainsCode return indicates whether the activity stream entry contains code or not.
	GetContainsCode() (*bool, error)
	// SetContainsCode sets whether the activity stream entry contains code or not.
	SetContainsCode(*bool) error

	// GetElement returns the name for the journal field associated with the activity stream entry.
	GetElement() (*string, error)
	// SetElement sets the name for the journal field associated with the activity stream entry.
	SetElement(*string) error

	// GetFieldLabel returns the display name for the journal field associated with the activity stream entry.
	GetFieldLabel() (*string, error)
	// SetFieldLabel sets the display name for the journal field associated with the activity stream entry.
	SetFieldLabel(*string) error

	// GetInitials returns the initials of the user who created the activity stream entry.
	GetInitials() (*string, error)
	// SetInitials sets the initials of the user who created the activity stream entry.
	SetInitials(*string) error

	// GetLoginName returns the full user name for the user who created the activity stream entry.
	GetLoginName() (*string, error)
	// SetLoginName sets the full user name for the user who created the activity stream entry.
	SetLoginName(*string) error

	// GetName returns the full user name for the user who created the activity stream entry.
	GetName() (*string, error)
	// SetName sets the full user name for the user who created the activity stream entry.
	SetName(*string) error

	// GetSysCreatedOn returns the date and time of creation for the activity stream entry, expressed in GMT timezone.
	GetSysCreatedOn() (*string, error)
	// SetSysCreatedOn sets the date and time of creation for the activity stream entry, expressed in GMT timezone.
	SetSysCreatedOn(*string) error

	// GetSysCreatedOnAdjusted returns the date and time of creation for the activity stream entry, expressed in logged-in user timezone.
	GetSysCreatedOnAdjusted() (*string, error)
	// SetSysCreatedOnAdjusted sets the date and time of creation for the activity stream entry, expressed in logged-in user timezone.
	SetSysCreatedOnAdjusted(*string) error

	// GetSysID returns the sys_id for the activity stream entry.
	GetSysID() (*string, error)
	// SetSysID sets the sys_id for the activity stream entry.
	SetSysID(*string) error

	// GetUserSysID returns the sys_id for user who created the activity stream entry.
	GetUserSysID() (*string, error)
	// SetUserSysID sets the sys_id for user who created the activity stream entry.
	SetUserSysID(*string) error

	// GetValue returns the value for this journal entry.
	GetValue() (*string, error)
	// SetValue sets the value for this journal entry.
	SetValue(*string) error
}

// ActivitiesResultEntryModel is the implementation of [ActivitiesResultEntry].
type ActivitiesResultEntryModel struct {
	core.BaseModel
}

// NewActivitiesResultEntry creates a new instance of [ActivitiesResultEntryModel].
func NewActivitiesResultEntry() *ActivitiesResultEntryModel {
	return &ActivitiesResultEntryModel{BaseModel: *core.NewBaseModel()}
}

// CreateActivitiesResultEntryFromDiscriminatorValue creates a new instance of [ActivitiesResultEntry].
func CreateActivitiesResultEntryFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewActivitiesResultEntry(), nil
}

// Serialize writes the objects properties to the current writer.
func (a *ActivitiesResultEntryModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(a) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeObjectValueFunc(attachmentKey, a.GetAttachment),
		internalSerialization.SerializeBoolFunc(containsCodeKey, a.GetContainsCode),
		internalSerialization.SerializeStringFunc(elementKey, a.GetElement),
		internalSerialization.SerializeStringFunc(fieldLabelKey, a.GetFieldLabel),
		internalSerialization.SerializeStringFunc(initialsKey, a.GetInitials),
		internalSerialization.SerializeStringFunc(loginNameKey, a.GetLoginName),
		internalSerialization.SerializeStringFunc(nameKey, a.GetName),
		internalSerialization.SerializeStringFunc(sysCreatedOnKey, a.GetSysCreatedOn),
		internalSerialization.SerializeStringFunc(sysCreatedOnAdjKey, a.GetSysCreatedOnAdjusted),
		internalSerialization.SerializeStringFunc(sysIDKey, a.GetSysID),
		internalSerialization.SerializeStringFunc(userSysIdKey, a.GetUserSysID),
		internalSerialization.SerializeStringFunc(valueKey, a.GetValue),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (a *ActivitiesResultEntryModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		attachmentKey:      internalSerialization.DeserializeObjectValueFunc(CreateActivitiesResultEntryAttachmentFromDiscriminatorValue, a.SetAttachment),
		containsCodeKey:    internalSerialization.DeserializeBoolFunc(a.SetContainsCode),
		elementKey:         internalSerialization.DeserializeStringFunc(a.SetElement),
		fieldLabelKey:      internalSerialization.DeserializeStringFunc(a.SetFieldLabel),
		initialsKey:        internalSerialization.DeserializeStringFunc(a.SetInitials),
		loginNameKey:       internalSerialization.DeserializeStringFunc(a.SetLoginName),
		nameKey:            internalSerialization.DeserializeStringFunc(a.SetName),
		sysCreatedOnKey:    internalSerialization.DeserializeStringFunc(a.SetSysCreatedOn),
		sysCreatedOnAdjKey: internalSerialization.DeserializeStringFunc(a.SetSysCreatedOnAdjusted),
		sysIDKey:           internalSerialization.DeserializeStringFunc(a.SetSysID),
		userSysIdKey:       internalSerialization.DeserializeStringFunc(a.SetUserSysID),
		valueKey:           internalSerialization.DeserializeStringFunc(a.SetValue),
	}
}

// GetBackingStore delegates to the embedded BaseModel.
func (a *ActivitiesResultEntryModel) GetBackingStore() kiotaStore.BackingStore {
	return a.BaseModel.GetBackingStore()
}

// GetAttachment returns the attachment for this activity entry.
func (a *ActivitiesResultEntryModel) GetAttachment() (ActivitiesResultEntryAttachment, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultEntryModel, ActivitiesResultEntryAttachment](a, attachmentKey)
}

// SetAttachment sets the attachment for this activity entry.
func (a *ActivitiesResultEntryModel) SetAttachment(val ActivitiesResultEntryAttachment) error {
	return store.DefaultBackedModelMutatorFunc(a, attachmentKey, val)
}

// GetContainsCode returns whether the entry contains code.
func (a *ActivitiesResultEntryModel) GetContainsCode() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultEntryModel, *bool](a, containsCodeKey)
}

// SetContainsCode sets whether the entry contains code.
func (a *ActivitiesResultEntryModel) SetContainsCode(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(a, containsCodeKey, val)
}

// GetElement returns the journal field name for this entry.
func (a *ActivitiesResultEntryModel) GetElement() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultEntryModel, *string](a, elementKey)
}

// SetElement sets the journal field name for this entry.
func (a *ActivitiesResultEntryModel) SetElement(val *string) error {
	return store.DefaultBackedModelMutatorFunc(a, elementKey, val)
}

// GetFieldLabel returns the display name for the journal field.
func (a *ActivitiesResultEntryModel) GetFieldLabel() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultEntryModel, *string](a, fieldLabelKey)
}

// SetFieldLabel sets the display name for the journal field.
func (a *ActivitiesResultEntryModel) SetFieldLabel(val *string) error {
	return store.DefaultBackedModelMutatorFunc(a, fieldLabelKey, val)
}

// GetInitials returns the initials of the user who created this entry.
func (a *ActivitiesResultEntryModel) GetInitials() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultEntryModel, *string](a, initialsKey)
}

// SetInitials sets the initials of the user who created this entry.
func (a *ActivitiesResultEntryModel) SetInitials(val *string) error {
	return store.DefaultBackedModelMutatorFunc(a, initialsKey, val)
}

// GetLoginName returns the login name of the user who created this entry.
func (a *ActivitiesResultEntryModel) GetLoginName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultEntryModel, *string](a, loginNameKey)
}

// SetLoginName sets the login name of the user who created this entry.
func (a *ActivitiesResultEntryModel) SetLoginName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(a, loginNameKey, val)
}

// GetName returns the full name of the user who created this entry.
func (a *ActivitiesResultEntryModel) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultEntryModel, *string](a, nameKey)
}

// SetName sets the full name of the user who created this entry.
func (a *ActivitiesResultEntryModel) SetName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(a, nameKey, val)
}

// GetSysCreatedOn returns the GMT creation time.
func (a *ActivitiesResultEntryModel) GetSysCreatedOn() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultEntryModel, *string](a, sysCreatedOnKey)
}

// SetSysCreatedOn sets the GMT creation time.
func (a *ActivitiesResultEntryModel) SetSysCreatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(a, sysCreatedOnKey, val)
}

// GetSysCreatedOnAdjusted returns the creation time in the logged-in user's timezone.
func (a *ActivitiesResultEntryModel) GetSysCreatedOnAdjusted() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultEntryModel, *string](a, sysCreatedOnAdjKey)
}

// SetSysCreatedOnAdjusted sets the creation time in the logged-in user's timezone.
func (a *ActivitiesResultEntryModel) SetSysCreatedOnAdjusted(val *string) error {
	return store.DefaultBackedModelMutatorFunc(a, sysCreatedOnAdjKey, val)
}

// GetSysID returns the sys_id of this entry.
func (a *ActivitiesResultEntryModel) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultEntryModel, *string](a, sysIDKey)
}

// SetSysID sets the sys_id of this entry.
func (a *ActivitiesResultEntryModel) SetSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(a, sysIDKey, val)
}

// GetUserSysID returns the sys_id of the user who created this entry.
func (a *ActivitiesResultEntryModel) GetUserSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultEntryModel, *string](a, userSysIdKey)
}

// SetUserSysID sets the sys_id of the user who created this entry.
func (a *ActivitiesResultEntryModel) SetUserSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(a, userSysIdKey, val)
}

// GetValue returns the journal entry text value.
func (a *ActivitiesResultEntryModel) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultEntryModel, *string](a, valueKey)
}

// SetValue sets the journal entry text value.
func (a *ActivitiesResultEntryModel) SetValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(a, valueKey, val)
}
