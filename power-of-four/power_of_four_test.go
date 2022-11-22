package poweroffour_test

import (
	"testing"

	poweroffour "github.com/SergeyParamoshkin/leetcode/power-of-four"
)

func Test_isPowerOfFour(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "n=16", args: args{n: 16}, want: true},
		{name: "n=5", args: args{n: 5}, want: false},
		{name: "n=1", args: args{n: 1}, want: true},
		{name: "n=-2147483648", args: args{n: -2147483648}, want: false},
	}

	t.Run("test", func(t *testing.T) {
		t.Parallel()
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := poweroffour.IsPowerOfFour(tt.args.n); got != tt.want {
					t.Errorf("isPowerOfFour() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}
