package core

import "testing"

func TestErrorMappingSet(t *testing.T) {
	eM := NewErrorMapping()
	eM.Set("404", "Not Found")

	if eM.Len() != 1 {
		t.Errorf("Expected length to be 1, but got %d", eM.Len())
	}

	msg, found := eM.Get(404)
	if !found || msg != "Not Found" {
		t.Errorf("Expected to find '404' with message 'Not Found', but got %v with message %s", found, msg)
	}
}

func TestErrorMappingGetExactMatch(t *testing.T) {
	eM := NewErrorMapping()
	eM.Set("404", "Not Found")

	msg, found := eM.Get(404)
	if !found || msg != "Not Found" {
		t.Errorf("Expected to find '404' with message 'Not Found', but got %v with message %s", found, msg)
	}
}

func TestErrorMappingGetRelativeMatch(t *testing.T) {
	eM := NewErrorMapping()
	eM.Set("4XX", "Client Error")

	msg, found := eM.Get(400)
	if !found || msg != "Client Error" {
		t.Errorf("Expected to find '4XX' with message 'Client Error' for 400, but got %v with message %s", found, msg)
	}

	msg, found = eM.Get(502)
	if found || msg != "" {
		t.Errorf("Expected not found and no message: got (%s)", msg)
	}

	eM.Set("5XX", "Server Error")
	msg, found = eM.Get(502)
	if !found || msg != "Server Error" {
		t.Errorf("Expected to find '5XX' with message 'Server Error' for 400, but got %v with message %s", found, msg)
	}
}

func TestErrorMappingGetNotFound(t *testing.T) {
	eM := NewErrorMapping()
	eM.Set("404", "Not Found")

	msg, found := eM.Get(500)
	if found || msg != "" {
		t.Errorf("Expected not to find '500', but got %v with message %s", found, msg)
	}
}

func TestErrorMappingGetNonErrorStatus(t *testing.T) {
	eM := NewErrorMapping()
	eM.Set("404", "Not Found")

	msg, found := eM.Get(200)
	if found || msg != "" {
		t.Errorf("Expected not to find '200', but got %v with message %s", found, msg)
	}
}
