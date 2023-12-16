package core

import (
	"testing"
)

func TestHttpMethod_String(t *testing.T) {
	tests := []Test{
		{
			Title:    "Test GET",
			Expected: "GET",
			Actual:   GET.String(),
		},
		{
			Title:    "Test POST",
			Expected: "POST",
			Actual:   POST.String(),
		},
		{
			Title:    "Test PATCH",
			Expected: "PATCH",
			Actual:   PATCH.String(),
		},
		{
			Title:    "Test DELETE",
			Expected: "DELETE",
			Actual:   DELETE.String(),
		},
		{
			Title:    "Test OPTIONS",
			Expected: "OPTIONS",
			Actual:   OPTIONS.String(),
		},
		{
			Title:    "Test CONNECT",
			Expected: "CONNECT",
			Actual:   CONNECT.String(),
		},
		{
			Title:    "Test PUT",
			Expected: "PUT",
			Actual:   PUT.String(),
		},
		{
			Title:    "Test TRACE",
			Expected: "TRACE",
			Actual:   TRACE.String(),
		},
		{
			Title:    "Test HEAD",
			Expected: "HEAD",
			Actual:   HEAD.String(),
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if result := test.Actual; result != test.Expected {
				t.Errorf("Expected: %s, Got: %s", test.Expected, test.Actual)
			}
		})
	}
}
