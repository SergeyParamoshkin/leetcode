package threesum

import "sort"

func ThreeSum(nums []int) [][]int {
	a := [][]int{}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[left]+nums[right]+nums[i] == 0 {
				item := []int{nums[i], nums[left], nums[right]}
				a = append(a, item)
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if nums[left]+nums[right]+nums[i] > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return a
}
