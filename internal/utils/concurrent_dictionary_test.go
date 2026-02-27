package utils

import (
	"testing"
)

func TestNewConcurrentDictionary(t *testing.T) {
	dict := NewConcurrentDictionary[string, int]()
	if dict == nil {
		t.Fatal("NewConcurrentDictionary returned nil")
	}
	if dict.mapping == nil {
		t.Error("mapping not initialized")
	}
}

func TestConcurrentDictionary_Get(t *testing.T) {
	d := NewConcurrentDictionary[string, int]()
	_ = d.Add("k", 1)
	var nilD *ConcurrentDictionary[string, int]

	tests := []struct {
		name     string
		dict     *ConcurrentDictionary[string, int]
		key      string
		expected int
		err      bool
	}{
		{"Exists", d, "k", 1, false},
		{"NotExists", d, "bad", 0, true},
		{"NilDict", nilD, "k", 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.dict.Get(tt.key)
			if (err != nil) != tt.err {
				t.Errorf("got err %v, expected err %v", err, tt.err)
			}
			if res != tt.expected {
				t.Errorf("got %v, expected %v", res, tt.expected)
			}
		})
	}
}

func TestConcurrentDictionary_Add(t *testing.T) {
	d := NewConcurrentDictionary[string, int]()
	var nilD *ConcurrentDictionary[string, int]

	tests := []struct {
		name string
		dict *ConcurrentDictionary[string, int]
		key  string
		val  int
		err  bool
	}{
		{"First", d, "k1", 1, false},
		{"Duplicate", d, "k1", 2, true},
		{"NilDict", nilD, "k", 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.dict.Add(tt.key, tt.val)
			if (err != nil) != tt.err {
				t.Errorf("got err %v, expected err %v", err, tt.err)
			}
		})
	}
}

func TestConcurrentDictionary_Update(t *testing.T) {
	d := NewConcurrentDictionary[string, int]()
	_ = d.Add("k", 1)
	var nilD *ConcurrentDictionary[string, int]

	tests := []struct {
		name string
		dict *ConcurrentDictionary[string, int]
		key  string
		val  int
		err  bool
	}{
		{"Ok", d, "k", 2, false},
		{"Missing", d, "bad", 3, true},
		{"NilDict", nilD, "k", 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.dict.Update(tt.key, tt.val)
			if (err != nil) != tt.err {
				t.Errorf("got err %v, expected err %v", err, tt.err)
			}
		})
	}
}

func TestConcurrentDictionary_Contains(t *testing.T) {
	d := NewConcurrentDictionary[string, int]()
	_ = d.Add("k", 1)
	var nilD *ConcurrentDictionary[string, int]

	tests := []struct {
		name     string
		dict     *ConcurrentDictionary[string, int]
		key      string
		expected bool
	}{
		{"True", d, "k", true},
		{"False", d, "bad", false},
		{"NilDict", nilD, "k", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.dict.Contains(tt.key) != tt.expected {
				t.Errorf("got %v, expected %v", tt.dict.Contains(tt.key), tt.expected)
			}
		})
	}
}

func TestConcurrentDictionary_Remove(t *testing.T) {
	d := NewConcurrentDictionary[string, int]()
	_ = d.Add("k", 1)
	var nilD *ConcurrentDictionary[string, int]

	tests := []struct {
		name string
		dict *ConcurrentDictionary[string, int]
		key  string
		err  bool
	}{
		{"Ok", d, "k", false},
		{"Missing", d, "bad", true},
		{"NilDict", nilD, "k", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.dict.Remove(tt.key)
			if (err != nil) != tt.err {
				t.Errorf("got err %v, expected err %v", err, tt.err)
			}
		})
	}
}

func TestConcurrentDictionary_Pop(t *testing.T) {
	d := NewConcurrentDictionary[string, int]()
	_ = d.Add("k", 1)
	var nilD *ConcurrentDictionary[string, int]

	tests := []struct {
		name     string
		dict     *ConcurrentDictionary[string, int]
		key      string
		expected int
		err      bool
	}{
		{"Ok", d, "k", 1, false},
		{"Missing", d, "bad", 0, true},
		{"NilDict", nilD, "k", 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.dict.Pop(tt.key)
			if (err != nil) != tt.err {
				t.Errorf("got err %v, expected err %v", err, tt.err)
			}
			if res != tt.expected {
				t.Errorf("got %v, expected %v", res, tt.expected)
			}
		})
	}
}
