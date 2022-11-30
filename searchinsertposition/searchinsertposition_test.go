package searchinsertposition_test

import (
	"testing"

	"github.com/SergeyParamoshkin/leetcode/searchinsertposition"
)

func TestSearchInsert(t *testing.T) {
	t.Parallel()

	type args struct {
		nums   []int
		target int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "okok", args: args{nums: []int{1, 3, 5, 6}, target: 5}, want: 2},
		{name: "okok", args: args{nums: []int{1, 3, 5, 6}, target: 2}, want: 1},
		{name: "okok", args: args{nums: []int{1, 3, 5, 6}, target: 7}, want: 4},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := searchinsertposition.SearchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("SearchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
