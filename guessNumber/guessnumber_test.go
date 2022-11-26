package guessnumber_test

import (
	"testing"

	guessNumber "github.com/SergeyParamoshkin/leetcode/guessnumber"
)

func Test_GuessNumber(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n=10",
			args: args{n: 10},
			want: 6,
		},
		{
			name: "n=1",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "n=2",
			args: args{n: 2},
			want: 1,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := guessNumber.GuessNumber(tt.args.n); got != tt.want {
				t.Errorf("guessNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
