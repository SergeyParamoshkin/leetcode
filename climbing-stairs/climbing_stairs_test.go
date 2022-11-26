package climbingstairs_test

import (
	"testing"

	climbingstairs "github.com/SergeyParamoshkin/leetcode/climbing-stairs"
)

func Test_climbStairs(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "n=2", args: args{n: 2}, want: 2},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := climbingstairs.ClimbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
