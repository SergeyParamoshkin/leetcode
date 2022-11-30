package searchinsertposition

func SearchInsert(nums []int, target int) int {
	const delimiter = 2

	pivot, left, right := 0, 0, len(nums)-1

	if nums[0] > target {
		return 0
	}

	for left <= right {
		pivot = left + (right-left)/delimiter
		if nums[pivot] == target {
			return pivot
		}

		if target < nums[pivot] {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}

	if nums[pivot] > target {
		return pivot
	}

	return right + 1
}
