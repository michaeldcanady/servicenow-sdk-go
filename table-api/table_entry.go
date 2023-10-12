package tableapi

import "reflect"

type TableEntry map[string]interface{}

// GetStringPtr retrieves a pointer to a string value from the TableEntry based on the given key.
func (te TableEntry) GetStringPtr(key string) *string {
    if value, ok := te[key].(string); ok {
        return &value
    }
    return nil
}

// GetIntPtr retrieves a pointer to an int value from the TableEntry based on the given key.
func (te TableEntry) GetIntPtr(key string) *int {
    if value, ok := te[key].(int); ok {
        return &value
    }
    return nil
}

// GetBoolPtr retrieves a pointer to a bool value from the TableEntry based on the given key.
func (te TableEntry) GetBoolPtr(key string) *bool {
    if value, ok := te[key].(bool); ok {
        return &value
    }
    return nil
}

// GetType returns the actual data type of the value associated with the given key.
func (te TableEntry) GetType(key string) reflect.Type {
    if value, ok := te[key]; ok {
        return reflect.TypeOf(value)
    }
    return nil
}
