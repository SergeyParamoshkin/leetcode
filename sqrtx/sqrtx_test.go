package sqrtx_test

import (
	"testing"

	"github.com/SergeyParamoshkin/leetcode/sqrtx"
)

func TestMySqrt(t *testing.T) {
	t.Parallel()

	type args struct {
		x int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "okok", args: args{x: 4}, want: 2},
		{name: "okok", args: args{x: 8}, want: 2},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := sqrtx.MySqrt(tt.args.x); got != tt.want {
				t.Errorf("MySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
