package serialization

import "encoding/json"

// JSONUnmarshaller for JSON unmarshalling
type JSONUnmarshaller struct{}

func (u *JSONUnmarshaller) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
