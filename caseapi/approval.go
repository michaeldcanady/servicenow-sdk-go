package caseapi

import (
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
)

const (
	approvalUnknown      = "unknown"
	approvalApproved     = "approved"
	approvalCancelled    = "cancelled"
	approvalDuplicate    = "duplicate"
	approvalNotRequired  = "not_required"
	approvalNotRequested = "not requested"
	approvalRejected     = "rejected"
	approvalRequested    = "requested"
)

// Approval specifies the UI approval for which to render the data.
type Approval int8

const (
	// ApprovalUnknown represents an unknown UI approval.
	ApprovalUnknown Approval = iota - 1
	ApprovalApproved
	ApprovalCancelled
	ApprovalDuplicate
	ApprovalNotRequired
	ApprovalNotRequested
	ApprovalRejected
	ApprovalRequested
)

// ParseApproval resolves the wire representation of a approval to a [Approval].
// Matching is case-insensitive.
func ParseApproval(s string) (interface{}, error) {
	if approval, ok := approvalValues[strings.ToLower(s)]; ok {
		return approval, nil
	}
	return ApprovalUnknown, unknownEnumValueError("approval", s)
}

var approvalStrings = map[Approval]string{
	ApprovalUnknown:      approvalUnknown,
	ApprovalApproved:     approvalApproved,
	ApprovalCancelled:    approvalCancelled,
	ApprovalDuplicate:    approvalDuplicate,
	ApprovalNotRequired:  approvalNotRequired,
	ApprovalNotRequested: approvalNotRequested,
	ApprovalRejected:     approvalRejected,
	ApprovalRequested:    approvalRequested,
}

// approvalValues is the lower-cased inverse of [approvalStrings], used by [ParseApproval].
var approvalValues = invertEnumStrings(approvalStrings, ApprovalUnknown)

// String returns the string representation of the Approval.
func (e Approval) String() string {
	return conversion.EnumString(approvalStrings, e, approvalUnknown)
}
