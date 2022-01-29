package threesum_test

import (
	"leetcode/threesum"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {

	type Test struct {
		Name string
		In   []int
		Want [][]int
	}
	testCases := []Test{
		{
			Name: "zero case",
			In:   []int{0, 0, 0, 0},
			Want: [][]int{{0, 0, 0}},
		},
		{
			Name: "ok",
			In:   []int{-1, 0, 1, 2, -1, -4},
			Want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.Want, threesum.ThreeSum(tc.In))
		})
	}
}
