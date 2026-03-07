package conversion

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCastCollection(t *testing.T) {
	tests := []struct {
		name       string
		collection []any
		want       []string
		wantErr    bool
	}{
		{
			name:       "Successful cast",
			collection: []any{"a", "b", "c"},
			want:       []string{"a", "b", "c"},
			wantErr:    false,
		},
		{
			name:       "Failed cast",
			collection: []any{"a", 1, "c"},
			want:       nil,
			wantErr:    true,
		},
		{
			name:       "Empty collection",
			collection: []any{},
			want:       []string{},
			wantErr:    false,
		},
		{
			name:       "Nil collection",
			collection: nil,
			want:       nil,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CastCollection[any, string](tt.collection)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestCollectionApply(t *testing.T) {
	tests := []struct {
		name       string
		collection []int
		mutator    func(int) (string, bool)
		want       []string
	}{
		{
			name:       "Successful apply",
			collection: []int{1, 2, 3},
			mutator: func(i int) (string, bool) {
				return fmt.Sprintf("%d", i), true
			},
			want: []string{"1", "2", "3"},
		},
		{
			name:       "Break early",
			collection: []int{1, 2, 3},
			mutator: func(i int) (string, bool) {
				if i == 2 {
					return "", false
				}
				return fmt.Sprintf("%d", i), true
			},
			want: []string{"1", "", ""}, // CollectionApply pre-allocates slice, remaining items are zero-valued
		},
		{
			name:       "Empty collection",
			collection: []int{},
			mutator: func(i int) (string, bool) {
				return "", true
			},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CollectionApply(tt.collection, tt.mutator)
			assert.Equal(t, tt.want, got)
		})
	}
}
