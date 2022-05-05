package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/2
		g := guess(mid)
		if g == 0 {
			return mid
		}
		if g == 1 {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}

	return 0
}
