package fibonacci_test

import (
	"testing"

	"github.com/SergeyParamoshkin/leetcode/fibonacci-number"
)

func Test_fib(t *testing.T) {
	t.Parallel()

	type args struct {
		number int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "n=2", args: args{number: 2}, want: 1},
		{name: "n=3", args: args{number: 3}, want: 2},
		{name: "n=4", args: args{number: 4}, want: 3},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := fibonacci.Fib(tt.args.number); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
