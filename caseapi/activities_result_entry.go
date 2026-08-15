package caseapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	attachmentKey   = "attachment"
	containsCodeKey = "contains_code"
	elementKey      = "element"
	fieldLabelKey   = "field_label"
	initialsKey     = "initials"
	loginNameKey    = "login_name"
	nameKey         = "name"
	// sysCreatedOnKey
	// sys_created_on_adjusted
	// sys_id
	userSysIdKey = "user_sys_id"
	// valueKey
)

var _ ActivitiesResultEntry = (*ActivitiesResultEntryModel)(nil)

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

	// TODO: date time format: YYYY-MM-DD hh:mm:ss
	// GetSysCreatedOn returns the date and time of creation for the activity stream entry, expressed in GMT timezone.
	GetSysCreatedOn() (*string, error)
	// SetSysCreatedOn sets the date and time of creation for the activity stream entry, expressed in GMT timezone.
	SetSysCreatedOn(*string) error

	// TODO: date time format: YYYY-MM-DD hh:mm:ss
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

type ActivitiesResultEntryModel struct {
	core.BaseModel
}

// GetContainsCode implements [ActivitiesResultEntry].
func (a *ActivitiesResultEntryModel) GetContainsCode() (*bool, error) {
	panic("unimplemented")
}

// GetElement implements [ActivitiesResultEntry].
func (a *ActivitiesResultEntryModel) GetElement() (*string, error) {
	panic("unimplemented")
}

// GetFieldLabel implements [ActivitiesResultEntry].
func (a *ActivitiesResultEntryModel) GetFieldLabel() (*string, error) {
	panic("unimplemented")
}

// GetInitials implements [ActivitiesResultEntry].
func (a *ActivitiesResultEntryModel) GetInitials() (*string, error) {
	panic("unimplemented")
}

// GetLoginName implements [ActivitiesResultEntry].
func (a *ActivitiesResultEntryModel) GetLoginName() (*string, error) {
	panic("unimplemented")
}

// GetName implements [ActivitiesResultEntry].
func (a *ActivitiesResultEntryModel) GetName() (*string, error) {
	panic("unimplemented")
}

// GetSysCreatedOn implements [ActivitiesResultEntry].
func (a *ActivitiesResultEntryModel) GetSysCreatedOn() (*string, error) {
	panic("unimplemented")
}

// GetSysCreatedOnAdjusted implements [ActivitiesResultEntry].
func (a *ActivitiesResultEntryModel) GetSysCreatedOnAdjusted() (*string, error) {
	panic("unimplemented")
}

// GetSysID implements [ActivitiesResultEntry].
func (a *ActivitiesResultEntryModel) GetSysID() (*string, error) {
	panic("unimplemented")
}

// GetUserSysID implements [ActivitiesResultEntry].
func (a *ActivitiesResultEntryModel) GetUserSysID() (*string, error) {
	panic("unimplemented")
}

// GetValue implements [ActivitiesResultEntry].
func (a *ActivitiesResultEntryModel) GetValue() (*string, error) {
	panic("unimplemented")
}

// SetContainsCode implements [ActivitiesResultEntry].
func (a *ActivitiesResultEntryModel) SetContainsCode(*bool) error {
	panic("unimplemented")
}

// SetElement implements [ActivitiesResultEntry].
func (a *ActivitiesResultEntryModel) SetElement(*string) error {
	panic("unimplemented")
}

// SetFieldLabel implements [ActivitiesResultEntry].
func (a *ActivitiesResultEntryModel) SetFieldLabel(*string) error {
	panic("unimplemented")
}

// SetInitials implements [ActivitiesResultEntry].
func (a *ActivitiesResultEntryModel) SetInitials(*string) error {
	panic("unimplemented")
}

// SetLoginName implements [ActivitiesResultEntry].
func (a *ActivitiesResultEntryModel) SetLoginName(*string) error {
	panic("unimplemented")
}

// SetName implements [ActivitiesResultEntry].
func (a *ActivitiesResultEntryModel) SetName(*string) error {
	panic("unimplemented")
}

// SetSysCreatedOn implements [ActivitiesResultEntry].
func (a *ActivitiesResultEntryModel) SetSysCreatedOn(*string) error {
	panic("unimplemented")
}

// SetSysCreatedOnAdjusted implements [ActivitiesResultEntry].
func (a *ActivitiesResultEntryModel) SetSysCreatedOnAdjusted(*string) error {
	panic("unimplemented")
}

// SetSysID implements [ActivitiesResultEntry].
func (a *ActivitiesResultEntryModel) SetSysID(*string) error {
	panic("unimplemented")
}

// SetUserSysID implements [ActivitiesResultEntry].
func (a *ActivitiesResultEntryModel) SetUserSysID(*string) error {
	panic("unimplemented")
}

// SetValue implements [ActivitiesResultEntry].
func (a *ActivitiesResultEntryModel) SetValue(*string) error {
	panic("unimplemented")
}

// GetAttachment implements [ActivitiesResultEntry].
func (a *ActivitiesResultEntryModel) GetAttachment() (ActivitiesResultEntryAttachment, error) {
	panic("unimplemented")
}

// GetBackingStore implements [ActivitiesResultEntry].
// Subtle: this method shadows the method (BaseModel).GetBackingStore of ActivitiesResultEntryModel.BaseModel.
func (a *ActivitiesResultEntryModel) GetBackingStore() kiotaStore.BackingStore {
	panic("unimplemented")
}

// SetAttachment implements [ActivitiesResultEntry].
func (a *ActivitiesResultEntryModel) SetAttachment(ActivitiesResultEntryAttachment) error {
	panic("unimplemented")
}

// NewActivitiesResultEntry creates a new instance of [ActivitiesResultEntryModel].
func NewActivitiesResultEntry() *ActivitiesResultEntryModel {
	return &ActivitiesResultEntryModel{BaseModel: *core.NewBaseModel()}
}

// CreateActivitiesResultEntryFromDiscriminatorValue creates a new instance of [ActivitiesResultEntry].
func CreateActivitiesResultEntryFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewActivitiesResultEntry(), nil
}

// GetFieldDeserializers implements [serialization.Parsable].
func (a *ActivitiesResultEntryModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	panic("unimplemented")
}

// Serialize implements [serialization.Parsable].
func (a *ActivitiesResultEntryModel) Serialize(writer serialization.SerializationWriter) error {
	panic("unimplemented")
}
