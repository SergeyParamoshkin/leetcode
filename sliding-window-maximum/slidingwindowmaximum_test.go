package slidingwindowmaximum_test

import (
	"reflect"
	"testing"

	slidingwindowmaximum "github.com/SergeyParamoshkin/leetcode/sliding-window-maximum"
)

func Test_maxSlidingWindow(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
		k    int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "okok", args: args{nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3}, want: []int{3, 3, 5, 5, 6, 7}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := slidingwindowmaximum.MaxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
