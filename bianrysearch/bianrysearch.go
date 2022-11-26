package bianrysearch

func Search(nums []int64, target int64) int64 {
	const delimiter = 2

	left := int64(0)
	right := int64(len(nums) - 1)

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
