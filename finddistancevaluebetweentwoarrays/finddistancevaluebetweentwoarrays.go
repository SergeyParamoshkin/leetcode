package finddistancevaluebetweentwoarrays

import "sort"

func abs(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}

func FindTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	// FIXME: dirty hack using the sort package
	sort.Ints(arr2)

	guess := 0

	for _, x := range arr1 {
		l, r := 0, len(arr2)-1

		for l <= r {
			mid := l + (r-l)/2
			diff := abs(x - arr2[mid])
			if diff > d {
				if x > arr2[mid] {
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
		}

		guess++
	}

	return guess
}
