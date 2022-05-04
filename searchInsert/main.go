package main

func searchInsert(nums []int, target int) int {
	pivot, left, right := 0, 0, len(nums)-1

	if nums[0] > target {
		return 0
	}

	for left <= right {
		pivot = left + (right-left)/2
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
