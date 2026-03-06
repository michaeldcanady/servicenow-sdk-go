package policyapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewPolicyDefinition(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Successful",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPolicyDefinition()
			assert.NotNil(t, p)
			assert.NotNil(t, p.GetBackingStore())
		})
	}
}

func TestCreatePolicyDefinitionFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Successful",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseNode := mocking.NewMockParseNode()
			parsable, err := CreatePolicyDefinitionFromDiscriminatorValue(parseNode)
			assert.NotNil(t, parsable)
			assert.Nil(t, err)
		})
	}
}

func TestPolicyDefinition_GetFieldDeserializers(t *testing.T) {
	tests := []struct {
		name         string
		expectedKeys []string
	}{
		{
			name: "Successful",
			expectedKeys: []string{
				policyDefinitionSysIdKey,
				policyDefinitionNameKey,
				policyDefinitionDescriptionKey,
				policyDefinitionActiveKey,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPolicyDefinition()
			deserializers := p.GetFieldDeserializers()
			assert.NotNil(t, deserializers)
			for _, key := range tt.expectedKeys {
				assert.Contains(t, deserializers, key)
			}
		})
	}
}

func TestPolicyDefinition_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		setup     func(*PolicyDefinition)
		expecting func(*mocking.MockSerializationWriter)
	}{
		{
			name: "Successful",
			setup: func(p *PolicyDefinition) {
				_ = p.SetSysId(newInternal.ToPointer("sys_id"))
				_ = p.SetName(newInternal.ToPointer("name"))
				_ = p.SetDescription(newInternal.ToPointer("description"))
				_ = p.SetActive(newInternal.ToPointer(true))
			},
			expecting: func(sw *mocking.MockSerializationWriter) {
				sw.On("WriteStringValue", policyDefinitionSysIdKey, mock.Anything).Return(nil)
				sw.On("WriteStringValue", policyDefinitionNameKey, mock.Anything).Return(nil)
				sw.On("WriteStringValue", policyDefinitionDescriptionKey, mock.Anything).Return(nil)
				sw.On("WriteBoolValue", policyDefinitionActiveKey, mock.Anything).Return(nil)
			},
		},
		{
			name:  "Nil_PolicyDefinition",
			setup: nil,
			expecting: func(sw *mocking.MockSerializationWriter) {
				// No calls expected
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p *PolicyDefinition
			if tt.setup != nil {
				p = NewPolicyDefinition()
				tt.setup(p)
			}
			writer := mocking.NewMockSerializationWriter()
			tt.expecting(writer)

			err := p.Serialize(writer)
			assert.Nil(t, err)
			writer.AssertExpectations(t)
		})
	}
}

func TestPolicyDefinition_Getters(t *testing.T) {
	tests := []struct {
		name  string
		key   string
		value interface{}
		call  func(*PolicyDefinition) (interface{}, error)
	}{
		{
			name:  "GetSysId",
			key:   policyDefinitionSysIdKey,
			value: newInternal.ToPointer("sys_id"),
			call:  func(p *PolicyDefinition) (interface{}, error) { return p.GetSysId() },
		},
		{
			name:  "GetName",
			key:   policyDefinitionNameKey,
			value: newInternal.ToPointer("name"),
			call:  func(p *PolicyDefinition) (interface{}, error) { return p.GetName() },
		},
		{
			name:  "GetDescription",
			key:   policyDefinitionDescriptionKey,
			value: newInternal.ToPointer("description"),
			call:  func(p *PolicyDefinition) (interface{}, error) { return p.GetDescription() },
		},
		{
			name:  "GetActive",
			key:   policyDefinitionActiveKey,
			value: newInternal.ToPointer(true),
			call:  func(p *PolicyDefinition) (interface{}, error) { return p.GetActive() },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStore := mocking.NewMockBackingStore()
			mockStore.On("Get", tt.key).Return(tt.value, nil)

			mockModel := mocking.NewMockModel()
			mockModel.On("GetBackingStore").Return(mockStore)

			p := &PolicyDefinition{Model: mockModel}

			res, err := tt.call(p)

			assert.Nil(t, err)
			assert.Equal(t, tt.value, res)
			mockStore.AssertExpectations(t)
			mockModel.AssertExpectations(t)
		})
	}
}

func TestPolicyDefinition_Setters(t *testing.T) {
	tests := []struct {
		name  string
		key   string
		value interface{}
		call  func(*PolicyDefinition) error
	}{
		{
			name:  "SetSysId",
			key:   policyDefinitionSysIdKey,
			value: newInternal.ToPointer("sys_id"),
			call: func(p *PolicyDefinition) error {
				return p.SetSysId(newInternal.ToPointer("sys_id"))
			},
		},
		{
			name:  "SetName",
			key:   policyDefinitionNameKey,
			value: newInternal.ToPointer("name"),
			call: func(p *PolicyDefinition) error {
				return p.SetName(newInternal.ToPointer("name"))
			},
		},
		{
			name:  "SetDescription",
			key:   policyDefinitionDescriptionKey,
			value: newInternal.ToPointer("description"),
			call: func(p *PolicyDefinition) error {
				return p.SetDescription(newInternal.ToPointer("description"))
			},
		},
		{
			name:  "SetActive",
			key:   policyDefinitionActiveKey,
			value: newInternal.ToPointer(true),
			call: func(p *PolicyDefinition) error {
				return p.SetActive(newInternal.ToPointer(true))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStore := mocking.NewMockBackingStore()
			mockStore.On("Set", tt.key, tt.value).Return(nil)

			mockModel := mocking.NewMockModel()
			mockModel.On("GetBackingStore").Return(mockStore)

			p := &PolicyDefinition{Model: mockModel}

			err := tt.call(p)

			assert.Nil(t, err)
			mockStore.AssertExpectations(t)
			mockModel.AssertExpectations(t)
		})
	}
}
