package guessnumber

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(num int) int {
	const a = 10
	if a < num {
		return 1
	}

	return -1
}

func GuessNumber(n int) int {
	const delimiter = 2

	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/delimiter

		gst := guess(mid)

		if gst == 0 {
			return mid
		}

		if gst == 1 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return 0
}
