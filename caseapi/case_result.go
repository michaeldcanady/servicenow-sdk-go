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
	// accountKey
	activeKey                  = "active"
	activeAccountEscalationKey = "active_account_escalation"
	activeEscalationKey        = "active_escalation"
	activityDueKey             = "activity_due"
	additionalAssigneeListKey  = "additional_assignee_list"
	approvalKey                = "approval"
	approvalHistoryKey         = "approval_history"
	approvalSetKey             = "approval_set"
	assetKey                   = "asset"
	//assignedToKey
	//assignmentGroupKey
	businessDurationKey          = "business_duration"
	businessImpactKey            = "business_impact"
	businessServiceKey           = "business_service"
	caseKey                      = "case"
	caseReportKey                = "case_report"
	causeKey                     = "cause"
	causedByKey                  = "caused_by"
	changeKey                    = "change"
	childCaseCreationProgressKey = "child_case_creation_progress"
	closedAtKey                  = "closed_at"
)

// CaseResult represents a single case object.
type CaseResult interface {
	serialization.Parsable
	kiotaStore.BackedModel

	// GetAccount returns the sys_id of the account record associated with the case.
	GetAccount() (Reference, error)
	// SetAccount sets the sys_id of the account record associated with the case.
	SetAccount(Reference) error

	// GetActive returns the flag that indicates whether the case is open and active.
	GetActive() (*bool, error)
	// SetActive sets the flag that indicates whether the case is open and active.
	SetActive(*bool) error

	// GetActiveAccountEscalation returns the sys_id of the active account escalation record associated with the case.
	GetActiveAccountEscalation() (*string, error)
	// SetActiveAccountEscalation sets the sys_id of the active account escalation record associated with the case.
	SetActiveAccountEscalation(*string) error

	// GetActiveEscalation returns the sys_id of the active escalation record associated with the case.
	GetActiveEscalation() (*string, error)
	// SetActiveEscalation sets the sys_id of the active escalation record associated with the case.
	SetActiveEscalation(*string) error

	// TODO: date
	// GetActivityDue returns the date for which the associated case is expected to be completed.
	GetActivityDue() (*string, error)
	// SetActivityDue sets the date for which the associated case is expected to be completed.
	SetActivityDue(*string) error

	// GetAdditionalAssigneeList returns the list of the sys_ids of the additional persons (other than the primary assignee) that have been assigned to the account.
	GetAdditionalAssigneeList() ([]*string, error)
	// SetAdditionalAssigneeList sets the list of the sys_ids of the additional persons (other than the primary assignee) that have been assigned to the account.
	SetAdditionalAssigneeList([]*string) error

	// GetApproval returns the type of approval required.
	GetApproval() (*Approval, error)
	// SetApproval sets the type of approval required.
	SetApproval(*Approval) error

	// GetApprovalHistory returns the list of all approvals associated with the case.
	GetApprovalHistory() (*string, error)
	// SetApprovalHistory sets the list of all approvals associated with the case.
	SetApprovalHistory(*string) error

	// TODO: date time
	// GetApprovalSet returns the date and time that the associated action was approved.
	GetApprovalSet() (*string, error)
	// SetApprovalSet sets the date and time that the associated action was approved.
	SetApprovalSet(*string) error

	// GetAsset returns the sys_id of the asset record associated with the case.
	GetAsset() (*string, error)
	// SetAsset sets the sys_id of the asset record associated with the case.
	SetAsset(*string) error

	// GetAssignedTo returns the sys_id of the person assigned to the case.
	GetAssignedTo() (*string, error)
	// SetAssignedTo sets the sys_id of the person assigned to the case.
	SetAssignedTo(*string) error

	// GetAssignmentGroup returns the sys_id of the customer service agent group assigned to the case.
	GetAssignmentGroup() (*string, error)
	// SetAssignmentGroup sets the sys_id of the customer service agent group assigned to the case.
	SetAssignmentGroup(*string) error

	// GetBusinessDuration returns the length in calendar work hours, work days, and work weeks that it took to complete the case.
	GetBusinessDuration() (*string, error)
	// SetBusinessDuration sets the length in calendar work hours, work days, and work weeks that it took to complete the case.
	SetBusinessDuration(*string) error

	// GetBusinessImpact returns the impact of the issue on the associated customer.
	GetBusinessImpact() (*string, error)
	// SetBusinessImpact sets the impact of the issue on the associated customer.
	SetBusinessImpact(*string) error

	// GetBusinessService returns the sys_id of the service record associated with the case.
	GetBusinessService() (*string, error)
	// SetBusinessService sets the sys_id of the service record associated with the case.
	SetBusinessService(*string) error

	// GetCase returns the case short description and case number.
	GetCase() (*string, error)
	// SetCase sets the case short description and case number.
	SetCase(*string) error

	// GetCaseReport returns the sys_id of the associated case report.
	GetCaseReport() (*string, error)
	// SetCaseReport sets the sys_id of the associated case report.
	SetCaseReport(*string) error

	// GetCategory returns the case category.
	GetCategory() (*Category, error)
	// GetCategory sets the case category.
	SetCategory(*Category) error

	// GetCause returns the details about the cause of the problem.
	GetCause() (*string, error)
	// SetCause sets the details about the cause of the problem.
	SetCause(*string) error

	// GetCausedBy returns the sys_id of the change request that caused the case to be created.
	GetCausedBy() (*string, error)
	// SetCausedBy sets the sys_id of the change request that caused the case to be created.
	SetCausedBy(*string) error

	// GetChange returns the sys_id of the change request that caused the case to be created.
	GetChange() (*string, error)
	// SetChange sets the sys_id of the change request that caused the case to be created.
	SetChange(*string) error

	// GetChildCaseCreationProgress returns the flag that indicates whether the case is a child case that was created from a major case.
	GetChildCaseCreationProgress() (*bool, error)
	// SetChildCaseCreationProgress sets the flag that indicates whether the case is a child case that was created from a major case.
	SetChildCaseCreationProgress(*bool) error

	// TODO: date time
	// GetClosedAt returns the date and time that the case was closed.
	GetClosedAt() (*string, error)
	// SetClosedAt sets the date and time that the case was closed.
	SetClosedAt(*string) error

	// TODO: continue from closed_by
}

// CaseResultModel implementation of CaseResult.
type CaseResultModel struct {
	core.BaseModel
}

// NewCaseResult creates a new instance of CaseResult.
func NewCaseResult() *CaseResultModel {
	return &CaseResultModel{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the objects properties to the current writer.
func (m *CaseResultModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIDKey, m.GetSysID),
		internalSerialization.SerializeStringFunc(numberKey, m.GetNumber),
		internalSerialization.SerializeStringFunc(shortDescriptionKey, m.GetShortDescription),
		internalSerialization.SerializeStringFunc(descriptionKey, m.GetDescription),
		internalSerialization.SerializeStringFunc(stateKey, m.GetState),
		internalSerialization.SerializeStringFunc(priorityKey, m.GetPriority),
		internalSerialization.SerializeStringFunc(categoryKey, m.GetCategory),
		internalSerialization.SerializeObjectValueFunc[Reference](assignmentGroupKey, m.GetAssignmentGroup),
		internalSerialization.SerializeObjectValueFunc[Reference](assignedToKey, m.GetAssignedTo),
		internalSerialization.SerializeObjectValueFunc[Reference](contactKey, m.GetContact),
		internalSerialization.SerializeObjectValueFunc[Reference](accountKey, m.GetAccount),
		internalSerialization.SerializeStringFunc(sysCreatedOnKey, m.GetSysCreatedOn),
		internalSerialization.SerializeStringFunc(sysUpdatedOnKey, m.GetSysUpdatedOn),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *CaseResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIDKey:            internalSerialization.DeserializeStringFunc(m.setSysID),
		numberKey:           internalSerialization.DeserializeStringFunc(m.setNumber),
		shortDescriptionKey: internalSerialization.DeserializeStringFunc(m.setShortDescription),
		descriptionKey:      internalSerialization.DeserializeStringFunc(m.setDescription),
		stateKey:            internalSerialization.DeserializeStringFunc(m.setState),
		priorityKey:         internalSerialization.DeserializeStringFunc(m.setPriority),
		categoryKey:         internalSerialization.DeserializeStringFunc(m.setCategory),
		assignmentGroupKey:  internalSerialization.DeserializeObjectValueFunc[Reference](CreateReferenceFromDiscriminatorValue, m.setAssignmentGroup),
		assignedToKey:       internalSerialization.DeserializeObjectValueFunc[Reference](CreateReferenceFromDiscriminatorValue, m.setAssignedTo),
		contactKey:          internalSerialization.DeserializeObjectValueFunc[Reference](CreateReferenceFromDiscriminatorValue, m.setContact),
		accountKey:          internalSerialization.DeserializeObjectValueFunc[Reference](CreateReferenceFromDiscriminatorValue, m.setAccount),
		sysCreatedOnKey:     internalSerialization.DeserializeStringFunc(m.setSysCreatedOn),
		sysUpdatedOnKey:     internalSerialization.DeserializeStringFunc(m.setSysUpdatedOn),
	}
}

// GetSysID ...
func (m *CaseResultModel) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, sysIDKey)
}
func (m *CaseResultModel) setSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIDKey, val)
}

// GetNumber ...
func (m *CaseResultModel) GetNumber() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, numberKey)
}
func (m *CaseResultModel) setNumber(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, numberKey, val)
}

// GetShortDescription ...
func (m *CaseResultModel) GetShortDescription() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, shortDescriptionKey)
}
func (m *CaseResultModel) setShortDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, shortDescriptionKey, val)
}

// GetDescription ...
func (m *CaseResultModel) GetDescription() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, descriptionKey)
}
func (m *CaseResultModel) setDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, descriptionKey, val)
}

// GetState ...
func (m *CaseResultModel) GetState() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, stateKey)
}
func (m *CaseResultModel) setState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, stateKey, val)
}

// GetPriority ...
func (m *CaseResultModel) GetPriority() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, priorityKey)
}
func (m *CaseResultModel) setPriority(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, priorityKey, val)
}

// GetCategory ...
func (m *CaseResultModel) GetCategory() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, categoryKey)
}
func (m *CaseResultModel) setCategory(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, categoryKey, val)
}

// GetAssignmentGroup ...
func (m *CaseResultModel) GetAssignmentGroup() (Reference, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, Reference](m, assignmentGroupKey)
}
func (m *CaseResultModel) setAssignmentGroup(val Reference) error {
	return store.DefaultBackedModelMutatorFunc(m, assignmentGroupKey, val)
}

// GetAssignedTo ...
func (m *CaseResultModel) GetAssignedTo() (Reference, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, Reference](m, assignedToKey)
}
func (m *CaseResultModel) setAssignedTo(val Reference) error {
	return store.DefaultBackedModelMutatorFunc(m, assignedToKey, val)
}

// GetContact ...
func (m *CaseResultModel) GetContact() (Reference, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, Reference](m, contactKey)
}
func (m *CaseResultModel) setContact(val Reference) error {
	return store.DefaultBackedModelMutatorFunc(m, contactKey, val)
}

// GetAccount ...
func (m *CaseResultModel) GetAccount() (Reference, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, Reference](m, accountKey)
}
func (m *CaseResultModel) setAccount(val Reference) error {
	return store.DefaultBackedModelMutatorFunc(m, accountKey, val)
}

// GetSysCreatedOn ...
func (m *CaseResultModel) GetSysCreatedOn() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, sysCreatedOnKey)
}
func (m *CaseResultModel) setSysCreatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysCreatedOnKey, val)
}

// GetSysUpdatedOn ...
func (m *CaseResultModel) GetSysUpdatedOn() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, sysUpdatedOnKey)
}
func (m *CaseResultModel) setSysUpdatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysUpdatedOnKey, val)
}

// CreateCaseResultFromDiscriminatorValue creates a new instance of CaseResult.
func CreateCaseResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCaseResult(), nil
}
