package finddistancevaluebetweentwoarrays_test

import (
	"testing"

	"github.com/SergeyParamoshkin/leetcode/finddistancevaluebetweentwoarrays"
)

func TestFindTheDistanceValue(t *testing.T) {
	t.Parallel()

	type args struct {
		arr1 []int
		arr2 []int
		d    int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "okok",
			args: args{
				arr1: []int{1, 4, 2, 3},
				arr2: []int{-4, -3, 6, 10, 20, 30},
				d:    3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		test := tt
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got := finddistancevaluebetweentwoarrays.FindTheDistanceValue(
				test.args.arr1,
				test.args.arr2,
				test.args.d); got != test.want {
				t.Errorf("FindTheDistanceValue() = %v, want %v", got, test.want)
			}
		})
	}
}
