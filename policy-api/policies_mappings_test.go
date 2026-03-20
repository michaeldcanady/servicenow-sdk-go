package policyapi

import (
	"testing"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewPoliciesMapping(t *testing.T) {
	p := NewPoliciesMapping()
	assert.NotNil(t, p)
	assert.NotNil(t, p.GetBackingStore())
}

func TestCreatePoliciesMappingsInputFromDiscriminatorValue(t *testing.T) {
	parseNode := mocking.NewMockParseNode()
	parsable, err := CreatePoliciesMappingsInputFromDiscriminatorValue(parseNode)
	assert.NotNil(t, parsable)
	assert.Nil(t, err)
}

func TestPoliciesMappingsInput_Getters(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name  string
		key   string
		value interface{}
		call  func(*PoliciesMapping) (interface{}, error)
	}{
		{"GetDescription", PoliciesMappingsResolvedDescription, newInternal.ToPointer("desc"), func(p *PoliciesMapping) (interface{}, error) { return p.GetDescription() }},
		{"GetDocument", PoliciesMappingsResolvedDocument, newInternal.ToPointer("doc"), func(p *PoliciesMapping) (interface{}, error) { return p.GetDocument() }},
		{"GetDocumentRef", PoliciesMappingsResolvedDocumentRef, NewRef(), func(p *PoliciesMapping) (interface{}, error) { return p.GetDocumentRef() }},
		{"GetError", PoliciesMappingsResolvedError, newInternal.NewMainError(), func(p *PoliciesMapping) (interface{}, error) { return p.GetError() }},
		{"GetException", PoliciesMappingsResolvedException, newInternal.ToPointer("exc"), func(p *PoliciesMapping) (interface{}, error) { return p.GetException() }},
		{"GetExceptionAllowed", PoliciesMappingsResolvedExceptionAllowed, newInternal.ToPointer(true), func(p *PoliciesMapping) (interface{}, error) { return p.GetExceptionAllowed() }},
		{"GetInputStatus", PoliciesMappingsResolvedInputStatus, newInternal.ToPointer(InputStatusValid), func(p *PoliciesMapping) (interface{}, error) { return p.GetInputStatus() }},
		{"GetLastUpdatedBy", PoliciesMappingsResolvedLastUpdatedBy, NewRef(), func(p *PoliciesMapping) (interface{}, error) { return p.GetLastUpdatedBy() }},
		{"GetNumber", PoliciesMappingsResolvedNumber, newInternal.ToPointer("123"), func(p *PoliciesMapping) (interface{}, error) { return p.GetNumber() }},
		{"GetPolicy", PoliciesMappingsResolvedPolicy, NewRef(), func(p *PoliciesMapping) (interface{}, error) { return p.GetPolicy() }},
		{"GetReason", PoliciesMappingsResolvedReason, newInternal.ToPointer("reason"), func(p *PoliciesMapping) (interface{}, error) { return p.GetReason() }},
		{"GetState", PoliciesMappingsResolvedState, newInternal.ToPointer(StateActive), func(p *PoliciesMapping) (interface{}, error) { return p.GetState() }},
		{"GetSysClassName", PoliciesMappingsResolvedSysClassName, newInternal.ToPointer("class"), func(p *PoliciesMapping) (interface{}, error) { return p.GetSysClassName() }},
		{"GetSysCreatedBy", PoliciesMappingsResolvedSysCreatedBy, newInternal.ToPointer("user"), func(p *PoliciesMapping) (interface{}, error) { return p.GetSysCreatedBy() }},
		{"GetSysCreatedOn", PoliciesMappingsResolvedSysCreatedOn, &now, func(p *PoliciesMapping) (interface{}, error) { return p.GetSysCreatedOn() }},
		{"GetSysId", PoliciesMappingsResolvedSysId, newInternal.ToPointer("id"), func(p *PoliciesMapping) (interface{}, error) { return p.GetSysId() }},
		{"GetSysUpdatedBy", PoliciesMappingsResolvedSysUpdatedBy, newInternal.ToPointer("user2"), func(p *PoliciesMapping) (interface{}, error) { return p.GetSysUpdatedBy() }},
		{"GetSysUpdatedOn", PoliciesMappingsResolvedSysUpdatedOn, &now, func(p *PoliciesMapping) (interface{}, error) { return p.GetSysUpdatedOn() }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStore := mocking.NewMockBackingStore()
			mockStore.On("Get", tt.key).Return(tt.value, nil)

			mockModel := mocking.NewMockModel()
			mockModel.On("GetBackingStore").Return(mockStore)

			p := &PoliciesMapping{Model: mockModel}

			res, err := tt.call(p)

			assert.Nil(t, err)
			assert.Equal(t, tt.value, res)
			mockStore.AssertExpectations(t)
		})
	}
}

func TestPoliciesMappingsInput_Setters(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name  string
		key   string
		value interface{}
		call  func(*PoliciesMapping) error
	}{
		{"SetDescription", PoliciesMappingsResolvedDescription, newInternal.ToPointer("desc"), func(p *PoliciesMapping) error { return p.SetDescription(newInternal.ToPointer("desc")) }},
		{"SetDocument", PoliciesMappingsResolvedDocument, newInternal.ToPointer("doc"), func(p *PoliciesMapping) error { return p.SetDocument(newInternal.ToPointer("doc")) }},
		{"SetDocumentRef", PoliciesMappingsResolvedDocumentRef, NewRef(), func(p *PoliciesMapping) error { return p.SetDocumentRef(NewRef()) }},
		{"SetError", PoliciesMappingsResolvedError, newInternal.NewMainError(), func(p *PoliciesMapping) error { return p.SetError(newInternal.NewMainError()) }},
		{"SetException", PoliciesMappingsResolvedException, newInternal.ToPointer("exc"), func(p *PoliciesMapping) error { return p.SetException(newInternal.ToPointer("exc")) }},
		{"SetExceptionAllowed", PoliciesMappingsResolvedExceptionAllowed, newInternal.ToPointer(true), func(p *PoliciesMapping) error { return p.SetExceptionAllowed(newInternal.ToPointer(true)) }},
		{"SetInputStatus", PoliciesMappingsResolvedInputStatus, newInternal.ToPointer(InputStatusValid), func(p *PoliciesMapping) error { return p.SetInputStatus(newInternal.ToPointer(InputStatusValid)) }},
		{"SetLastUpdatedBy", PoliciesMappingsResolvedLastUpdatedBy, NewRef(), func(p *PoliciesMapping) error { return p.SetLastUpdatedBy(NewRef()) }},
		{"SetNumber", PoliciesMappingsResolvedNumber, newInternal.ToPointer("123"), func(p *PoliciesMapping) error { return p.SetNumber(newInternal.ToPointer("123")) }},
		{"SetPolicy", PoliciesMappingsResolvedPolicy, NewRef(), func(p *PoliciesMapping) error { return p.SetPolicy(NewRef()) }},
		{"SetReason", PoliciesMappingsResolvedReason, newInternal.ToPointer("reason"), func(p *PoliciesMapping) error { return p.SetReason(newInternal.ToPointer("reason")) }},
		{"SetState", PoliciesMappingsResolvedState, newInternal.ToPointer(StateActive), func(p *PoliciesMapping) error { return p.SetState(newInternal.ToPointer(StateActive)) }},
		{"SetSysClassName", PoliciesMappingsResolvedSysClassName, newInternal.ToPointer("class"), func(p *PoliciesMapping) error { return p.SetSysClassName(newInternal.ToPointer("class")) }},
		{"SetSysCreatedBy", PoliciesMappingsResolvedSysCreatedBy, newInternal.ToPointer("user"), func(p *PoliciesMapping) error { return p.SetSysCreatedBy(newInternal.ToPointer("user")) }},
		{"SetSysCreatedOn", PoliciesMappingsResolvedSysCreatedOn, &now, func(p *PoliciesMapping) error { return p.SetSysCreatedOn(&now) }},
		{"SetSysId", PoliciesMappingsResolvedSysId, newInternal.ToPointer("id"), func(p *PoliciesMapping) error { return p.SetSysId(newInternal.ToPointer("id")) }},
		{"SetSysUpdatedBy", PoliciesMappingsResolvedSysUpdatedBy, newInternal.ToPointer("user2"), func(p *PoliciesMapping) error { return p.SetSysUpdatedBy(newInternal.ToPointer("user2")) }},
		{"SetSysUpdatedOn", PoliciesMappingsResolvedSysUpdatedOn, &now, func(p *PoliciesMapping) error { return p.SetSysUpdatedOn(&now) }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStore := mocking.NewMockBackingStore()
			mockStore.On("Set", tt.key, mock.Anything).Return(nil)

			mockModel := mocking.NewMockModel()
			mockModel.On("GetBackingStore").Return(mockStore)

			p := &PoliciesMapping{Model: mockModel}

			err := tt.call(p)

			assert.Nil(t, err)
			mockStore.AssertExpectations(t)
		})
	}
}

func TestPoliciesMappingsInput_Serialize(t *testing.T) {
	p := NewPoliciesMapping()
	_ = p.SetDescription(newInternal.ToPointer("desc"))
	_ = p.SetInputStatus(newInternal.ToPointer(InputStatusValid))
	_ = p.SetState(newInternal.ToPointer(StateActive))
	_ = p.SetExceptionAllowed(newInternal.ToPointer(true))
	_ = p.SetSysCreatedOn(&time.Time{})
	_ = p.SetSysUpdatedOn(&time.Time{})
	_ = p.SetDocumentRef(NewRef())
	_ = p.SetError(newInternal.NewMainError())
	_ = p.SetLastUpdatedBy(NewRef())
	_ = p.SetPolicy(NewRef())

	writer := mocking.NewMockSerializationWriter()
	writer.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
	writer.On("WriteBoolValue", mock.Anything, mock.Anything).Return(nil)
	writer.On("WriteTimeValue", mock.Anything, mock.Anything).Return(nil)
	writer.On("WriteObjectValue", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	err := p.Serialize(writer)
	assert.Nil(t, err)
	writer.AssertExpectations(t)

	var nilP *PoliciesMapping
	err = nilP.Serialize(writer)
	assert.Nil(t, err)
}

func TestPoliciesMappingsInput_GetFieldDeserializers(t *testing.T) {
	p := NewPoliciesMapping()
	deser := p.GetFieldDeserializers()
	assert.NotNil(t, deser)
	assert.Equal(t, 18, len(deser))
}
