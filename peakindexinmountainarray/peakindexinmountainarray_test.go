package peakindexinmountainarray_test

import (
	"testing"

	peakIndexInMountaInArray "github.com/SergeyParamoshkin/leetcode/peakindexinmountainarray"
)

func Test_peakIndexInMountainArray(t *testing.T) {
	t.Parallel()

	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "ok", args: args{arr: []int{0, 1, 0}}, want: 1},
		{name: "ok", args: args{arr: []int{0, 2, 1, 0}}, want: 1},
		{name: "ok", args: args{arr: []int{0, 10, 5, 2}}, want: 1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := peakIndexInMountaInArray.PeakIndexInMountainArray(tt.args.arr); got != tt.want {
				t.Errorf("peakIndexInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
