package caseapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// CaseResult represents a single case object.
type CaseResult interface {
	serialization.Parsable
	kiotaStore.BackedModel
}

type CaseResultModel struct {
	newInternal.BaseModel
}

func NewCaseResult() *CaseResultModel {
	return &CaseResultModel{BaseModel: *newInternal.NewBaseModel()}
}

func (m *CaseResultModel) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	val, _ := m.GetBackingStore().Get("additionalData")
	if val != nil {
		return writer.WriteAdditionalData(val.(map[string]interface{}))
	}
	return nil
}

func (m *CaseResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		"*": func(n serialization.ParseNode) error {
			val, err := n.GetRawValue()
			if err != nil {
				return err
			}
			m.GetBackingStore().Set("additionalData", val)
			return nil
		},
	}
}

func CreateCaseResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCaseResult(), nil
}

// ActivitiesResult represents case activities.
type ActivitiesResult interface {
	serialization.Parsable
	kiotaStore.BackedModel
}

type ActivitiesResultModel struct {
	newInternal.BaseModel
}

func NewActivitiesResult() *ActivitiesResultModel {
	return &ActivitiesResultModel{BaseModel: *newInternal.NewBaseModel()}
}

func (m *ActivitiesResultModel) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	val, _ := m.GetBackingStore().Get("additionalData")
	if val != nil {
		return writer.WriteAdditionalData(val.(map[string]interface{}))
	}
	return nil
}

func (m *ActivitiesResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		"*": func(n serialization.ParseNode) error {
			val, err := n.GetRawValue()
			if err != nil {
				return err
			}
			m.GetBackingStore().Set("additionalData", val)
			return nil
		},
	}
}

func CreateActivitiesResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewActivitiesResult(), nil
}

// FieldValuesResult represents field values.
type FieldValuesResult interface {
	serialization.Parsable
	kiotaStore.BackedModel
}

type FieldValuesResultModel struct {
	newInternal.BaseModel
}

func NewFieldValuesResult() *FieldValuesResultModel {
	return &FieldValuesResultModel{BaseModel: *newInternal.NewBaseModel()}
}

func (m *FieldValuesResultModel) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	val, _ := m.GetBackingStore().Get("additionalData")
	if val != nil {
		return writer.WriteAdditionalData(val.(map[string]interface{}))
	}
	return nil
}

func (m *FieldValuesResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		"*": func(n serialization.ParseNode) error {
			val, err := n.GetRawValue()
			if err != nil {
				return err
			}
			m.GetBackingStore().Set("additionalData", val)
			return nil
		},
	}
}

func CreateFieldValuesResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewFieldValuesResult(), nil
}
