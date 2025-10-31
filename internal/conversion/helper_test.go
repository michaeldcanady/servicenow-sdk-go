package conversion

import (
	"errors"
	"testing"

	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/stretchr/testify/assert"
)

func TestAs2(t *testing.T) {
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "String Pointer && Strict",
			test: func(t *testing.T) {
				in := internal.ToPointer("test")
				out := ""

				err := As2(in, &out, true)

				assert.Equal(t, errors.New("cannot convert type '*string' to type string"), err)
				assert.Equal(t, "", out)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
