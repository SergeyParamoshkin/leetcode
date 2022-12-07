package sqrtx

func MySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 2, x/2

	for left <= right {
		pivot := left + (right-left)/2

		num := pivot * pivot

		if num == x {
			return pivot
		}

		if num > x {
			right = pivot - 1
		}

		if num < x {
			left = pivot + 1
		}
	}

	return right
}
