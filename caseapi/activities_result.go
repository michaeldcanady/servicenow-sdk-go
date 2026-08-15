package caseapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	tableKey                = "table"
	displayValueKey         = "display_value"
	userFullNameKey         = "user_full_name"
	sysCreatedOnAdjustedKey = "sys_created_on_adjusted"
	thumbnailPathKey        = "thumbnail_path"
	userLoginKey            = "user_login"
	userSysIDKey            = "user_sys_id"
	entriesKey              = "entries"
	journalFieldsKey        = "journal_fields"
)

var _ ActivitiesResult = (*ActivitiesResultModel)(nil)

// ActivitiesResult represents case activities.
type ActivitiesResult interface {
	serialization.Parsable
	kiotaStore.BackedModel

	// GetDisplayValue returns the display value for the case number for the CSM case.
	GetDisplayValue() (*string, error)
	// SetDisplayValue sets the display value for the case number for the CSM case.
	SetDisplayValue(*string) error
	// GetJournalFields returns the journal fields present in the CSM case.
	GetJournalFields() ([]ActivitiesResultField, error)
	// SetJournalFields sets the journal fields present in the CSM case.
	SetJournalFields([]ActivitiesResultField) error
	// TODO: is technically an enum but only one possible value?
	// GetLabel returns the display label for the table in which the CSM case record appears.
	GetLabel() (*string, error)
	// SetLabel sets the display label for the table in which the CSM case record appears.
	SetLabel(*string) error
	// GetNumber returns the case number for the CSM case.
	GetNumber() (*string, error)
	// SetNumber sets the case number for the CSM case.
	SetNumber(*string) error

	// GetShortDescription returns the short description for the CSM case.
	GetShortDescription() (*string, error)
	// SetShortDescription sets the short description for the CSM case.
	SetShortDescription(*string) error
	GetSysID() (*string, error)
	SetSysID(*string) error
	// TODO: Date time
	GetSysCreatedOn() (*string, error)
	SetSysCreatedOn(*string) error

	// TODO: datetime YYYY-MM-DD hh:mm:ss
	// GetSysCreatedOnAdjusted returns the date and time of case creation, expressed in logged-in user timezone.
	GetSysCreatedOnAdjusted() (*string, error)
	// GetSysCreatedOnAdjusted sets the date and time of case creation, expressed in logged-in user timezone.
	SetSysCreatedOnAdjusted(*string) error

	// GetTable returns the name of the table in which the CSM case appears.
	GetTable() (*string, error)
	// SetTable sets the name of the table in which the CSM case appears.
	SetTable(*string) error

	// TODO: should known urls like this be http.url?
	// GetThumbnailPath returns the direct link to get thumbnail for attachment (only for attachment of type images.)
	GetThumbnailPath() (*string, error)
	// SetThumbnailPath sets the direct link to get thumbnail for attachment (only for attachment of type images.)
	SetThumbnailPath(*string) error

	// GetUserFullName returns the full name for logged-in user.
	GetUserFullName() (*string, error)
	// SetUserFullName sets the full name for logged-in user.
	SetUserFullName(*string) error

	// GetUserLogin returns the user ID for logged-in user.
	GetUserLogin() (*string, error)
	// SetUserLogin sets the user ID for logged-in user.
	SetUserLogin(*string) error

	// GetUserSysID returns the sys_id for logged-in user.
	GetUserSysID() (*string, error)
	// SetUserSysID sets the sys_id for logged-in user.
	SetUserSysID(*string) error

	// GetEntries returns the list that represents activity stream entries from the CSM case.
	GetEntries() ([]ActivitiesResultEntry, error)
	// SetEntries sets the list that represents activity stream entries from the CSM case.
	SetEntries([]ActivitiesResultEntry) error
}

// ActivitiesResultModel implementation of [ActivitiesResult].
type ActivitiesResultModel struct {
	core.BaseModel
}

// NewActivitiesResult creates a new instance of [ActivitiesResultModel].
func NewActivitiesResult() *ActivitiesResultModel {
	return &ActivitiesResultModel{BaseModel: *core.NewBaseModel()}
}

// CreateActivitiesResultFromDiscriminatorValue creates a new instance of [ActivitiesResult].
func CreateActivitiesResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewActivitiesResult(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *ActivitiesResultModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(displayValueKey, m.GetDisplayValue),
		internalSerialization.SerializeCollectionOfObjectValuesFunc(entriesKey, m.GetEntries),
		internalSerialization.SerializeCollectionOfObjectValuesFunc(journalFieldsKey, m.GetJournalFields),
		internalSerialization.SerializeStringFunc(labelKey, m.GetLabel),
		internalSerialization.SerializeStringFunc(numberKey, m.GetNumber),
		internalSerialization.SerializeStringFunc(shortDescriptionKey, m.GetShortDescription),
		internalSerialization.SerializeStringFunc(sysCreatedOnKey, m.GetSysCreatedOn),
		internalSerialization.SerializeStringFunc(sysCreatedOnAdjustedKey, m.GetSysCreatedOnAdjusted),
		internalSerialization.SerializeStringFunc(sysIDKey, m.GetSysID),
		internalSerialization.SerializeStringFunc(tableKey, m.GetTable),
		internalSerialization.SerializeStringFunc(thumbnailPathKey, m.GetThumbnailPath),
		internalSerialization.SerializeStringFunc(userFullNameKey, m.GetUserFullName),
		internalSerialization.SerializeStringFunc(userLoginKey, m.GetUserLogin),
		internalSerialization.SerializeStringFunc(userSysIDKey, m.GetUserSysID),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *ActivitiesResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		displayValueKey:         internalSerialization.DeserializeStringFunc(m.SetDisplayValue),
		entriesKey:              internalSerialization.DeserializeCollectionOfObjectValuesFunc(CreateActivitiesResultEntryFromDiscriminatorValue, m.SetEntries),
		journalFieldsKey:        internalSerialization.DeserializeCollectionOfObjectValuesFunc(CreateActivitiesResultFieldFromDiscriminatorValue, m.SetJournalFields),
		labelKey:                internalSerialization.DeserializeStringFunc(m.SetLabel),
		numberKey:               internalSerialization.DeserializeStringFunc(m.SetNumber),
		shortDescriptionKey:     internalSerialization.DeserializeStringFunc(m.SetShortDescription),
		sysCreatedOnKey:         internalSerialization.DeserializeStringFunc(m.SetSysCreatedOn),
		sysCreatedOnAdjustedKey: internalSerialization.DeserializeStringFunc(m.SetSysCreatedOnAdjusted),
		sysIDKey:                internalSerialization.DeserializeStringFunc(m.SetSysID),
		tableKey:                internalSerialization.DeserializeStringFunc(m.SetTable),
		thumbnailPathKey:        internalSerialization.DeserializeStringFunc(m.SetThumbnailPath),
		userFullNameKey:         internalSerialization.DeserializeStringFunc(m.SetUserFullName),
		userLoginKey:            internalSerialization.DeserializeStringFunc(m.SetUserLogin),
		userSysIDKey:            internalSerialization.DeserializeStringFunc(m.SetUserSysID),
	}
}

// GetEntries returns the list that represents activity stream entries from the CSM case.
func (m *ActivitiesResultModel) GetEntries() ([]ActivitiesResultEntry, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, []ActivitiesResultEntry](m, entriesKey)
}

// SetEntries sets the list that represents activity stream entries from the CSM case.
func (m *ActivitiesResultModel) SetEntries(val []ActivitiesResultEntry) error {
	return store.DefaultBackedModelMutatorFunc(m, entriesKey, val)
}

// GetJournalFields returns the list of journal field names present in the CSM case.
func (m *ActivitiesResultModel) GetJournalFields() ([]ActivitiesResultField, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, []ActivitiesResultField](m, journalFieldsKey)
}

// SetJournalFields sets the list of journal field names present in the CSM case.
func (m *ActivitiesResultModel) SetJournalFields(val []ActivitiesResultField) error {
	return store.DefaultBackedModelMutatorFunc(m, journalFieldsKey, val)
}

// GetDisplayValue implements [ActivitiesResult].
func (m *ActivitiesResultModel) GetDisplayValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, displayValueKey)
}

// GetLabel implements [ActivitiesResult].
func (m *ActivitiesResultModel) GetLabel() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, labelKey)
}

// GetNumber implements [ActivitiesResult].
func (m *ActivitiesResultModel) GetNumber() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, numberKey)
}

// GetShortDescription implements [ActivitiesResult].
func (m *ActivitiesResultModel) GetShortDescription() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, shortDescriptionKey)
}

// GetSysCreatedOnAdjusted implements [ActivitiesResult].
func (m *ActivitiesResultModel) GetSysCreatedOnAdjusted() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, sysCreatedOnAdjustedKey)
}

// GetTable implements [ActivitiesResult].
func (m *ActivitiesResultModel) GetTable() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, tableKey)
}

// GetThumbnailPath implements [ActivitiesResult].
func (m *ActivitiesResultModel) GetThumbnailPath() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, thumbnailPathKey)
}

// GetUserFullName implements [ActivitiesResult].
func (m *ActivitiesResultModel) GetUserFullName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, userFullNameKey)
}

// GetUserLogin implements [ActivitiesResult].
func (m *ActivitiesResultModel) GetUserLogin() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, userLoginKey)
}

// GetUserSysID implements [ActivitiesResult].
func (m *ActivitiesResultModel) GetUserSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, userSysIDKey)
}

// SetDisplayValue implements [ActivitiesResult].
func (m *ActivitiesResultModel) SetDisplayValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, displayValueKey, val)
}

// SetLabel implements [ActivitiesResult].
func (m *ActivitiesResultModel) SetLabel(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, labelKey, val)
}

// SetNumber implements [ActivitiesResult].
func (m *ActivitiesResultModel) SetNumber(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, numberKey, val)
}

// SetShortDescription implements [ActivitiesResult].
func (m *ActivitiesResultModel) SetShortDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, shortDescriptionKey, val)
}

// SetSysCreatedOn implements [ActivitiesResult].
func (m *ActivitiesResultModel) SetSysCreatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysCreatedOnKey, val)
}

// SetSysCreatedOnAdjusted implements [ActivitiesResult].
func (m *ActivitiesResultModel) SetSysCreatedOnAdjusted(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysCreatedOnAdjustedKey, val)
}

// SetSysID implements [ActivitiesResult].
func (m *ActivitiesResultModel) SetSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIDKey, val)
}

// SetTable implements [ActivitiesResult].
func (m *ActivitiesResultModel) SetTable(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, tableKey, val)
}

// SetThumbnailPath implements [ActivitiesResult].
func (m *ActivitiesResultModel) SetThumbnailPath(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, thumbnailPathKey, val)
}

// SetUserFullName implements [ActivitiesResult].
func (m *ActivitiesResultModel) SetUserFullName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, userFullNameKey, val)
}

// SetUserLogin implements [ActivitiesResult].
func (m *ActivitiesResultModel) SetUserLogin(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, userLoginKey, val)
}

// SetUserSysID implements [ActivitiesResult].
func (m *ActivitiesResultModel) SetUserSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, userSysIDKey, val)
}

// GetSysID ...
func (m *ActivitiesResultModel) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, sysIDKey)
}

// GetSysCreatedOn ...
func (m *ActivitiesResultModel) GetSysCreatedOn() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, sysCreatedOnKey)
}
