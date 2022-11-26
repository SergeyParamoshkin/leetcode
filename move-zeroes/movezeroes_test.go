package movezeroes_test

import (
	"testing"

	movezeroes "github.com/SergeyParamoshkin/leetcode/move-zeroes"
)

func TestMoveZeroes(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "okok", args: args{nums: []int{}}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			movezeroes.MoveZeroes(tt.args.nums)
		})
	}
}
