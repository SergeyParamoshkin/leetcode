package poweroffour

import "math"

func IsPowerOfFour(number int) bool {
	if number <= 0 {
		return false
	}

	if number == 1 {
		return true
	}

	c := math.Sqrt(float64(number))

	return int(c)%4 == 0
}
