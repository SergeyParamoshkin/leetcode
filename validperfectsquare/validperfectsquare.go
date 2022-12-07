package validperfectsquare

func IsPerfectSquare(num int) bool {
	const delimiter = 2

	if num < delimiter {
		return true
	}

	left, right := 2, num/delimiter

	for left <= right {
		/*
			num = 14
			left = 2
			right = 7

			left + (right - left) / delimiter
			2 + (7 - 2) / 2 = 4
		*/
		pivot := left + (right-left)/delimiter //
		guess := pivot * pivot

		if guess == num {
			return true
		}

		if guess > num {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}

	return false
}
