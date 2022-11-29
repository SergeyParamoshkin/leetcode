package removeelement_test

import (
	"testing"

	rdfa "github.com/SergeyParamoshkin/leetcode/removeelement"
)

func TestRemoveDuplicates(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
		val  int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "okok", args: args{nums: []int{3, 2, 2, 3}, val: 3}, want: 2},
		{name: "okok", args: args{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2}, want: 5},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := rdfa.RemoveElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("RemoveDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
