package core

import "testing"

func TestNewFragment(t *testing.T) {
	field := "exampleField"
	operator := Is
	value := "exampleValue"

	fragment := NewFragment(field, operator, value)

	if fragment.Field != field {
		t.Errorf("Expected Field to be %s, got %s", field, fragment.Field)
	}

	if fragment.RelationalOperator != operator {
		t.Errorf("Expected RelationalOperator to be %v, got %v", operator, fragment.RelationalOperator)
	}

	if fragment.Value != value {
		t.Errorf("Expected Value to be %v, got %v", value, fragment.Value)
	}
}

func TestFragmentSetNext(t *testing.T) {
	fragment1 := NewFragment("field1", Is, "value1")
	fragment2 := NewFragment("field2", IsNot, "value2")

	fragment1.SetNext(fragment2, And)

	if fragment1.next != fragment2 {
		t.Error("Expected fragment1's next to be fragment2")
	}

	if fragment1.LogicalOperator != And {
		t.Errorf("Expected fragment1's LogicalOperator to be And, got %v", fragment1.LogicalOperator)
	}
}

func TestFragmentIterate(t *testing.T) {
	fragment1 := NewFragment("field1", Is, "value1")
	fragment2 := NewFragment("field2", IsNot, "value2")
	fragment3 := NewFragment("field3", GreaterThan, "value3")

	fragment1.SetNext(fragment2, And)
	fragment2.SetNext(fragment3, Or)

	fragments := []*Fragment{}
	fragment1.Iterate(func(f *Fragment) bool {
		fragments = append(fragments, f)
		return true
	})

	if len(fragments) != 3 {
		t.Errorf("Expected 3 fragments, got %d", len(fragments))
	}

	if fragments[0] != fragment1 || fragments[1] != fragment2 || fragments[2] != fragment3 {
		t.Error("Fragments are not in the expected order")
	}

	fragments = []*Fragment{}
	fragment1.Iterate(func(f *Fragment) bool {
		fragments = append(fragments, f)
		return false
	})

	if len(fragments) != 1 {
		t.Errorf("Expected 1 fragments, got %d", len(fragments))
	}

	if fragments[0] != fragment1 {
		t.Error("Fragments are not in the expected order")
	}
}

func TestFragmentString(t *testing.T) {
	fragment := NewFragment("exampleField", IsNot, "exampleValue")

	expected := "exampleField!=exampleValue"
	result := fragment.String()

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	fragment = NewFragment("exampleField", IsEmpty, nil)

	expected = "exampleFieldISEMPTY"
	result = fragment.String()

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
