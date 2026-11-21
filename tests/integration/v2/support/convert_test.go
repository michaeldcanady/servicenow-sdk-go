// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package support

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToInt32(t *testing.T) {
	tests := []struct {
		name    string
		in      int
		want    int32
		wantErr bool
	}{
		{name: "zero", in: 0, want: 0},
		{name: "negative", in: -5, want: -5},
		{name: "max int32", in: math.MaxInt32, want: math.MaxInt32},
		{name: "min int32", in: math.MinInt32, want: math.MinInt32},
		{name: "above max int32", in: math.MaxInt32 + 1, wantErr: true},
		{name: "below min int32", in: math.MinInt32 - 1, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt32(tt.in)
			if tt.wantErr {
				require.Error(t, err, "ToInt32(%d)", tt.in)
				return
			}
			require.NoError(t, err, "ToInt32(%d)", tt.in)
			require.Equal(t, tt.want, got, "ToInt32(%d)", tt.in)
		})
	}
}
