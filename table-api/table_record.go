package tableapi

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

type TableRecord interface {
	Get(string) (RecordElement, error)
	Set(string, RecordElement) error
	serialization.Parsable
	store.BackedModel
}

type enumerationStyle int64

const (
	enumerationStyleAll enumerationStyle = iota
	enumerationStyleOnlyChanged
	enumerationStyleOnlyChangedToNil
)

type tableRecord struct {
	backingStore     store.BackingStore
	enumerationStyle enumerationStyle
}

func NewTableRecord() TableRecord {
	return &tableRecord{
		backingStore:     store.BackingStoreFactoryInstance(),
		enumerationStyle: enumerationStyleAll,
	}
}

func CreateTableRecordFromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	raw, err := parseNode.GetRawValue()
	if err != nil {
		return nil, err
	}
	record, ok := raw.(map[string]interface{})
	if !ok {
		// TODO: define error
		return nil, nil
	}
	tableRecord := NewTableRecord()

	for key := range record {
		if err := tableRecord.Set(key, nil); err != nil {
			return nil, err
		}
	}
	return tableRecord, nil
}

func (tR *tableRecord) GetBackingStore() store.BackingStore {
	return tR.backingStore
}

// Serialize writes the objects properties to the current writer.
func (tR *tableRecord) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(tR) {
		return nil
	}

	for key, value := range tR.enumerate() {
		actualValue, err := value.value()
		if err != nil {
			return err
		}
		if err := writer.WriteAnyValue(key, actualValue); err != nil {
			return err
		}
	}
	return nil
}

func (tR *tableRecord) enumerate() map[string]RecordElement {
	enumerator := make(map[string]RecordElement, 0)

	// Helper function to enumerate with the current return mode
	enumerateWithReturnMode := func(returnOnlyChangedValues bool) {
		original := tR.GetBackingStore().GetReturnOnlyChangedValues()
		tR.GetBackingStore().SetReturnOnlyChangedValues(returnOnlyChangedValues)
		for key, value := range tR.GetBackingStore().Enumerate() {
			typedValue, _ := value.(RecordElement)
			enumerator[key] = typedValue
		}
		tR.GetBackingStore().SetReturnOnlyChangedValues(original)
	}

	switch tR.enumerationStyle {
	case enumerationStyleOnlyChanged:
		enumerateWithReturnMode(true)
	case enumerationStyleAll:
		enumerateWithReturnMode(false)
	case enumerationStyleOnlyChangedToNil:
		keys := tR.GetBackingStore().EnumerateKeysForValuesChangedToNil()
		for _, key := range keys {
			value, _ := tR.GetBackingStore().Get(key)
			typedValue, _ := value.(RecordElement)
			enumerator[key] = typedValue
		}
	}

	return enumerator
}

func (tR *tableRecord) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	fieldDeserializers := map[string]func(serialization.ParseNode) error{}

	for key := range tR.backingStore.Enumerate() {
		fieldDeserializers[key] = func(pn serialization.ParseNode) error {
			tR.GetBackingStore().SetInitializationCompleted(false)
			val, err := pn.GetRawValue()
			if err != nil {
				return nil
			}
			_, ok := val.(map[string]interface{})
			var elem RecordElement
			if !ok {
				elem = NewRecordElement()
				if err := elem.SetValue(val); err != nil {
					return err
				}
			} else {
				var ok bool
				rawElem, err := pn.GetObjectValue(CreateRecordElementFromDiscriminatorValue)
				if err != nil {
					return err
				}
				elem, ok = rawElem.(RecordElement)
				if !ok {
					// TODO: define error
					return nil
				}
			}
			if err := tR.Set(key, elem); err != nil {
				return err
			}
			tR.GetBackingStore().SetInitializationCompleted(true)
			return nil
		}
	}

	return fieldDeserializers
}

func (tR *tableRecord) Get(key string) (RecordElement, error) {
	if internal.IsNil(tR) {
		return nil, nil
	}

	val, err := tR.GetBackingStore().Get(key)
	if err != nil {
		return nil, err
	}

	elem, ok := val.(RecordElement)
	if !ok {
		return nil, errors.New("elem is not RecordElement")
	}

	return elem, nil
}

func (tR *tableRecord) Set(key string, elem RecordElement) error {
	if internal.IsNil(tR) {
		return nil
	}

	return tR.GetBackingStore().Set(key, elem)
}
