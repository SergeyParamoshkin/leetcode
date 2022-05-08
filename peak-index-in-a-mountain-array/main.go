package main

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] == arr[mid-1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left

	return 0
}
