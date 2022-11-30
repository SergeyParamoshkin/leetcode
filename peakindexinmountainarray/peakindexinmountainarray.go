package peakindexinmountainarray

func PeakIndexInMountainArray(arr []int) int {
	i := 0
	for arr[i] < arr[i+1] {
		i++
	}

	return i
}
