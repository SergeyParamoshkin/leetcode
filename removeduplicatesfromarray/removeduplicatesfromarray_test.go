package removeduplicatesfromarray_test

import (
	"testing"

	rdfa "github.com/SergeyParamoshkin/leetcode/removeduplicatesfromarray"
)

func TestRemoveDuplicates(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "okok", args: args{nums: []int{1, 1, 2}}, want: 2},
		{name: "okok", args: args{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}, want: 5},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := rdfa.RemoveDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("RemoveDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
