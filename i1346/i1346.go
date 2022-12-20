package i1346

import "sort"

func CheckIfExist(arr []int) bool {
	sort.Ints(arr)
	l, r := 0, len(arr)-1

	for r > l {
		d := arr[l] * 2
		if d < arr[r] {
			l++
		} else if d > arr[r] {
			r = lower_bound(arr, r, 2*arr[l])
		} else {
			return true
		}
	}

	return false
}

func lower_bound(v []int, l int, t int) int {
	r := len(v)
	for l <= r {
		m := l + (r-l)/2
		if v[m] < t {
			l = m
		} else if v[m] == t {
			return m
		} else {
			r = m
		}
	}

	return l
}
