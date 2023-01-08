package bianrysearch

func Search(nums []int, target int) int {
	const delimiter = 2

	left := 0
	right := len(nums) - 1

	for left <= right {
		pivot := left + (right-left)/delimiter
		if nums[pivot] == target {
			return pivot
		}

		if target < nums[pivot] {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}

	return -1
}

func Search2(nums []int, target int) int {
	left, right := -1, len(nums)
	for right > left+1 {
		m := (left + right) / 2
		if nums[m] >= target {
			right = m
		} else {
			left = m
		}
	}

	return right
}
