package main

func main() {

}

func search(nums []int64, target int64) int64 {
	pivot, left, right := int64(0), int64(0), int64(len(nums)-1)
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

	return -1
}
