package tableapi

import (
	"bytes"
	"encoding/json"
)

// TableValue2 is a struct that can represent different formats of data
type TableValue2 struct {
	Link         string     `json:"link,omitempty"`          // link is the URL of the data
	Value        *DataValue `json:"value,omitempty"`         // value is the ID of the data
	DisplayValue *DataValue `json:"display_value,omitempty"` // display_value is the name of the data
}

func (d *TableValue2) isEmpty() bool {
	return d.Link == "" && d.Value == nil && d.DisplayValue == nil
}

func (d *TableValue2) Unmarshal(data []byte) error {
	// create a decoder to decode the JSON data
	dec := json.NewDecoder(bytes.NewReader(data))

	// decode the JSON data into a token
	tok, err := dec.Token()
	if err != nil {
		return err
	}

	// check the type of the token
	switch v := tok.(type) {
	case string:
		d.Value = &DataValue{value: v}
	default:
		// Not sure how to make this error
		_ = json.Unmarshal(data, d)
	}

	return nil
}
