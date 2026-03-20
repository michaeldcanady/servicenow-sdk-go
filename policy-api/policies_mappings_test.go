package policyapi

import (
	"testing"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
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
		{"GetDescription", PoliciesMappingsResolvedDescription, utils.ToPointer("desc"), func(p *PoliciesMapping) (interface{}, error) { return p.GetDescription() }},
		{"GetDocument", PoliciesMappingsResolvedDocument, utils.ToPointer("doc"), func(p *PoliciesMapping) (interface{}, error) { return p.GetDocument() }},
		{"GetDocumentRef", PoliciesMappingsResolvedDocumentRef, NewRef(), func(p *PoliciesMapping) (interface{}, error) { return p.GetDocumentRef() }},
		{"GetError", PoliciesMappingsResolvedError, model.NewMainError(), func(p *PoliciesMapping) (interface{}, error) { return p.GetError() }},
		{"GetException", PoliciesMappingsResolvedException, utils.ToPointer("exc"), func(p *PoliciesMapping) (interface{}, error) { return p.GetException() }},
		{"GetExceptionAllowed", PoliciesMappingsResolvedExceptionAllowed, utils.ToPointer(true), func(p *PoliciesMapping) (interface{}, error) { return p.GetExceptionAllowed() }},
		{"GetInputStatus", PoliciesMappingsResolvedInputStatus, utils.ToPointer(InputStatusValid), func(p *PoliciesMapping) (interface{}, error) { return p.GetInputStatus() }},
		{"GetLastUpdatedBy", PoliciesMappingsResolvedLastUpdatedBy, NewRef(), func(p *PoliciesMapping) (interface{}, error) { return p.GetLastUpdatedBy() }},
		{"GetNumber", PoliciesMappingsResolvedNumber, utils.ToPointer("123"), func(p *PoliciesMapping) (interface{}, error) { return p.GetNumber() }},
		{"GetPolicy", PoliciesMappingsResolvedPolicy, NewRef(), func(p *PoliciesMapping) (interface{}, error) { return p.GetPolicy() }},
		{"GetReason", PoliciesMappingsResolvedReason, utils.ToPointer("reason"), func(p *PoliciesMapping) (interface{}, error) { return p.GetReason() }},
		{"GetState", PoliciesMappingsResolvedState, utils.ToPointer(StateActive), func(p *PoliciesMapping) (interface{}, error) { return p.GetState() }},
		{"GetSysClassName", PoliciesMappingsResolvedSysClassName, utils.ToPointer("class"), func(p *PoliciesMapping) (interface{}, error) { return p.GetSysClassName() }},
		{"GetSysCreatedBy", PoliciesMappingsResolvedSysCreatedBy, utils.ToPointer("user"), func(p *PoliciesMapping) (interface{}, error) { return p.GetSysCreatedBy() }},
		{"GetSysCreatedOn", PoliciesMappingsResolvedSysCreatedOn, &now, func(p *PoliciesMapping) (interface{}, error) { return p.GetSysCreatedOn() }},
		{"GetSysId", PoliciesMappingsResolvedSysId, utils.ToPointer("id"), func(p *PoliciesMapping) (interface{}, error) { return p.GetSysId() }},
		{"GetSysUpdatedBy", PoliciesMappingsResolvedSysUpdatedBy, utils.ToPointer("user2"), func(p *PoliciesMapping) (interface{}, error) { return p.GetSysUpdatedBy() }},
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
		{"SetDescription", PoliciesMappingsResolvedDescription, utils.ToPointer("desc"), func(p *PoliciesMapping) error { return p.SetDescription(utils.ToPointer("desc")) }},
		{"SetDocument", PoliciesMappingsResolvedDocument, utils.ToPointer("doc"), func(p *PoliciesMapping) error { return p.SetDocument(utils.ToPointer("doc")) }},
		{"SetDocumentRef", PoliciesMappingsResolvedDocumentRef, NewRef(), func(p *PoliciesMapping) error { return p.SetDocumentRef(NewRef()) }},
		{"SetError", PoliciesMappingsResolvedError, model.NewMainError(), func(p *PoliciesMapping) error { return p.SetError(model.NewMainError()) }},
		{"SetException", PoliciesMappingsResolvedException, utils.ToPointer("exc"), func(p *PoliciesMapping) error { return p.SetException(utils.ToPointer("exc")) }},
		{"SetExceptionAllowed", PoliciesMappingsResolvedExceptionAllowed, utils.ToPointer(true), func(p *PoliciesMapping) error { return p.SetExceptionAllowed(utils.ToPointer(true)) }},
		{"SetInputStatus", PoliciesMappingsResolvedInputStatus, utils.ToPointer(InputStatusValid), func(p *PoliciesMapping) error { return p.SetInputStatus(utils.ToPointer(InputStatusValid)) }},
		{"SetLastUpdatedBy", PoliciesMappingsResolvedLastUpdatedBy, NewRef(), func(p *PoliciesMapping) error { return p.SetLastUpdatedBy(NewRef()) }},
		{"SetNumber", PoliciesMappingsResolvedNumber, utils.ToPointer("123"), func(p *PoliciesMapping) error { return p.SetNumber(utils.ToPointer("123")) }},
		{"SetPolicy", PoliciesMappingsResolvedPolicy, NewRef(), func(p *PoliciesMapping) error { return p.SetPolicy(NewRef()) }},
		{"SetReason", PoliciesMappingsResolvedReason, utils.ToPointer("reason"), func(p *PoliciesMapping) error { return p.SetReason(utils.ToPointer("reason")) }},
		{"SetState", PoliciesMappingsResolvedState, utils.ToPointer(StateActive), func(p *PoliciesMapping) error { return p.SetState(utils.ToPointer(StateActive)) }},
		{"SetSysClassName", PoliciesMappingsResolvedSysClassName, utils.ToPointer("class"), func(p *PoliciesMapping) error { return p.SetSysClassName(utils.ToPointer("class")) }},
		{"SetSysCreatedBy", PoliciesMappingsResolvedSysCreatedBy, utils.ToPointer("user"), func(p *PoliciesMapping) error { return p.SetSysCreatedBy(utils.ToPointer("user")) }},
		{"SetSysCreatedOn", PoliciesMappingsResolvedSysCreatedOn, &now, func(p *PoliciesMapping) error { return p.SetSysCreatedOn(&now) }},
		{"SetSysId", PoliciesMappingsResolvedSysId, utils.ToPointer("id"), func(p *PoliciesMapping) error { return p.SetSysId(utils.ToPointer("id")) }},
		{"SetSysUpdatedBy", PoliciesMappingsResolvedSysUpdatedBy, utils.ToPointer("user2"), func(p *PoliciesMapping) error { return p.SetSysUpdatedBy(utils.ToPointer("user2")) }},
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
	_ = p.SetDescription(utils.ToPointer("desc"))
	_ = p.SetInputStatus(utils.ToPointer(InputStatusValid))
	_ = p.SetState(utils.ToPointer(StateActive))
	_ = p.SetExceptionAllowed(utils.ToPointer(true))
	_ = p.SetSysCreatedOn(&time.Time{})
	_ = p.SetSysUpdatedOn(&time.Time{})
	_ = p.SetDocumentRef(NewRef())
	_ = p.SetError(model.NewMainError())
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
