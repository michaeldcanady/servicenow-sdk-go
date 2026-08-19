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

	// GetAssignedTo returns the person assigned to the case.
	GetAssignedTo() (Reference, error)
	// SetAssignedTo sets the person assigned to the case.
	SetAssignedTo(Reference) error

	// GetAssignmentGroup returns the customer service agent group assigned to the case.
	GetAssignmentGroup() (Reference, error)
	// SetAssignmentGroup sets the customer service agent group assigned to the case.
	SetAssignmentGroup(Reference) error

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
	GetCategory() (*string, error)
	// SetCategory sets the case category.
	SetCategory(*string) error

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

	// GetContact returns the contact associated with the case.
	GetContact() (Reference, error)
	// SetContact sets the contact associated with the case.
	SetContact(Reference) error

	// GetDescription returns the case description.
	GetDescription() (*string, error)
	// SetDescription sets the case description.
	SetDescription(*string) error

	// GetNumber returns the case number.
	GetNumber() (*string, error)
	// SetNumber sets the case number.
	SetNumber(*string) error

	// GetPriority returns the case priority.
	GetPriority() (*string, error)
	// SetPriority sets the case priority.
	SetPriority(*string) error

	// GetShortDescription returns the short description of the case.
	GetShortDescription() (*string, error)
	// SetShortDescription sets the short description of the case.
	SetShortDescription(*string) error

	// GetState returns the state of the case.
	GetState() (*string, error)
	// SetState sets the state of the case.
	SetState(*string) error

	// GetSysCreatedOn returns the creation timestamp.
	GetSysCreatedOn() (*string, error)
	// SetSysCreatedOn sets the creation timestamp.
	SetSysCreatedOn(*string) error

	// GetSysID returns the sys_id of the case.
	GetSysID() (*string, error)
	// SetSysID sets the sys_id of the case.
	SetSysID(*string) error

	// GetSysUpdatedOn returns the last-updated timestamp.
	GetSysUpdatedOn() (*string, error)
	// SetSysUpdatedOn sets the last-updated timestamp.
	SetSysUpdatedOn(*string) error
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

// SetAccount implements [CaseResult].
func (m *CaseResultModel) SetAccount(val Reference) error {
	return store.DefaultBackedModelMutatorFunc(m, accountKey, val)
}

// GetActive returns whether the case is active.
func (m *CaseResultModel) GetActive() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *bool](m, activeKey)
}

// SetActive sets whether the case is active.
func (m *CaseResultModel) SetActive(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, activeKey, val)
}

// GetActiveAccountEscalation returns the active account escalation sys_id.
func (m *CaseResultModel) GetActiveAccountEscalation() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, activeAccountEscalationKey)
}

// SetActiveAccountEscalation sets the active account escalation sys_id.
func (m *CaseResultModel) SetActiveAccountEscalation(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, activeAccountEscalationKey, val)
}

// GetActiveEscalation returns the active escalation sys_id.
func (m *CaseResultModel) GetActiveEscalation() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, activeEscalationKey)
}

// SetActiveEscalation sets the active escalation sys_id.
func (m *CaseResultModel) SetActiveEscalation(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, activeEscalationKey, val)
}

// GetActivityDue returns the activity due date.
func (m *CaseResultModel) GetActivityDue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, activityDueKey)
}

// SetActivityDue sets the activity due date.
func (m *CaseResultModel) SetActivityDue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, activityDueKey, val)
}

// GetAdditionalAssigneeList returns the list of additional assignee sys_ids.
func (m *CaseResultModel) GetAdditionalAssigneeList() ([]*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, []*string](m, additionalAssigneeListKey)
}

// SetAdditionalAssigneeList sets the list of additional assignee sys_ids.
func (m *CaseResultModel) SetAdditionalAssigneeList(val []*string) error {
	return store.DefaultBackedModelMutatorFunc(m, additionalAssigneeListKey, val)
}

// GetApproval returns the approval type.
func (m *CaseResultModel) GetApproval() (*Approval, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *Approval](m, approvalKey)
}

// SetApproval sets the approval type.
func (m *CaseResultModel) SetApproval(val *Approval) error {
	return store.DefaultBackedModelMutatorFunc(m, approvalKey, val)
}

// GetApprovalHistory returns the approval history.
func (m *CaseResultModel) GetApprovalHistory() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, approvalHistoryKey)
}

// SetApprovalHistory sets the approval history.
func (m *CaseResultModel) SetApprovalHistory(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, approvalHistoryKey, val)
}

// GetApprovalSet returns the approval set timestamp.
func (m *CaseResultModel) GetApprovalSet() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, approvalSetKey)
}

// SetApprovalSet sets the approval set timestamp.
func (m *CaseResultModel) SetApprovalSet(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, approvalSetKey, val)
}

// GetAsset returns the asset sys_id.
func (m *CaseResultModel) GetAsset() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, assetKey)
}

// SetAsset sets the asset sys_id.
func (m *CaseResultModel) SetAsset(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, assetKey, val)
}

// SetAssignedTo implements [CaseResult].
func (m *CaseResultModel) SetAssignedTo(val Reference) error {
	return store.DefaultBackedModelMutatorFunc(m, assignedToKey, val)
}

// SetAssignmentGroup implements [CaseResult].
func (m *CaseResultModel) SetAssignmentGroup(val Reference) error {
	return store.DefaultBackedModelMutatorFunc(m, assignmentGroupKey, val)
}

// GetBusinessDuration returns the business duration.
func (m *CaseResultModel) GetBusinessDuration() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, businessDurationKey)
}

// SetBusinessDuration sets the business duration.
func (m *CaseResultModel) SetBusinessDuration(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, businessDurationKey, val)
}

// GetBusinessImpact returns the business impact.
func (m *CaseResultModel) GetBusinessImpact() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, businessImpactKey)
}

// SetBusinessImpact sets the business impact.
func (m *CaseResultModel) SetBusinessImpact(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, businessImpactKey, val)
}

// GetBusinessService returns the business service sys_id.
func (m *CaseResultModel) GetBusinessService() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, businessServiceKey)
}

// SetBusinessService sets the business service sys_id.
func (m *CaseResultModel) SetBusinessService(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, businessServiceKey, val)
}

// GetCase returns the case short description and number.
func (m *CaseResultModel) GetCase() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, caseKey)
}

// SetCase sets the case short description and number.
func (m *CaseResultModel) SetCase(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, caseKey, val)
}

// GetCaseReport returns the case report sys_id.
func (m *CaseResultModel) GetCaseReport() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, caseReportKey)
}

// SetCaseReport sets the case report sys_id.
func (m *CaseResultModel) SetCaseReport(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, caseReportKey, val)
}

// SetCategory implements [CaseResult].
func (m *CaseResultModel) SetCategory(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, categoryKey, val)
}

// GetCause returns the cause of the problem.
func (m *CaseResultModel) GetCause() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, causeKey)
}

// SetCause sets the cause of the problem.
func (m *CaseResultModel) SetCause(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, causeKey, val)
}

// GetCausedBy returns the sys_id of the causing change request.
func (m *CaseResultModel) GetCausedBy() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, causedByKey)
}

// SetCausedBy sets the sys_id of the causing change request.
func (m *CaseResultModel) SetCausedBy(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, causedByKey, val)
}

// GetChange returns the sys_id of the change request.
func (m *CaseResultModel) GetChange() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, changeKey)
}

// SetChange sets the sys_id of the change request.
func (m *CaseResultModel) SetChange(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, changeKey, val)
}

// GetChildCaseCreationProgress returns the child case creation progress flag.
func (m *CaseResultModel) GetChildCaseCreationProgress() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *bool](m, childCaseCreationProgressKey)
}

// SetChildCaseCreationProgress sets the child case creation progress flag.
func (m *CaseResultModel) SetChildCaseCreationProgress(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, childCaseCreationProgressKey, val)
}

// GetClosedAt returns the closed timestamp.
func (m *CaseResultModel) GetClosedAt() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, closedAtKey)
}

// SetClosedAt sets the closed timestamp.
func (m *CaseResultModel) SetClosedAt(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, closedAtKey, val)
}

// SetContact implements [CaseResult].
func (m *CaseResultModel) SetContact(val Reference) error {
	return store.DefaultBackedModelMutatorFunc(m, contactKey, val)
}

// SetDescription implements [CaseResult].
func (m *CaseResultModel) SetDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, descriptionKey, val)
}

// SetNumber implements [CaseResult].
func (m *CaseResultModel) SetNumber(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, numberKey, val)
}

// SetPriority implements [CaseResult].
func (m *CaseResultModel) SetPriority(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, priorityKey, val)
}

// SetShortDescription implements [CaseResult].
func (m *CaseResultModel) SetShortDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, shortDescriptionKey, val)
}

// SetState implements [CaseResult].
func (m *CaseResultModel) SetState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, stateKey, val)
}

// SetSysCreatedOn implements [CaseResult].
func (m *CaseResultModel) SetSysCreatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysCreatedOnKey, val)
}

// SetSysID implements [CaseResult].
func (m *CaseResultModel) SetSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIDKey, val)
}

// SetSysUpdatedOn implements [CaseResult].
func (m *CaseResultModel) SetSysUpdatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysUpdatedOnKey, val)
}

// CreateCaseResultFromDiscriminatorValue creates a new instance of CaseResult.
func CreateCaseResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCaseResult(), nil
}
