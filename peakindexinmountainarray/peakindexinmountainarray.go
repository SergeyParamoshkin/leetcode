package peakindexinmountainarray

func PeakIndexInMountainArray(arr []int) int {
	low, height := 0, len(arr)-1

	const delimiter = 2
	for low < height {
		mid := low + (height-low)/delimiter
		if arr[mid] < arr[mid+1] {
			low = mid + 1
		} else {
			height = mid
		}
	}

	return low
}
