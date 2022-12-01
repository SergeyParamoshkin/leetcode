package validperfectsquare_test

import (
	"testing"

	"github.com/SergeyParamoshkin/leetcode/validperfectsquare"
)

func TestIsPerfectSquare(t *testing.T) {
	t.Parallel()

	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		// {name: "okok", args: args{num: 16}, want: true},
		{name: "okok", args: args{num: 14}, want: false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := validperfectsquare.IsPerfectSquare(tt.args.num); got != tt.want {
				t.Errorf("IsPerfectSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
