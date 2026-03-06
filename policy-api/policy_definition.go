package policyapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	policyDefinitionSysIdKey       = "sys_id"
	policyDefinitionNameKey        = "name"
	policyDefinitionDescriptionKey = "description"
	policyDefinitionActiveKey      = "active"
)

// PolicyDefinition represents a Service-Now Policy Definition.
// It contains metadata about a specific policy available in the system.
type PolicyDefinition struct {
	newInternal.Model
}

// NewPolicyDefinition creates a new instance of PolicyDefinition.
func NewPolicyDefinition() *PolicyDefinition {
	return &PolicyDefinition{
		Model: newInternal.NewBaseModel(),
	}
}

// CreatePolicyDefinitionFromDiscriminatorValue creates a new PolicyDefinition from a ParseNode.
func CreatePolicyDefinitionFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewPolicyDefinition(), nil
}

// GetFieldDeserializers returns the deserialization information for this object.
func (p *PolicyDefinition) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		policyDefinitionSysIdKey:       internalSerialization.DeserializeStringFunc(p.SetSysId),
		policyDefinitionNameKey:        internalSerialization.DeserializeStringFunc(p.SetName),
		policyDefinitionDescriptionKey: internalSerialization.DeserializeStringFunc(p.SetDescription),
		policyDefinitionActiveKey:      internalSerialization.DeserializeBoolFunc(p.SetActive),
	}
}

// Serialize writes the objects properties to the current writer.
func (p *PolicyDefinition) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(p) {
		return nil
	}

	{
		val, err := p.GetSysId()
		if err != nil {
			return err
		}
		if val != nil {
			err = writer.WriteStringValue(policyDefinitionSysIdKey, val)
			if err != nil {
				return err
			}
		}
	}
	{
		val, err := p.GetName()
		if err != nil {
			return err
		}
		if val != nil {
			err = writer.WriteStringValue(policyDefinitionNameKey, val)
			if err != nil {
				return err
			}
		}
	}
	{
		val, err := p.GetDescription()
		if err != nil {
			return err
		}
		if val != nil {
			err = writer.WriteStringValue(policyDefinitionDescriptionKey, val)
			if err != nil {
				return err
			}
		}
	}
	{
		val, err := p.GetActive()
		if err != nil {
			return err
		}
		if val != nil {
			err = writer.WriteBoolValue(policyDefinitionActiveKey, val)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// GetSysId returns the unique identifier (sys_id) of the policy definition.
func (p *PolicyDefinition) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), policyDefinitionSysIdKey)
}

// SetSysId sets the unique identifier (sys_id) of the policy definition.
func (p *PolicyDefinition) SetSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), policyDefinitionSysIdKey, val)
}

// GetName returns the display name of the policy definition.
func (p *PolicyDefinition) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), policyDefinitionNameKey)
}

// SetName sets the display name of the policy definition.
func (p *PolicyDefinition) SetName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), policyDefinitionNameKey, val)
}

// GetDescription returns the functional description of what the policy does.
func (p *PolicyDefinition) GetDescription() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), policyDefinitionDescriptionKey)
}

// SetDescription sets the functional description of what the policy does.
func (p *PolicyDefinition) SetDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), policyDefinitionDescriptionKey, val)
}

// GetActive returns whether the policy definition is currently active and available for use.
func (p *PolicyDefinition) GetActive() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](p.GetBackingStore(), policyDefinitionActiveKey)
}

// SetActive sets whether the policy definition is currently active and available for use.
func (p *PolicyDefinition) SetActive(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), policyDefinitionActiveKey, val)
}
