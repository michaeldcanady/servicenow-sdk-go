package core

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yosida95/uritemplate/v3"
)

func TestToQueryMap(t *testing.T) {

	inputs := []struct {
		Input       interface{}
		ShouldError bool
		Expected    map[string]string
		CheckErr    func(error) bool
	}{
		{
			Input: struct {
				Param1 string `query:"param_1"`
				Param2 int    `query:"param_2"`
				Param3 bool   `query:"param_3"`
			}{
				Param1: "value1",
				Param2: 5,
				Param3: true,
			},
			ShouldError: false,
			Expected:    map[string]string{"param_1": "value1", "param_2": "5", "param_3": "1"},
			CheckErr:    nil,
		},
	}

	for _, input := range inputs {

		paramMap, err := ToQueryMap(input.Input)

		if input.ShouldError {
			if !input.CheckErr(err) {
				t.Errorf("Expected error, got nil")
			}
		} else if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		assert.Equal(t, input.Expected, paramMap)

	}
}

func TestNormalizeVarNames(t *testing.T) {

	input := []string{"VaR1", "vAr2", "VAr3", "vAR4", "Var5", "vaR6", "VAR7"}
	expected := map[string]string{"var1": "VaR1", "var2": "vAr2", "var3": "VAr3", "var4": "vAR4", "var5": "Var5", "var6": "vaR6", "var7": "VAR7"}

	actual := normalizeVarNames(input)

	assert.Equal(t, expected, actual)
}

func TestAddParametersWithOriginalNames(t *testing.T) {

	params := map[string]string{"var1": "val1", "var2": "val2"}
	normalizedNames := map[string]string{"var1": "VaR1", "var2": "vAr2"}

	values := addParametersWithOriginalNames(params, normalizedNames, nil)
	expected := uritemplate.Values{"VaR1": uritemplate.String("val1"), "vAr2": uritemplate.String("val2")}

	assert.Equal(t, expected, values)
}

func pick[M ~map[K]V, K comparable, V any](m M) K {

	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	index := rand.Intn(len(keys)) // generate a random index
	return keys[index]
}

func TestGetKeyWithOriginalName(t *testing.T) {

	normalizedNames := map[string]string{"var1": "VaR1", "var2": "vAr2", "var3": "VAr3", "var4": "vAR4", "var5": "Var5", "var6": "vaR6", "var7": "VAR7"}

	randomKey := pick(normalizedNames)

	value := getKeyWithOriginalName(randomKey, normalizedNames)

	assert.Equal(t, normalizedNames[randomKey], value)
}

func TestIsPointer(t *testing.T) {

	s := "test"
	i := 42

	f := func() {}

	inputs := []struct {
		Input    interface{}
		Expected bool
	}{
		{
			Input:    &s,
			Expected: true,
		},
		{
			Input:    s, // this is a string value, not a pointer
			Expected: false,
		},
		{
			Input:    i, // this is an int value, not a pointer
			Expected: false,
		},
		{
			Input:    &i, // this is a pointer to an int value
			Expected: true,
		},
		{
			Input:    nil, // this is a nil value, not a pointer
			Expected: false,
		},
		{
			Input:    (*int)(nil), // this is a nil pointer to an int type
			Expected: true,
		},
		{
			Input:    []int{1, 2, 3}, // this is a slice value, not a pointer
			Expected: false,
		},
		{
			Input:    &[3]int{1, 2, 3}, // this is a pointer to an array value
			Expected: true,
		},
		{
			Input:    map[string]int{"a": 1, "b": 2}, // this is a map value, not a pointer
			Expected: false,
		},
		{
			Input:    &map[string]int{"a": 1, "b": 2}, // this is a pointer to a map value
			Expected: true,
		},
		{
			Input:    f, // this is a function value, not a pointer
			Expected: false,
		},
		{
			Input:    &f, // this is a pointer to a function value
			Expected: true,
		},
		{
			Input:    struct{}{}, // this is a struct value, not a pointer
			Expected: false,
		},
		{
			Input:    &struct{}{}, // this is a pointer to a struct value
			Expected: true,
		},
	}

	for _, input := range inputs {
		actual := IsPointer(input.Input)
		assert.Equal(t, input.Expected, actual)
	}
}
