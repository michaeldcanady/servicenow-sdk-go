package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewConcurrentDictionary tests the creation of a new ConcurrentDictionary.
func TestNewConcurrentDictionary(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful Initialization",
			test: func(t *testing.T) {
				dict := NewConcurrentDictionary[string, int]()
				assert.NotNil(t, dict, "NewConcurrentDictionary should not return nil")
				assert.Empty(t, dict.mapping, "NewConcurrentDictionary should initialize with an empty map")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestConcurrentDictionary_Get(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Existing Key",
			test: func(t *testing.T) {
				dict := NewConcurrentDictionary[string, int]()
				dict.mapping["key"] = 42

				value, err := dict.Get("key")
				assert.NoError(t, err, "Expected no error when retrieving an existing key")
				assert.Equal(t, 42, value, "Expected value 42 for key 'key'")
			},
		},
		{
			name: "Non-Existing Key",
			test: func(t *testing.T) {
				dict := NewConcurrentDictionary[string, int]()

				value, err := dict.Get("missing")
				assert.Error(t, err, "Expected an error when retrieving a non-existing key")
				assert.Zero(t, value, "Expected zero value for a non-existing key")
			},
		},
		{
			name: "Nil type",
			test: func(t *testing.T) {
				dict := (*ConcurrentDictionary[string, int])(nil)

				value, err := dict.Get("missing")
				assert.Equal(t, 0, value)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestConcurrentDictionary_Add(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Add New Key",
			test: func(t *testing.T) {
				dict := NewConcurrentDictionary[string, int]()
				err := dict.Add("key", 42)
				assert.NoError(t, err, "Expected no error when adding a new key")
				assert.Equal(t, 42, dict.mapping["key"], "Expected value 42 for key 'key'")
			},
		},
		{
			name: "Add Existing Key",
			test: func(t *testing.T) {
				dict := NewConcurrentDictionary[string, int]()
				dict.mapping["key"] = 42
				err := dict.Add("key", 100)
				assert.Error(t, err, "Expected an error when adding an existing key")
			},
		},
		{
			name: "Nil type",
			test: func(t *testing.T) {
				dict := (*ConcurrentDictionary[string, int])(nil)

				err := dict.Add("missing", 100)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestConcurrentDictionary_Update(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Update Existing Key",
			test: func(t *testing.T) {
				dict := NewConcurrentDictionary[string, int]()
				dict.mapping["key"] = 42
				err := dict.Update("key", 100)
				assert.NoError(t, err, "Expected no error when updating an existing key")
				assert.Equal(t, 100, dict.mapping["key"], "Expected updated value 100 for key 'key'")
			},
		},
		{
			name: "Update Non-Existing Key",
			test: func(t *testing.T) {
				dict := NewConcurrentDictionary[string, int]()
				err := dict.Update("key", 42)
				assert.Error(t, err, "Expected an error when updating a non-existing key")
			},
		},
		{
			name: "Nil type",
			test: func(t *testing.T) {
				dict := (*ConcurrentDictionary[string, int])(nil)

				err := dict.Update("missing", 100)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestConcurrentDictionary_Contains(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Key Exists",
			test: func(t *testing.T) {
				dict := NewConcurrentDictionary[string, int]()
				dict.mapping["key"] = 42

				result := dict.Contains("key")
				assert.True(t, result, "Expected Contains to return true for an existing key")
			},
		},
		{
			name: "Key Does Not Exist",
			test: func(t *testing.T) {
				dict := NewConcurrentDictionary[string, int]()

				result := dict.Contains("missing")
				assert.False(t, result, "Expected Contains to return false for a non-existing key")
			},
		},
		{
			name: "Empty Dictionary",
			test: func(t *testing.T) {
				dict := NewConcurrentDictionary[string, int]()

				result := dict.Contains("key")
				assert.False(t, result, "Expected Contains to return false for an empty dictionary")
			},
		},
		{
			name: "Nil type",
			test: func(t *testing.T) {
				dict := (*ConcurrentDictionary[string, int])(nil)

				contains := dict.Contains("missing")
				assert.False(t, contains)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestConcurrentDictionary_Remove(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Remove Existing Key",
			test: func(t *testing.T) {
				dict := NewConcurrentDictionary[string, int]()
				dict.mapping["key"] = 42
				err := dict.Remove("key")
				assert.NoError(t, err, "Expected no error when removing an existing key")
				_, exists := dict.mapping["key"]
				assert.False(t, exists, "Expected key to be removed from the dictionary")
			},
		},
		{
			name: "Remove Non-Existing Key",
			test: func(t *testing.T) {
				dict := NewConcurrentDictionary[string, int]()
				err := dict.Remove("missing")
				assert.Error(t, err, "Expected an error when removing a non-existing key")
			},
		},
		{
			name: "Nil type",
			test: func(t *testing.T) {
				dict := (*ConcurrentDictionary[string, int])(nil)

				err := dict.Remove("missing")
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestConcurrentDictionary_Pop(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Pop Existing Key",
			test: func(t *testing.T) {
				dict := NewConcurrentDictionary[string, int]()
				dict.mapping["key"] = 42
				value, err := dict.Pop("key")
				assert.NoError(t, err, "Expected no error when popping an existing key")
				assert.Equal(t, 42, value, "Expected value 42 for key 'key'")
				_, exists := dict.mapping["key"]
				assert.False(t, exists, "Expected key to be removed after Pop")
			},
		},
		{
			name: "Pop Non-Existing Key",
			test: func(t *testing.T) {
				dict := NewConcurrentDictionary[string, int]()
				value, err := dict.Pop("missing")
				assert.Error(t, err, "Expected an error when popping a non-existing key")
				assert.Zero(t, value, "Expected zero value for a non-existing key")
			},
		},
		{
			name: "Nil type",
			test: func(t *testing.T) {
				dict := (*ConcurrentDictionary[string, int])(nil)

				value, err := dict.Pop("missing")
				assert.Nil(t, err)
				assert.Zero(t, value, "Expected zero value for null pointer")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
