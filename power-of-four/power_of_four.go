package poweroffour

import "math"

func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	
	if n == 1 {
		return true
	}

	c := math.Sqrt(float64(n))

	return int(c)%4 == 0
}
