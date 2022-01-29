package threesum_test

import (
	"reflect"
	"testing"

	"github.com/SergeyParamoshkin/leetcode/threesum"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "zero case",
			args: args{[]int{0, 0, 0, 0}},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "ok",
			args: args{[]int{-1, 0, 1, 2, -1, -4}},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threesum.ThreeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				// t.Errorf("ThreeSum() = %v, want %v", got, tt.want)
				assert.Equal(t, tt.want, threesum.ThreeSum(tt.args.nums))
			}
		})
	}
}
