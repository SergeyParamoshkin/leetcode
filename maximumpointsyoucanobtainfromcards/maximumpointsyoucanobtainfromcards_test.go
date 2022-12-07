package maximumpointsyoucanobtainfromcards_test

import (
	"testing"

	"github.com/SergeyParamoshkin/leetcode/maximumpointsyoucanobtainfromcards"
)

func Test_maxScore(t *testing.T) {
	t.Parallel()

	type args struct {
		cardPoints []int
		k          int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "okok", args: args{cardPoints: []int{1, 2, 3, 4, 5, 6, 1}, k: 3}, want: 12},
		{name: "okok", args: args{cardPoints: []int{9, 7, 7, 9, 7, 7, 9}, k: 7}, want: 55},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := maximumpointsyoucanobtainfromcards.MaxScore(tt.args.cardPoints, tt.args.k); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
