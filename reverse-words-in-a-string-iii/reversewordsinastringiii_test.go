package reversewordsinastringiii_test

import (
	"testing"

	reversewordsinastringiii "github.com/SergeyParamoshkin/leetcode/reverse-words-in-a-string-iii"
)

func TestReverseWords(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "okok", args: args{s: "Let's take LeetCode contest"}, want: "teL ekat edoCteeL tsetnoc"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := reversewordsinastringiii.ReverseWords(tt.args.s); got != tt.want {
				t.Errorf("ReverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
