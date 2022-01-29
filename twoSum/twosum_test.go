package twosum_test

import (
	"testing"

	"github.com/SergeyParamoshkin/leetcode/twosum"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	a := []int{2, 7, 11, 15}
	got := twosum.TwoSum(a, 9)
	assert.Equal(t, []int{0, 1}, got)
}
