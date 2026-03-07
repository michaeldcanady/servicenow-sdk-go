package policyapi

import (
	"testing"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewPoliciesMappingsInput(t *testing.T) {
	p := NewPoliciesMappingsInput()
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
		call  func(*PoliciesMappingsInput) (interface{}, error)
	}{
		{"GetDescription", PoliciesMappingsInputsResolvedDescription, newInternal.ToPointer("desc"), func(p *PoliciesMappingsInput) (interface{}, error) { return p.GetDescription() }},
		{"GetDocument", PoliciesMappingsInputsResolvedDocument, newInternal.ToPointer("doc"), func(p *PoliciesMappingsInput) (interface{}, error) { return p.GetDocument() }},
		{"GetDocumentRef", PoliciesMappingsInputsResolvedDocumentRef, NewRef(), func(p *PoliciesMappingsInput) (interface{}, error) { return p.GetDocumentRef() }},
		{"GetError", PoliciesMappingsInputsResolvedError, newInternal.NewMainError(), func(p *PoliciesMappingsInput) (interface{}, error) { return p.GetError() }},
		{"GetException", PoliciesMappingsInputsResolvedException, newInternal.ToPointer("exc"), func(p *PoliciesMappingsInput) (interface{}, error) { return p.GetException() }},
		{"GetExceptionAllowed", PoliciesMappingsInputsResolvedExceptionAllowed, newInternal.ToPointer(true), func(p *PoliciesMappingsInput) (interface{}, error) { return p.GetExceptionAllowed() }},
		{"GetInputStatus", PoliciesMappingsInputsResolvedInputStatus, newInternal.ToPointer(InputStatusValid), func(p *PoliciesMappingsInput) (interface{}, error) { return p.GetInputStatus() }},
		{"GetLastUpdatedBy", PoliciesMappingsInputsResolvedLastUpdatedBy, NewRef(), func(p *PoliciesMappingsInput) (interface{}, error) { return p.GetLastUpdatedBy() }},
		{"GetNumber", PoliciesMappingsInputsResolvedNumber, newInternal.ToPointer("123"), func(p *PoliciesMappingsInput) (interface{}, error) { return p.GetNumber() }},
		{"GetPolicy", PoliciesMappingsInputsResolvedPolicy, NewRef(), func(p *PoliciesMappingsInput) (interface{}, error) { return p.GetPolicy() }},
		{"GetReason", PoliciesMappingsInputsResolvedReason, newInternal.ToPointer("reason"), func(p *PoliciesMappingsInput) (interface{}, error) { return p.GetReason() }},
		{"GetState", PoliciesMappingsInputsResolvedState, newInternal.ToPointer(StateActive), func(p *PoliciesMappingsInput) (interface{}, error) { return p.GetState() }},
		{"GetSysClassName", PoliciesMappingsInputsResolvedSysClassName, newInternal.ToPointer("class"), func(p *PoliciesMappingsInput) (interface{}, error) { return p.GetSysClassName() }},
		{"GetSysCreatedBy", PoliciesMappingsInputsResolvedSysCreatedBy, newInternal.ToPointer("user"), func(p *PoliciesMappingsInput) (interface{}, error) { return p.GetSysCreatedBy() }},
		{"GetSysCreatedOn", PoliciesMappingsInputsResolvedSysCreatedOn, &now, func(p *PoliciesMappingsInput) (interface{}, error) { return p.GetSysCreatedOn() }},
		{"GetSysId", PoliciesMappingsInputsResolvedSysId, newInternal.ToPointer("id"), func(p *PoliciesMappingsInput) (interface{}, error) { return p.GetSysId() }},
		{"GetSysUpdatedBy", PoliciesMappingsInputsResolvedSysUpdatedBy, newInternal.ToPointer("user2"), func(p *PoliciesMappingsInput) (interface{}, error) { return p.GetSysUpdatedBy() }},
		{"GetSysUpdatedOn", PoliciesMappingsInputsResolvedSysUpdatedOn, &now, func(p *PoliciesMappingsInput) (interface{}, error) { return p.GetSysUpdatedOn() }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStore := mocking.NewMockBackingStore()
			mockStore.On("Get", tt.key).Return(tt.value, nil)

			mockModel := mocking.NewMockModel()
			mockModel.On("GetBackingStore").Return(mockStore)

			p := &PoliciesMappingsInput{Model: mockModel}

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
		call  func(*PoliciesMappingsInput) error
	}{
		{"SetDescription", PoliciesMappingsInputsResolvedDescription, newInternal.ToPointer("desc"), func(p *PoliciesMappingsInput) error { return p.SetDescription(newInternal.ToPointer("desc")) }},
		{"SetDocument", PoliciesMappingsInputsResolvedDocument, newInternal.ToPointer("doc"), func(p *PoliciesMappingsInput) error { return p.SetDocument(newInternal.ToPointer("doc")) }},
		{"SetDocumentRef", PoliciesMappingsInputsResolvedDocumentRef, NewRef(), func(p *PoliciesMappingsInput) error { return p.SetDocumentRef(NewRef()) }},
		{"SetError", PoliciesMappingsInputsResolvedError, newInternal.NewMainError(), func(p *PoliciesMappingsInput) error { return p.SetError(newInternal.NewMainError()) }},
		{"SetException", PoliciesMappingsInputsResolvedException, newInternal.ToPointer("exc"), func(p *PoliciesMappingsInput) error { return p.SetException(newInternal.ToPointer("exc")) }},
		{"SetExceptionAllowed", PoliciesMappingsInputsResolvedExceptionAllowed, newInternal.ToPointer(true), func(p *PoliciesMappingsInput) error { return p.SetExceptionAllowed(newInternal.ToPointer(true)) }},
		{"SetInputStatus", PoliciesMappingsInputsResolvedInputStatus, newInternal.ToPointer(InputStatusValid), func(p *PoliciesMappingsInput) error { return p.SetInputStatus(newInternal.ToPointer(InputStatusValid)) }},
		{"SetLastUpdatedBy", PoliciesMappingsInputsResolvedLastUpdatedBy, NewRef(), func(p *PoliciesMappingsInput) error { return p.SetLastUpdatedBy(NewRef()) }},
		{"SetNumber", PoliciesMappingsInputsResolvedNumber, newInternal.ToPointer("123"), func(p *PoliciesMappingsInput) error { return p.SetNumber(newInternal.ToPointer("123")) }},
		{"SetPolicy", PoliciesMappingsInputsResolvedPolicy, NewRef(), func(p *PoliciesMappingsInput) error { return p.SetPolicy(NewRef()) }},
		{"SetReason", PoliciesMappingsInputsResolvedReason, newInternal.ToPointer("reason"), func(p *PoliciesMappingsInput) error { return p.SetReason(newInternal.ToPointer("reason")) }},
		{"SetState", PoliciesMappingsInputsResolvedState, newInternal.ToPointer(StateActive), func(p *PoliciesMappingsInput) error { return p.SetState(newInternal.ToPointer(StateActive)) }},
		{"SetSysClassName", PoliciesMappingsInputsResolvedSysClassName, newInternal.ToPointer("class"), func(p *PoliciesMappingsInput) error { return p.SetSysClassName(newInternal.ToPointer("class")) }},
		{"SetSysCreatedBy", PoliciesMappingsInputsResolvedSysCreatedBy, newInternal.ToPointer("user"), func(p *PoliciesMappingsInput) error { return p.SetSysCreatedBy(newInternal.ToPointer("user")) }},
		{"SetSysCreatedOn", PoliciesMappingsInputsResolvedSysCreatedOn, &now, func(p *PoliciesMappingsInput) error { return p.SetSysCreatedOn(&now) }},
		{"SetSysId", PoliciesMappingsInputsResolvedSysId, newInternal.ToPointer("id"), func(p *PoliciesMappingsInput) error { return p.SetSysId(newInternal.ToPointer("id")) }},
		{"SetSysUpdatedBy", PoliciesMappingsInputsResolvedSysUpdatedBy, newInternal.ToPointer("user2"), func(p *PoliciesMappingsInput) error { return p.SetSysUpdatedBy(newInternal.ToPointer("user2")) }},
		{"SetSysUpdatedOn", PoliciesMappingsInputsResolvedSysUpdatedOn, &now, func(p *PoliciesMappingsInput) error { return p.SetSysUpdatedOn(&now) }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStore := mocking.NewMockBackingStore()
			mockStore.On("Set", tt.key, mock.Anything).Return(nil)

			mockModel := mocking.NewMockModel()
			mockModel.On("GetBackingStore").Return(mockStore)

			p := &PoliciesMappingsInput{Model: mockModel}

			err := tt.call(p)

			assert.Nil(t, err)
			mockStore.AssertExpectations(t)
		})
	}
}

func TestPoliciesMappingsInput_Serialize(t *testing.T) {
	p := NewPoliciesMappingsInput()
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

	var nilP *PoliciesMappingsInput
	err = nilP.Serialize(writer)
	assert.Nil(t, err)
}

func TestPoliciesMappingsInput_GetFieldDeserializers(t *testing.T) {
	p := NewPoliciesMappingsInput()
	deser := p.GetFieldDeserializers()
	assert.NotNil(t, deser)
	assert.Equal(t, 18, len(deser))
}
