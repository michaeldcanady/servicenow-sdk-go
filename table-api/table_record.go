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

type tableRecord struct {
	backingStore store.BackingStore
}

func NewTableRecord() TableRecord {
	return &tableRecord{
		backingStore: store.BackingStoreFactoryInstance(),
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

	for key, _ := range record {
		tableRecord.Set(key, nil)
	}
	return tableRecord, nil
}

func (tR *tableRecord) GetBackingStore() store.BackingStore {
	return tR.backingStore
}

// Serialize writes the objects properties to the current writer.
func (tE *tableRecord) Serialize(writer serialization.SerializationWriter) error {
	return nil
}

func (tR *tableRecord) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	fieldDeserializers := map[string]func(serialization.ParseNode) error{}

	for key, _ := range tR.backingStore.Enumerate() {
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
